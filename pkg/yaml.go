package kyaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func GetMappingKeyNode(n *yaml.Node, key string, create bool, expectedKind yaml.Kind, expectedTag string) (*yaml.Node, error) {
	if n.Kind == yaml.DocumentNode {
		n = n.Content[0]
	}
	if n.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("node is not a mapping")
	}
	for i := 0; i < len(n.Content); i += 2 {
		keyNode := n.Content[i]
		if keyNode.Value == key {
			valueNode := n.Content[i+1]
			if valueNode.Kind != expectedKind {
				return nil, fmt.Errorf("expected value node kind to be %v, got %v", expectedKind, valueNode.Kind)
			} else if valueNode.Tag != expectedTag {
				return nil, fmt.Errorf("expected value node tag to be %s, got %s", expectedTag, valueNode.Tag)
			} else {
				return valueNode, nil
			}
		}
	}
	if create {
		valueNode := &yaml.Node{Kind: expectedKind, Tag: expectedTag}
		n.Content = append(
			n.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: key},
			valueNode,
		)
		return valueNode, nil
	}
	return nil, nil
}

func GetMappingKeyValue(n *yaml.Node, key string, expectedKind yaml.Kind, expectedTag string) (interface{}, error) {
	var value interface{}
	if valueNode, err := GetMappingKeyNode(n, key, false, expectedKind, expectedTag); err != nil {
		return nil, err
	} else if valueNode == nil {
		return nil, nil
	} else if err := valueNode.Decode(&value); err != nil {
		return nil, fmt.Errorf("failed to decode value node: %w", err)
	} else {
		return value, nil
	}
}

func SetMappingKeyValue(n *yaml.Node, key string, value interface{}, expectedKind yaml.Kind, expectedTag string) error {
	if valueNode, err := GetMappingKeyNode(n, key, true, expectedKind, expectedTag); err != nil {
		return fmt.Errorf("failed to get value node: %w", err)
	} else if err := valueNode.Encode(value); err != nil {
		return fmt.Errorf("failed to set value in node: %w", err)
	} else {
		return nil
	}
}
