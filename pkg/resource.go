package kyaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type RNode struct {
	N *yaml.Node
}

func (r *RNode) GetAPIVersion() (string, error) {
	if value, err := GetMappingKeyValue(r.N, "apiVersion", yaml.ScalarNode, "!!str"); err != nil {
		return "", err
	} else if value == nil {
		return "", nil
	} else {
		return value.(string), nil
	}
}

func (r *RNode) GetKind() (string, error) {
	if value, err := GetMappingKeyValue(r.N, "kind", yaml.ScalarNode, "!!str"); err != nil {
		return "", err
	} else if value == nil {
		return "", nil
	} else {
		return value.(string), nil
	}
}

func (r *RNode) GetAPIGroupAndVersion() (string, string, error) {
	if value, err := r.GetAPIVersion(); err != nil {
		return "", "", fmt.Errorf("failed getting apiVersion: %w", err)
	} else if value == "" {
		return "", "", fmt.Errorf("apiVersion is missing")
	} else if lastSlashIndex := strings.LastIndex(value, "/"); lastSlashIndex < 0 {
		return "", value, nil
	} else {
		return value[0:lastSlashIndex], value[lastSlashIndex+1:], nil
	}
}

func (r *RNode) IsValid() (bool, error) {
	if apiVersion, err := r.GetAPIVersion(); err != nil {
		return false, fmt.Errorf("failed getting apiVersion: %w", err)
	} else if apiVersion == "" {
		return false, fmt.Errorf("apiVersion is missing")
	} else if kind, err := r.GetKind(); err != nil {
		return false, fmt.Errorf("failed getting kind: %w", err)
	} else if kind == "" {
		return false, fmt.Errorf("kind is missing")
	} else {
		return true, nil
	}
}
