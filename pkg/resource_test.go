package kyaml

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestRNodeGetAPIVersion(t *testing.T) {
	n := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("apiVersion: v1\nkind: ServiceAccount"), n); err != nil {
		t.Fatal(err)
	}

	r := &RNode{N: n}
	if apiVersion, err := r.GetAPIVersion(); err != nil {
		t.Errorf("failed getting apiVersion: %v", err)
	} else if apiVersion != "v1" {
		t.Errorf("unexpected apiVersion: %s", apiVersion)
	}
}

func TestRNodeGetAPIVersionWhenMissing(t *testing.T) {
	n := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("kind: ServiceAccount"), n); err != nil {
		t.Fatal(err)
	}

	r := &RNode{N: n}
	if apiVersion, err := r.GetAPIVersion(); err != nil {
		t.Errorf("failed getting apiVersion: %v", err)
	} else if apiVersion != "" {
		t.Errorf("unexpected apiVersion: %s", apiVersion)
	}
}

func TestRNodeGetKind(t *testing.T) {
	n := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("apiVersion: v1\nkind: ServiceAccount"), n); err != nil {
		t.Fatal(err)
	}

	r := &RNode{N: n}
	if kind, err := r.GetKind(); err != nil {
		t.Errorf("failed getting kind: %v", err)
	} else if kind != "ServiceAccount" {
		t.Errorf("unexpected kind: %s", kind)
	}
}

func TestRNodeGetKindWhenMissing(t *testing.T) {
	n := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("apiVersion: v1"), n); err != nil {
		t.Fatal(err)
	}

	r := &RNode{N: n}
	if kind, err := r.GetKind(); err != nil {
		t.Errorf("failed getting kind: %v", err)
	} else if kind != "" {
		t.Errorf("unexpected kind: %s", kind)
	}
}
