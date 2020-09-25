package hamt

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
)

// TODO : the synthetic Node that acts like a single unified map goes here, and so does its Prototype and whatever Config that entails.

var _ ipld.Node = (*Node)(nil)

type Node struct {
	_HashMapRoot
}

func (*Node) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (*Node) LookupByString(string) (ipld.Node, error) {
	panic("todo")
}
func (*Node) LookupByNode(ipld.Node) (ipld.Node, error) {
	panic("todo")
}
func (*Node) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.Node"}.LookupByIndex(0)
}
func (*Node) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	panic("todo")
}
func (*Node) MapIterator() ipld.MapIterator {
	panic("todo")
}
func (*Node) ListIterator() ipld.ListIterator {
	return nil
}
func (*Node) Length() int {
	return -1
}
func (*Node) IsAbsent() bool {
	return false
}
func (*Node) IsNull() bool {
	return false
}
func (*Node) AsBool() (bool, error) {
	return mixins.Map{"hamt.Node"}.AsBool()
}
func (*Node) AsInt() (int, error) {
	return mixins.Map{"hamt.Node"}.AsInt()
}
func (*Node) AsFloat() (float64, error) {
	return mixins.Map{"hamt.Node"}.AsFloat()
}
func (*Node) AsString() (string, error) {
	return mixins.Map{"hamt.Node"}.AsString()
}
func (*Node) AsBytes() ([]byte, error) {
	return mixins.Map{"hamt.Node"}.AsBytes()
}
func (*Node) AsLink() (ipld.Link, error) {
	return mixins.Map{"hamt.Node"}.AsLink()
}

var _ ipld.NodePrototype = (*Prototype)(nil)

type Prototype struct {
}

func (Prototype) NewBuilder() ipld.NodeBuilder {
	return &builder{}
}
