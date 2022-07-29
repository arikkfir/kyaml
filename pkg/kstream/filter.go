package kstream

import (
	"context"
	"fmt"
	. "github.com/arikkfir/gstream/pkg/types"
	"github.com/arikkfir/kyaml/pkg"
	"gopkg.in/yaml.v3"
)

func FilterResource(includes []kyaml.TargetingFilter, excludes []kyaml.TargetingFilter) NodeTransformer {
	return func(ctx context.Context, n *yaml.Node, c chan *yaml.Node) error {
		r := &kyaml.RNode{N: n}
		if matches, err := r.IsMatchingFilter(includes, excludes); err != nil {
			return fmt.Errorf("failed matching filter to node: %w", err)
		} else if matches {
			c <- n
		}
		return nil
	}
}
