package kstream

import (
	"context"
	. "github.com/arikkfir/gstream/pkg/types"
	"github.com/arikkfir/kyaml/pkg"
	"gopkg.in/yaml.v3"
)

func LabelResource(name string, value interface{}) NodeProcessor {
	return func(ctx context.Context, n *yaml.Node) error {
		r := &pkg.RNode{N: n}
		return r.SetLabel(name, value)
	}
}
