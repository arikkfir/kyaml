package kstream

import (
	"context"
	"github.com/arikkfir/gstream/pkg/types"
	"github.com/arikkfir/kyaml/pkg"
	"gopkg.in/yaml.v3"
)

func AnnotateResource(name string, value interface{}) types.NodeProcessor {
	return func(ctx context.Context, n *yaml.Node) error {
		r := &kyaml.RNode{N: n}
		return r.SetAnnotation(name, value)
	}
}
