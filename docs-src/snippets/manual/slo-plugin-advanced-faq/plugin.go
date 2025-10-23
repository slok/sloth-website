package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	prommodel "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/promql/parser"

	"github.com/slok/sloth/pkg/common/conventions"
	"github.com/slok/sloth/pkg/common/model"
	promutils "github.com/slok/sloth/pkg/common/utils/prometheus"
	pluginslov1 "github.com/slok/sloth/pkg/prometheus/plugin/slo/v1"
)

const (
	PluginVersion = "prometheus/slo/v1"
	PluginID      = "github.com/slok/sloth/issues/618"
)

const (
	metricNameBadDefault           = "slo:bad_events:count%s"
	metricNameTotalDefault         = "slo:total_events:count%s"
	sliErrorTotalWindowQueryMetric = `sum_over_time(%[1]s%[4]s[%[3]s]) / ignoring (sloth_window) sum_over_time(%[2]s%[4]s[%[3]s])`
)

type Config struct {
	AggregationWindow prommodel.Duration `json:"aggregation_window,omitempty"`
}

func NewPlugin(configData json.RawMessage, _ pluginslov1.AppUtils) (pluginslov1.Plugin, error) {
	config := Config{}
	err := json.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	if config.AggregationWindow == 0 {
		config.AggregationWindow = prommodel.Duration(5 * time.Minute)
	}

	return plugin{config: config}, nil
}

type plugin struct {
	config Config
}

func (p plugin) ProcessSLO(ctx context.Context, request *pluginslov1.Request, result *pluginslov1.Result) error {
	if request.SLO.SLI.Events == nil {
		return fmt.Errorf("only supported with SLI events type")
	}

	errorMetricCounter, err := getCounterMetricFromSLIQuery(request.SLO.SLI.Events.ErrorQuery)
	if err != nil {
		return fmt.Errorf("could not get counter metric from error events SLI query: %w", err)
	}

	totalMetricCounter, err := getCounterMetricFromSLIQuery(request.SLO.SLI.Events.TotalQuery)
	if err != nil {
		return fmt.Errorf("could not get counter metric from total events SLI query: %w", err)
	}

	// Add precomputed buckets into a new rule group.
	badRec := rulefmt.Rule{
		Record: fmt.Sprintf(metricNameBadDefault, p.config.AggregationWindow),
		Expr:   fmt.Sprintf("increase(%s[%s])", errorMetricCounter, p.config.AggregationWindow),
		Labels: result.SLORules.SLIErrorRecRules.Rules[0].Labels,
	}
	totalRec := rulefmt.Rule{
		Record: fmt.Sprintf(metricNameTotalDefault, p.config.AggregationWindow),
		Expr:   fmt.Sprintf("increase(%s[%s])", totalMetricCounter, p.config.AggregationWindow),
		Labels: result.SLORules.SLIErrorRecRules.Rules[0].Labels,
	}
	result.SLORules.ExtraRules = append(result.SLORules.ExtraRules, model.PromRuleGroup{
		Name:     "sloth-slo-events-precomputed-" + request.SLO.ID,
		Interval: time.Duration(p.config.AggregationWindow),
		Rules:    []rulefmt.Rule{badRec, totalRec},
	})

	// Replace long window query to use the precomputed ones.
	sliLongWindowRecord := conventions.GetSLIErrorMetric(request.SLO.TimeWindow)
	sloLabels := promutils.LabelsToPromFilter(conventions.GetSLOIDPromLabels(request.SLO))
	for i, r := range result.SLORules.SLIErrorRecRules.Rules {
		if r.Record == sliLongWindowRecord {
			result.SLORules.SLIErrorRecRules.Rules[i].Expr = fmt.Sprintf(sliErrorTotalWindowQueryMetric, badRec.Record, totalRec.Record, promutils.TimeDurationToPromStr(request.SLO.TimeWindow), sloLabels)
			break
		}
	}

	return nil
}

func getCounterMetricFromSLIQuery(query string) (string, error) {
	tmpl, err := template.New("").Parse(query)
	if err != nil {
		return "", fmt.Errorf("could not parse events SLI query: %w", err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, map[string]any{"window": "5m"})
	if err != nil {
		return "", fmt.Errorf("could not execute events SLI query template: %w", err)
	}

	inferredMetric, err := inferMetricFromQuery(b.String())
	if err != nil {
		return "", fmt.Errorf("could not infer metrics from query: %w", err)
	}
	return inferredMetric, nil
}

const levelStep = 2

// Adapted from https://github.com/cloudflare/pint/blob/main/cmd/pint/parse.go.
func inferMetricFromQuery(query string) (string, error) {
	expr, err := parser.ParseExpr(query)
	if err != nil {
		return "", err
	}
	res := inferMetric(expr, 0)
	if res == "" {
		return "", fmt.Errorf("could not infer metric from query")
	}
	return res, nil
}

func inferMetric(node parser.Node, level int) string {
	level += levelStep

	switch n := node.(type) {
	case *parser.MatrixSelector:
		return n.VectorSelector.String()
	case *parser.VectorSelector:
		return n.String()
	case parser.Expressions:
		for _, e := range n {
			res := inferMetric(e, level+levelStep)
			if res != "" {
				return res
			}
		}
	case *parser.AggregateExpr:
		level += levelStep
		res := inferMetric(n.Expr, level+levelStep)
		if res != "" {
			return res
		}
	case *parser.BinaryExpr:
		level += levelStep
		res := inferMetric(n.LHS, level+levelStep)
		if res != "" {
			return res
		}
		res = inferMetric(n.RHS, level+levelStep)
		if res != "" {
			return res
		}
	case *parser.Call:
		level += levelStep
		res := inferMetric(n.Args, level+levelStep)
		if res != "" {
			return res
		}
	case *parser.ParenExpr:
		level += levelStep
		res := inferMetric(n.Expr, level+levelStep)
		if res != "" {
			return res
		}
	case *parser.SubqueryExpr:
		level += levelStep
		res := inferMetric(n.Expr, level+levelStep)
		if res != "" {
			return res
		}
	default:
		level += levelStep
	}

	return ""
}
