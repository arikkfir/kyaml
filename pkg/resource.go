package pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type RNode struct {
	N *yaml.Node
}

func (r *RNode) GetAPIVersion() (string, error) {
	if value, err := GetMappingKeyValue(r.N, "apiVersion"); err != nil {
		return "", err
	} else {
		return value.(string), nil
	}
}

func (r *RNode) GetKind() (string, error) {
	if value, err := GetMappingKeyValue(r.N, "kind"); err != nil {
		return "", err
	} else {
		return value.(string), nil
	}
}

func (r *RNode) GetAPIGroupAndVersion() (string, string, error) {
	if apiVersion, err := r.GetAPIVersion(); err != nil {
		return "", "", fmt.Errorf("failed getting apiVersion: %w", err)
	} else if lastSlashIndex := strings.LastIndex(apiVersion, "/"); lastSlashIndex < 0 {
		return "", apiVersion, nil
	} else {
		return apiVersion[0:lastSlashIndex], apiVersion[lastSlashIndex+1:], nil
	}
}
