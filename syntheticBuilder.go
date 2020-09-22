package hamt

import "github.com/ipld/go-ipld-prime"

// TODO : the synthetic NodeBuilder + NodeAssembler that act like a single unified map go here

var _ ipld.NodeBuilder = (*builder)(nil)

type builder struct {
	hashAlg    string
	bucketSize int

	node *Node
}

func (b *builder) Build() ipld.Node { return b.node }
func (b *builder) Reset()           { b.node = nil }

func (b *builder) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	b.node = &Node{
		hashAlg:    b.hashAlg,
		bucketSize: b.bucketSize,
	}
	return assembler{&b.node._InteriorNode}, nil
}
func (b *builder) BeginList(sizeHint int) (ipld.ListAssembler, error) { return nil, nil }
func (b *builder) AssignNull() error                                  { return nil }
func (b *builder) AssignBool(bool) error                              { return nil }
func (b *builder) AssignInt(int) error                                { return nil }
func (b *builder) AssignFloat(float64) error                          { return nil }
func (b *builder) AssignString(string) error                          { return nil }
func (b *builder) AssignBytes([]byte) error                           { return nil }
func (b *builder) AssignLink(ipld.Link) error                         { return nil }
func (b *builder) AssignNode(ipld.Node) error                         { return nil }
func (b *builder) Prototype() ipld.NodePrototype                      { return nil }

type assembler struct {
	node *_InteriorNode
}

func (a assembler) AssembleKey() ipld.NodeAssembler {
	return nil
}
func (a assembler) AssembleValue() ipld.NodeAssembler {
	return nil
}
func (a assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	return nil, nil
}
func (a assembler) Finish() error {
	return nil
}
func (a assembler) KeyPrototype() ipld.NodePrototype {
	return _Bytes__Prototype{}
}
func (a assembler) ValuePrototype(k string) ipld.NodePrototype {
	return _Any__Prototype{}
}
