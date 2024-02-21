package hello

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	return nil
}

var HelloApp = gorf.GorfBaseApp{
	Name:         "Hello",
	RouteHandler: Urls,
	SetUpHandler: setup,
}
