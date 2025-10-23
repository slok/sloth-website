package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"testing/fstest"

	sloth "github.com/slok/sloth/pkg/lib"
)

const (
	remotePluginURL = "https://gist.githubusercontent.com/slok/9408924d844049b675a8da9185422b40/raw/e36b7dd881f3465cb568db3edf5afb0d38b5427f/sloth-slo-plugin.go"
)

const sloSpec = `---
version: "prometheus/v1"
service: "myservice"
slo_plugins:
  chain:
  - id: "sloth-examples/add_description_to_info/v1"
slos:
  - name: "requests-availability"
    objective: 99.9
    description: "Common SLO based on availability for HTTP request responses."
    sli:
      events:
        error_query: sum(rate(http_request_duration_seconds_count{job="myservice",code=~"(5..|429)"}[{{.window}}]))
        total_query: sum(rate(http_request_duration_seconds_count{job="myservice"}[{{.window}}]))
    alerting:
      page_alert:
        disable: true
      ticket_alert:
        disable: true
`

func main() {
	ctx := context.Background()

	// Download the remote plugin.
	resp, err := http.Get(remotePluginURL)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("failed to download remote plugin")
	}

	pluginB, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Load the downloaded plugin in an in-memory fs.
	f := fstest.MapFS{}
	f["remote-example/plugin.go"] = &fstest.MapFile{Data: pluginB}

	gen, err := sloth.NewPrometheusSLOGenerator(sloth.PrometheusSLOGeneratorConfig{PluginsFS: []fs.FS{f}})
	if err != nil {
		panic(fmt.Errorf("could not create SLO generator: %w", err))
	}

	// Generate SLOs.
	result, err := gen.GenerateFromRaw(ctx, []byte(sloSpec))
	if err != nil {
		panic(fmt.Sprintf("could not generate SLOs: %v", err))
	}
	err = sloth.WriteResultAsPrometheusStd(ctx, *result, os.Stdout)
	if err != nil {
		panic(fmt.Sprintf("could not write result: %v", err))
	}
}
