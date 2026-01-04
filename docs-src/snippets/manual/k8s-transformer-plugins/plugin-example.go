package plugin

import (
	"context"
	"github.com/slok/sloth/pkg/common/model"
	k8sutils "github.com/slok/sloth/pkg/common/utils/k8s"
	plugink8stransformv1 "github.com/slok/sloth/pkg/prometheus/plugin/k8stransform/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	PluginVersion = "prometheus/k8stransform/v1"
	PluginID      = "mycompany.com/prometheus-configmap/v1"
)

func NewPlugin() (plugink8stransformv1.Plugin, error) {
	return plugin{}, nil
}

type plugin struct{}

func (p plugin) TransformK8sObjects(
	ctx context.Context,
	kmeta model.K8sMeta,
	sloResult model.PromSLOGroupResult,
) (*plugink8stransformv1.K8sObjects, error) {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("v1")
	u.SetKind("ConfigMap")
	u.SetNamespace(kmeta.Namespace)
	u.SetName(kmeta.Name)
	u.SetLabels(kmeta.Labels)

	// Collect all rules
	groups := []model.PromRuleGroup{}
	for _, slo := range sloResult.SLOResults {
		groups = append(groups,
			slo.PrometheusRules.SLIErrorRecRules,
			slo.PrometheusRules.MetadataRecRules,
			slo.PrometheusRules.AlertRules,
		)
	}

	// Convert to YAML
	rulesYAML, err := k8sutils.UnstructuredToYAMLString(
		map[string]any{"groups": groups},
	)
	if err != nil {
		return nil, err
	}

	u.Object["data"] = map[string]any{
		"prometheus-rules.yaml": rulesYAML,
	}

	return &plugink8stransformv1.K8sObjects{
		Items: []*unstructured.Unstructured{u},
	}, nil
}
