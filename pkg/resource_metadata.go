package pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/labels"
)

func (r *RNode) GetNamespace() (string, error) {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", false, yaml.MappingNode, "!!map")
	if err != nil {
		return "", fmt.Errorf("failed to get metadata node: %w", err)
	} else if metadataNode == nil {
		return "", nil
	} else if value, err := GetMappingKeyValue(metadataNode, "namespace", yaml.ScalarNode, "!!str"); err != nil {
		return "", fmt.Errorf("failed to get namespace: %w", err)
	} else if value == nil {
		return "", nil
	} else {
		return value.(string), nil
	}
}

func (r *RNode) SetNamespace(value string) error {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", true, yaml.MappingNode, "!!map")
	if err != nil {
		return fmt.Errorf("failed to get metadata node: %w", err)
	} else if err := SetMappingKeyValue(metadataNode, "namespace", value, yaml.ScalarNode, "!!str"); err != nil {
		return fmt.Errorf("failed to set namespace value: %w", err)
	} else {
		return nil
	}
}

func (r *RNode) GetName() (string, error) {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", false, yaml.MappingNode, "!!map")
	if err != nil {
		return "", fmt.Errorf("failed to get metadata node: %w", err)
	} else if metadataNode == nil {
		return "", nil
	} else if value, err := GetMappingKeyValue(metadataNode, "name", yaml.ScalarNode, "!!str"); err != nil {
		return "", fmt.Errorf("failed to get name: %w", err)
	} else if value == nil {
		return "", nil
	} else {
		return value.(string), nil
	}
}

func (r *RNode) SetName(value string) error {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", true, yaml.MappingNode, "!!map")
	if err != nil {
		return fmt.Errorf("failed to get metadata node: %w", err)
	} else if err := SetMappingKeyValue(metadataNode, "name", value, yaml.ScalarNode, "!!str"); err != nil {
		return fmt.Errorf("failed to set name value: %w", err)
	} else {
		return nil
	}
}

func (r *RNode) GetAnnotation(key string) (string, error) {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", false, yaml.MappingNode, "!!map")
	if err != nil {
		return "", fmt.Errorf("failed to get metadata node: %w", err)
	} else if metadataNode == nil {
		return "", nil
	} else if annotationsNode, err := GetMappingKeyNode(metadataNode, "annotations", false, yaml.MappingNode, "!!map"); err != nil {
		return "", fmt.Errorf("failed to get annotations node: %w", err)
	} else if annotationsNode == nil {
		return "", nil
	} else if value, err := GetMappingKeyValue(annotationsNode, key, yaml.ScalarNode, "!!str"); err != nil {
		return "", fmt.Errorf("failed to get annotation value: %w", err)
	} else if value == nil {
		return "", nil
	} else {
		return value.(string), nil
	}
}

func (r *RNode) SetAnnotation(key string, value interface{}) error {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", true, yaml.MappingNode, "!!map")
	if err != nil {
		return fmt.Errorf("failed to get metadata node: %w", err)
	} else if annotationsNode, err := GetMappingKeyNode(metadataNode, "annotations", true, yaml.MappingNode, "!!map"); err != nil {
		return fmt.Errorf("failed to get annotations node: %w", err)
	} else if err := SetMappingKeyValue(annotationsNode, key, value, yaml.ScalarNode, "!!str"); err != nil {
		return fmt.Errorf("failed to set annotation value: %w", err)
	} else {
		return nil
	}
}

func (r *RNode) GetLabels() (map[string]string, error) {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", false, yaml.MappingNode, "!!map")
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata node: %w", err)
	} else if metadataNode == nil {
		return nil, nil
	} else if labelsNode, err := GetMappingKeyNode(metadataNode, "labels", false, yaml.MappingNode, "!!map"); err != nil {
		return nil, fmt.Errorf("failed to get labels node: %w", err)
	} else if labelsNode == nil {
		return nil, nil
	} else {
		m := make(map[string]string)
		for i := 0; i < len(labelsNode.Content); i += 2 {
			key := labelsNode.Content[i].Value
			value := labelsNode.Content[i].Value
			m[key] = value
		}
		return m, nil
	}
}

func (r *RNode) GetLabel(key string) (string, error) {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", false, yaml.MappingNode, "!!map")
	if err != nil {
		return "", fmt.Errorf("failed to get metadata node: %w", err)
	} else if metadataNode == nil {
		return "", nil
	} else if labelsNode, err := GetMappingKeyNode(metadataNode, "labels", false, yaml.MappingNode, "!!map"); err != nil {
		return "", fmt.Errorf("failed to get labels node: %w", err)
	} else if labelsNode == nil {
		return "", nil
	} else if value, err := GetMappingKeyValue(labelsNode, key, yaml.ScalarNode, "!!str"); err != nil {
		return "", fmt.Errorf("failed to get label value: %w", err)
	} else {
		return value.(string), nil
	}
}

func (r *RNode) SetLabel(key string, value interface{}) error {
	metadataNode, err := GetMappingKeyNode(r.N, "metadata", true, yaml.MappingNode, "!!map")
	if err != nil {
		return fmt.Errorf("failed to get metadata node: %w", err)
	} else if labelsNode, err := GetMappingKeyNode(metadataNode, "labels", true, yaml.MappingNode, "!!map"); err != nil {
		return fmt.Errorf("failed to get labels node: %w", err)
	} else if err := SetMappingKeyValue(labelsNode, key, value, yaml.ScalarNode, "!!str"); err != nil {
		return fmt.Errorf("failed to set label value: %w", err)
	} else {
		return nil
	}
}

func (r *RNode) IsMatchingLabelSelector(selector string) (bool, error) {
	s, err := labels.Parse(selector)
	if err != nil {
		return false, fmt.Errorf("failed parsing label selector '%s': %w", selector, err)
	}
	labelsMap, err := r.GetLabels()
	if err != nil {
		return false, fmt.Errorf("failed getting node labels: %w", err)
	}
	return s.Matches(labels.Set(labelsMap)), nil
}
