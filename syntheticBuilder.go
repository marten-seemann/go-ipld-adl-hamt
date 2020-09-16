package hamt

import "github.com/ipld/go-ipld-prime"

// TODO : the synthetic NodeBuilder + NodeAssembler that act like a single unified map go here

var _ ipld.NodeBuilder = (*builder)(nil)

type builder struct {
}

func (b *builder) Build() ipld.Node { return nil }
func (b *builder) Reset() {}

func (b *builder) BeginMap(sizeHint int) (ipld.MapAssembler, error) { return nil, nil }
func (b *builder) BeginList(sizeHint int) (ipld.ListAssembler, error) { return nil, nil }
func (b *builder) AssignNull() error { return nil }
func (b *builder) AssignBool(bool) error { return nil }
func (b *builder) AssignInt(int) error { return nil }
func (b *builder) AssignFloat(float64) error { return nil }
func (b *builder) AssignString(string) error { return nil }
func (b *builder) AssignBytes([]byte) error { return nil }
func (b *builder) AssignLink(ipld.Link) error { return nil }
func (b *builder) AssignNode(ipld.Node) error { return nil }
func (b *builder) Prototype() ipld.NodePrototype { return nil }
