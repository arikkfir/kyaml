package kyaml

import "fmt"

type TargetingFilter struct {
	APIVersion    string `json:"apiVersion" yaml:"apiVersion" mapstructure:"apiVersion"`
	Kind          string `json:"kind" yaml:"kind" mapstructure:"kind"`
	Namespace     string `json:"namespace" yaml:"namespace" mapstructure:"namespace"`
	Name          string `json:"name" yaml:"name" mapstructure:"name"`
	LabelSelector string `json:"labelSelector" yaml:"labelSelector" mapstructure:"labelSelector"`
}

func (r *RNode) IsMatchingFilter(includes []TargetingFilter, excludes []TargetingFilter) (bool, error) {
	if apiVersion, err := r.GetAPIVersion(); err != nil {
		return false, fmt.Errorf("failed getting apiVersion: %w", err)
	} else if kind, err := r.GetKind(); err != nil {
		return false, fmt.Errorf("failed getting kind: %w", err)
	} else if namespace, err := r.GetNamespace(); err != nil {
		return false, fmt.Errorf("failed getting namespace: %w", err)
	} else if name, err := r.GetName(); err != nil {
		return false, fmt.Errorf("failed getting name: %w", err)
	} else {
		included := len(includes) == 0
		excluded := false
		for _, f := range includes {
			if f.APIVersion != "" && f.APIVersion != apiVersion {
				continue
			} else if f.Kind != "" && f.Kind != kind {
				continue
			} else if f.Namespace != "" && f.Namespace != namespace {
				continue
			} else if f.Name != "" && f.Name != name {
				continue
			} else if f.LabelSelector != "" {
				if matches, err := r.IsMatchingLabelSelector(f.LabelSelector); err != nil {
					return false, fmt.Errorf("failed matching label selector '%s' to node: %w", f.LabelSelector, err)
				} else if !matches {
					continue
				}
			}
			included = true
		}
		for _, f := range excludes {
			if f.APIVersion != "" && f.APIVersion != apiVersion {
				continue
			} else if f.Kind != "" && f.Kind != kind {
				continue
			} else if f.Namespace != "" && f.Namespace != namespace {
				continue
			} else if f.Name != "" && f.Name != name {
				continue
			} else if f.LabelSelector != "" {
				if matches, err := r.IsMatchingLabelSelector(f.LabelSelector); err != nil {
					return false, fmt.Errorf("failed matching label selector '%s' to node: %w", f.LabelSelector, err)
				} else if !matches {
					continue
				}
			}
			excluded = true
		}
		if included && !excluded {
			return true, nil
		}
		return false, nil
	}
}
