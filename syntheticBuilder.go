package hamt

import (
	"fmt"
	"os"

	"github.com/ipld/go-ipld-prime"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/twmb/murmur3"
)

const bitWidth = 3 // TODO: configurable?

// TODO : the synthetic NodeBuilder + NodeAssembler that act like a single unified map go here

var (
	_ = fmt.Sprint
	_ = os.Stdout
)

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
		_HashMapRoot: _HashMapRoot{
			hashAlg:    _String{b.hashAlg},
			bucketSize: _Int{b.bucketSize},
			_map:       _Bytes{make([]byte, 2<<bitWidth)},
		},
	}
	return &assembler{&b.node._HashMapRoot}, nil
}
func (b *builder) BeginList(sizeHint int) (ipld.ListAssembler, error) { panic("todo: error?") }
func (b *builder) AssignNull() error                                  { panic("todo: error?") }
func (b *builder) AssignBool(bool) error                              { panic("todo: error?") }
func (b *builder) AssignInt(int) error                                { panic("todo: error?") }
func (b *builder) AssignFloat(float64) error                          { panic("todo: error?") }
func (b *builder) AssignString(string) error                          { panic("todo: error?") }
func (b *builder) AssignBytes([]byte) error                           { panic("todo: error?") }
func (b *builder) AssignLink(ipld.Link) error                         { panic("todo: error?") }
func (b *builder) AssignNode(ipld.Node) error                         { panic("todo: error?") }
func (b *builder) Prototype() ipld.NodePrototype                      { panic("todo: error?") }

type assembler struct {
	node *_HashMapRoot
}

type assembleState uint8

const (
	stateInitial assembleState = iota
	stateMidKey
	stateExpectValue
	stateMidValue
	stateFinished
)

func (a *assembler) AssembleKey() ipld.NodeAssembler {
	return keyAssembler{a}
}

func (a *assembler) AssembleValue() ipld.NodeAssembler {
	return valueAssembler{a}
}

func (a *assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	return nil, nil
}

func (a *assembler) Finish() error {
	return nil
}

func (a *assembler) KeyPrototype() ipld.NodePrototype {
	return _Bytes__Prototype{}
}

func (a *assembler) ValuePrototype(k string) ipld.NodePrototype {
	return _Any__Prototype{}
}

type keyAssembler struct {
	parent *assembler
}

func (keyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.BytesAssembler{"bytes"}.BeginMap(0)
}

func (keyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.BytesAssembler{"bytes"}.BeginList(0)
}

func (keyAssembler) AssignNull() error {
	return mixins.BytesAssembler{"bytes"}.AssignNull()
}

func (keyAssembler) AssignBool(bool) error {
	return mixins.BytesAssembler{"bytes"}.AssignBool(false)
}

func (keyAssembler) AssignInt(int) error {
	return mixins.BytesAssembler{"bytes"}.AssignInt(0)
}

func (keyAssembler) AssignFloat(float64) error {
	return mixins.BytesAssembler{"bytes"}.AssignFloat(0)
}

func (a keyAssembler) AssignString(s string) error {
	return a.AssignBytes([]byte(s))
}

func (a keyAssembler) AssignBytes(b []byte) error {
	depth := 0

	hasher := murmur3.New128() // TODO: configurable
	hasher.Write(b)
	h := hasher.Sum(nil)

	from := depth * bitWidth
	index := rangedInt(h, from, from+bitWidth)
	// fmt.Fprintf(os.Stderr, "%x\n", h)
	// println("index:", index)

	node := a.parent.node
	dataIndex := onesCountRange(node._map.x, index)
	exists := bitsetGet(node._map.x, index)
	if !exists {
		bucket := _Bucket{[]_BucketEntry{{
			_Bytes{b},
			_Value{},
		}}}
		node.data.x = append(node.data.x[:dataIndex], append([]_Element{{bucket}}, node.data.x[dataIndex:]...)...)
	} else {
	}

	return nil
}

func (keyAssembler) AssignLink(ipld.Link) error {
	return mixins.BytesAssembler{"bytes"}.AssignLink(nil)
}

func (a keyAssembler) AssignNode(v ipld.Node) error {
	vs, err := v.AsString()
	if err != nil {
		return err
	}
	return a.AssignString(vs)
}

func (keyAssembler) Prototype() ipld.NodePrototype {
	return basicnode.Prototype__Bytes{}
}

type valueAssembler struct {
	parent *assembler
}

func (valueAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.BytesAssembler{"bytes"}.BeginMap(0)
}

func (valueAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.BytesAssembler{"bytes"}.BeginList(0)
}

func (valueAssembler) AssignNull() error {
	return mixins.BytesAssembler{"bytes"}.AssignNull()
}

func (valueAssembler) AssignBool(bool) error {
	return mixins.BytesAssembler{"bytes"}.AssignBool(false)
}

func (valueAssembler) AssignInt(int) error {
	return mixins.BytesAssembler{"bytes"}.AssignInt(0)
}

func (valueAssembler) AssignFloat(float64) error {
	return mixins.BytesAssembler{"bytes"}.AssignFloat(0)
}

func (a valueAssembler) AssignString(s string) error {
	return a.AssignBytes([]byte(s))
}

func (a valueAssembler) AssignBytes(b []byte) error {
	// TODO do work
	return nil
}

func (valueAssembler) AssignLink(ipld.Link) error {
	return mixins.BytesAssembler{"bytes"}.AssignLink(nil)
}

func (a valueAssembler) AssignNode(v ipld.Node) error {
	vs, err := v.AsString()
	if err != nil {
		return err
	}
	return a.AssignString(vs)
}

func (valueAssembler) Prototype() ipld.NodePrototype {
	return basicnode.Prototype__Bytes{}
}
