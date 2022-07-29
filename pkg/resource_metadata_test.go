package kyaml

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestRNodeGetNamespace(t *testing.T) {
	n1 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {namespace: ns, name: r}"), n1); err != nil {
		t.Fatal(err)
	}
	r1 := &RNode{N: n1}
	if ns, err := r1.GetNamespace(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if ns != "ns" {
		t.Errorf("unexpected namespace: %s", ns)
	}

	n2 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {name: r}"), n2); err != nil {
		t.Fatal(err)
	}
	r2 := &RNode{N: n2}
	if ns, err := r2.GetNamespace(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if ns != "" {
		t.Errorf("unexpected namespace: %s", ns)
	}

	n3 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("kind: k"), n3); err != nil {
		t.Fatal(err)
	}
	r3 := &RNode{N: n3}
	if ns, err := r3.GetNamespace(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if ns != "" {
		t.Errorf("unexpected namespace: %s", ns)
	}
}

func TestRNodeSetNamespace(t *testing.T) {
	n1 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {namespace: ns1, name: r}"), n1); err != nil {
		t.Fatal(err)
	}
	r1 := &RNode{N: n1}
	if ns, err := r1.GetNamespace(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if ns != "ns1" {
		t.Errorf("unexpected namespace: %s", ns)
	} else if err := r1.SetNamespace("ns2"); err != nil {
		t.Errorf("failed setting namespace: %v", err)
	} else if ns, err := r1.GetNamespace(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if ns != "ns2" {
		t.Errorf("unexpected namespace: %s", ns)
	}
}

func TestRNodeGetName(t *testing.T) {
	n1 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {namespace: ns, name: r}"), n1); err != nil {
		t.Fatal(err)
	}
	r1 := &RNode{N: n1}
	if ns, err := r1.GetName(); err != nil {
		t.Errorf("failed getting name: %v", err)
	} else if ns != "r" {
		t.Errorf("unexpected name: %s", ns)
	}

	n2 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {namespace: ns}"), n2); err != nil {
		t.Fatal(err)
	}
	r2 := &RNode{N: n2}
	if nm, err := r2.GetName(); err != nil {
		t.Errorf("failed getting name: %v", err)
	} else if nm != "" {
		t.Errorf("unexpected name: %s", nm)
	}

	n3 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("kind: k"), n3); err != nil {
		t.Fatal(err)
	}
	r3 := &RNode{N: n3}
	if nm, err := r3.GetName(); err != nil {
		t.Errorf("failed getting name: %v", err)
	} else if nm != "" {
		t.Errorf("unexpected name: %s", nm)
	}
}

func TestRNodeSetName(t *testing.T) {
	n1 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {namespace: ns1, name: r1}"), n1); err != nil {
		t.Fatal(err)
	}
	r1 := &RNode{N: n1}
	if nm, err := r1.GetName(); err != nil {
		t.Errorf("failed getting namespace: %v", err)
	} else if nm != "r1" {
		t.Errorf("unexpected name: %s", nm)
	} else if err := r1.SetName("r2"); err != nil {
		t.Errorf("failed setting name: %v", err)
	} else if nm, err := r1.GetName(); err != nil {
		t.Errorf("failed getting name: %v", err)
	} else if nm != "r2" {
		t.Errorf("unexpected name: %s", nm)
	}
}

func TestRNodeGetAnnotation(t *testing.T) {
	n1 := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("metadata: {annotations: {foo: bar}}"), n1); err != nil {
		t.Fatal(err)
	}
	r1 := &RNode{N: n1}
	if v, err := r1.GetAnnotation("foo"); err != nil {
		t.Errorf("failed getting annotation: %v", err)
	} else if v != "bar" {
		t.Errorf("unexpected value: %s", v)
	}
	if v, err := r1.GetAnnotation("foo1"); err != nil {
		t.Errorf("failed getting annotation: %v", err)
	} else if v != "" {
		t.Errorf("unexpected value: %s", v)
	}
}

func TestRNodeSetAnnotation(t *testing.T) {
	n := &yaml.Node{}
	if err := yaml.Unmarshal([]byte("kind: k"), n); err != nil {
		t.Fatal(err)
	}
	r := &RNode{N: n}
	if err := r.SetAnnotation("foo1", "bar1"); err != nil {
		t.Errorf("failed setting annotation: %v", err)
	} else if v, err := r.GetAnnotation("foo1"); err != nil {
		t.Errorf("failed getting annotation: %v", err)
	} else if v != "bar1" {
		t.Errorf("unexpected value: %s", v)
	} else if err := r.SetAnnotation("foo2", "bar2"); err != nil {
		t.Errorf("failed setting annotation: %v", err)
	} else if v, err := r.GetAnnotation("foo2"); err != nil {
		t.Errorf("failed getting annotation: %v", err)
	} else if v != "bar2" {
		t.Errorf("unexpected value: %s", v)
	}
}
