package main

import (
	"encoding/json"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo
	name string
}

type provider struct {
	app.Compo
	name    string
	version string
}

var (
	schemas = &tfjson.ProviderSchemas{}
)

func init() {
	err := json.Unmarshal(aws, schemas)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (h *home) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			&provider{name: "aws", version: "v3.2.0"},
		),
	)
}

func (p *provider) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text(p.name),
		),
		app.P().Body(
			app.Text(p.version),
			app.Br(),
			app.Text(fmt.Sprintf("Resources = %d", count(schemas.Schemas[p.name].ResourceSchemas))),
			app.Br(),
			app.Text(fmt.Sprintf("Data Sources = %d", count(schemas.Schemas[p.name].DataSourceSchemas))),
		),
	)
}

func main() {
	app.Route("/", &home{})
	app.Run()
}

func count(m map[string]*tfjson.Schema) int {
	return len(m)
}
