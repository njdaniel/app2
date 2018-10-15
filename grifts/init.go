package grifts

import (
	"gobuffalo-example/app2/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
