package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	icon := app.Icon{
		Default: "/web/terraform-icon.png",
	}

	h := &app.Handler{
		Title:  "Terraform Generate",
		Author: "Kevin Head",
		Icon:   icon,
	}

	if err := http.ListenAndServe(":8000", h); err != nil {
		panic(err)
	}
}
