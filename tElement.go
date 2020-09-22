package hamt

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Element struct {
	tag uint
	x1  _HashMapNode
	x2  _Link__HashMapNode
	x3  _Bucket
}
type Element = *_Element

type _Element__iface interface {
	_Element__member()
}

func (_HashMapNode) _Element__member()       {}
func (_Link__HashMapNode) _Element__member() {}
func (_Bucket) _Element__member()            {}
func (n _Element) AsInterface() _Element__iface {
	switch n.tag {
	case 1:
		return &n.x1
	case 2:
		return &n.x2
	case 3:
		return &n.x3
	default:
		panic("invalid union state; how did you create this object?")
	}
}

type _Element__Maybe struct {
	m schema.Maybe
	v Element
}
type MaybeElement = *_Element__Maybe

func (m MaybeElement) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeElement) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeElement) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeElement) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Absent
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeElement) Must() Element {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var (
	memberName__Element_HashMapNode       = _String{"HashMapNode"}
	memberName__Element_Link__HashMapNode = _String{"Link__HashMapNode"}
	memberName__Element_Bucket            = _String{"Bucket"}
)
var _ ipld.Node = (Element)(&_Element{})
var _ schema.TypedNode = (Element)(&_Element{})

func (Element) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Element) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "HashMapNode":
		if n.tag != 1 {
			return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return &n.x1, nil
	case "Link__HashMapNode":
		if n.tag != 2 {
			return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return &n.x2, nil
	case "Bucket":
		if n.tag != 3 {
			return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return &n.x3, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n Element) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (Element) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.Element"}.LookupByIndex(0)
}
func (n Element) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n Element) MapIterator() ipld.MapIterator {
	return &_Element__MapItr{n, false}
}

type _Element__MapItr struct {
	n    Element
	done bool
}

func (itr *_Element__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.done {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.n.tag {
	case 1:
		k, v = &memberName__Element_HashMapNode, &itr.n.x1
	case 2:
		k, v = &memberName__Element_Link__HashMapNode, &itr.n.x2
	case 3:
		k, v = &memberName__Element_Bucket, &itr.n.x3
	default:
		panic("unreachable")
	}
	itr.done = true
	return
}
func (itr *_Element__MapItr) Done() bool {
	return itr.done
}

func (Element) ListIterator() ipld.ListIterator {
	return nil
}
func (Element) Length() int {
	return 1
}
func (Element) IsAbsent() bool {
	return false
}
func (Element) IsNull() bool {
	return false
}
func (Element) AsBool() (bool, error) {
	return mixins.Map{"hamt.Element"}.AsBool()
}
func (Element) AsInt() (int, error) {
	return mixins.Map{"hamt.Element"}.AsInt()
}
func (Element) AsFloat() (float64, error) {
	return mixins.Map{"hamt.Element"}.AsFloat()
}
func (Element) AsString() (string, error) {
	return mixins.Map{"hamt.Element"}.AsString()
}
func (Element) AsBytes() ([]byte, error) {
	return mixins.Map{"hamt.Element"}.AsBytes()
}
func (Element) AsLink() (ipld.Link, error) {
	return mixins.Map{"hamt.Element"}.AsLink()
}
func (Element) Prototype() ipld.NodePrototype {
	return _Element__Prototype{}
}

type _Element__Prototype struct{}

func (_Element__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Element__Builder
	nb.Reset()
	return &nb
}

type _Element__Builder struct {
	_Element__Assembler
}

func (nb *_Element__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Element__Builder) Reset() {
	var w _Element
	var m schema.Maybe
	*nb = _Element__Builder{_Element__Assembler{w: &w, m: &m}}
}

type _Element__Assembler struct {
	w     *_Element
	m     *schema.Maybe
	state maState

	cm  schema.Maybe
	ca1 _HashMapNode__Assembler

	ca2 _Link__HashMapNode__Assembler

	ca3 _Bucket__Assembler
	ca  uint
}

func (na *_Element__Assembler) reset() {
	na.state = maState_initial
	switch na.ca {
	case 0:
		return
	case 1:
		na.ca1.reset()

	case 2:
		na.ca2.reset()

	case 3:
		na.ca3.reset()
	default:
		panic("unreachable")
	}
	na.ca = 0
	na.cm = schema.Maybe_Absent
}
func (na *_Element__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_Element{}
	}
	return na, nil
}
func (_Element__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"hamt.Element"}.BeginList(0)
}
func (na *_Element__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"hamt.Element"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Element__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignBool(false)
}
func (_Element__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignInt(0)
}
func (_Element__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignFloat(0)
}
func (_Element__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignString("")
}
func (_Element__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignBytes(nil)
}
func (_Element__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"hamt.Element"}.AssignLink(nil)
}
func (na *_Element__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Element); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "hamt.Element", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Element__Assembler) Prototype() ipld.NodePrototype {
	return _Element__Prototype{}
}
func (ma *_Element__Assembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.state = maState_initial
		return true
	default:
		return false
	}
}
func (ma *_Element__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on for the moment, but we'll still be erroring shortly.
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	if ma.ca != 0 {
		return nil, schema.ErrNotUnionStructure{TypeName: "hamt.Element", Detail: "cannot add another entry -- a union can only contain one thing!"}
	}
	switch k {
	case "HashMapNode":
		ma.state = maState_midValue
		ma.ca = 1
		ma.w.tag = 1
		ma.ca1.w = &ma.w.x1
		ma.ca1.m = &ma.cm
		return &ma.ca1, nil
	case "Link__HashMapNode":
		ma.state = maState_midValue
		ma.ca = 2
		ma.w.tag = 2
		ma.ca2.w = &ma.w.x2
		ma.ca2.m = &ma.cm
		return &ma.ca2, nil
	case "Bucket":
		ma.state = maState_midValue
		ma.ca = 3
		ma.w.tag = 3
		ma.ca3.w = &ma.w.x3
		ma.ca3.m = &ma.cm
		return &ma.ca3, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName: "hamt.Element", Key: &_String{k}}
	}
}
func (ma *_Element__Assembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on for the moment, but we'll still be erroring shortly... or rather, the keyassembler will be.
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_Element__KeyAssembler)(ma)
}
func (ma *_Element__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.ca {
	case 0:
		ma.ca1.w = &ma.w.x1
		ma.ca1.m = &ma.cm
		return &ma.ca1
	case 1:
		ma.ca2.w = &ma.w.x2
		ma.ca2.m = &ma.cm
		return &ma.ca2
	case 2:
		ma.ca3.w = &ma.w.x3
		ma.ca3.m = &ma.cm
		return &ma.ca3
	default:
		panic("unreachable")
	}
}
func (ma *_Element__Assembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	if ma.ca == 0 {
		return schema.ErrNotUnionStructure{TypeName: "hamt.Element", Detail: "a union must have exactly one entry (not none)!"}
	}
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Element__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_Element__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	switch k {
	case "HashMapNode":
		return _HashMapNode__Prototype{}
	case "Link__HashMapNode":
		return _Link__HashMapNode__Prototype{}
	case "Bucket":
		return _Bucket__Prototype{}
	default:
		return nil
	}
}

type _Element__KeyAssembler _Element__Assembler

func (_Element__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.BeginMap(0)
}
func (_Element__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.BeginList(0)
}
func (na *_Element__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignNull()
}
func (_Element__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignBool(false)
}
func (_Element__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignInt(0)
}
func (_Element__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignFloat(0)
}
func (ka *_Element__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	if ka.ca != 0 {
		return schema.ErrNotUnionStructure{TypeName: "hamt.Element", Detail: "cannot add another entry -- a union can only contain one thing!"}
	}
	switch k {
	case "HashMapNode":
		ka.ca = 1
		ka.w.tag = 1
		ka.state = maState_expectValue
		return nil
	case "Link__HashMapNode":
		ka.ca = 2
		ka.w.tag = 2
		ka.state = maState_expectValue
		return nil
	case "Bucket":
		ka.ca = 3
		ka.w.tag = 3
		ka.state = maState_expectValue
		return nil
	default:
		return ipld.ErrInvalidKey{TypeName: "hamt.Element", Key: &_String{k}} // TODO: error quality: ErrInvalidUnionDiscriminant ?
	}
	return nil
}
func (_Element__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignBytes(nil)
}
func (_Element__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"hamt.Element.KeyAssembler"}.AssignLink(nil)
}
func (ka *_Element__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_Element__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (Element) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Element) Representation() ipld.Node {
	return (*_Element__Repr)(n)
}

type _Element__Repr _Element

var _ ipld.Node = &_Element__Repr{}

func (n *_Element__Repr) ReprKind() ipld.ReprKind {
	switch n.tag {
	case 1:
		return ipld.ReprKind_Map
	case 2:
		return ipld.ReprKind_Link
	case 3:
		return ipld.ReprKind_List
	default:
		panic("unreachable")
	}
}
func (n *_Element__Repr) LookupByString(key string) (ipld.Node, error) {
	switch n.tag {
	case 1:
		return n.x1.Representation().LookupByString(key)
	default:
		return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "LookupByString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: n.ReprKind()}
	}
}
func (n *_Element__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	switch n.tag {
	case 1:
		return n.x1.Representation().LookupByNode(key)
	case 3:
		return n.x3.Representation().LookupByNode(key)
	default:
		return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "LookupByNode", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: n.ReprKind()}
	}
}
func (n *_Element__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch n.tag {
	case 3:
		return n.x3.Representation().LookupByIndex(idx)
	default:
		return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "LookupByIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: n.ReprKind()}
	}
}
func (n *_Element__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	switch n.tag {
	case 1:
		return n.x1.Representation().LookupBySegment(seg)
	case 3:
		return n.x3.Representation().LookupBySegment(seg)
	default:
		return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "LookupBySegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: n.ReprKind()}
	}
}
func (n *_Element__Repr) MapIterator() ipld.MapIterator {
	switch n.tag {
	case 1:
		return n.x1.Representation().MapIterator()
	default:
		return nil
	}
}
func (n *_Element__Repr) ListIterator() ipld.ListIterator {
	switch n.tag {
	case 3:
		return n.x3.Representation().ListIterator()
	default:
		return nil
	}
}
func (n *_Element__Repr) Length() int {
	switch n.tag {
	case 1:
		return n.x1.Representation().Length()
	case 3:
		return n.x3.Representation().Length()
	default:
		return -1
	}
}
func (n *_Element__Repr) IsAbsent() bool {
	return false
}
func (n *_Element__Repr) IsNull() bool {
	return false
}
func (n *_Element__Repr) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: n.ReprKind()}
}
func (n *_Element__Repr) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: n.ReprKind()}
}
func (n *_Element__Repr) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: n.ReprKind()}
}
func (n *_Element__Repr) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: n.ReprKind()}
}
func (n *_Element__Repr) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: n.ReprKind()}
}
func (n *_Element__Repr) AsLink() (ipld.Link, error) {
	switch n.tag {
	case 2:
		return n.x2.Representation().AsLink()
	default:
		return nil, ipld.ErrWrongKind{TypeName: "hamt.Element.Repr", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: n.ReprKind()}
	}
}
func (_Element__Repr) Prototype() ipld.NodePrototype {
	return _Element__ReprPrototype{}
}

type _Element__ReprPrototype struct{}

func (_Element__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Element__ReprBuilder
	nb.Reset()
	return &nb
}

type _Element__ReprBuilder struct {
	_Element__ReprAssembler
}

func (nb *_Element__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Element__ReprBuilder) Reset() {
	var w _Element
	var m schema.Maybe
	*nb = _Element__ReprBuilder{_Element__ReprAssembler{w: &w, m: &m}}
}

type _Element__ReprAssembler struct {
	w   *_Element
	m   *schema.Maybe
	ca1 _HashMapNode__ReprAssembler
	ca2 _Link__HashMapNode__ReprAssembler
	ca3 _Bucket__ReprAssembler
	ca  uint
}

func (na *_Element__ReprAssembler) reset() {
	switch na.ca {
	case 0:
		return
	case 1:
		na.ca1.reset()
	case 2:
		na.ca2.reset()
	case 3:
		na.ca3.reset()
	default:
		panic("unreachable")
	}
	na.ca = 0
}
func (na *_Element__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	if na.w == nil {
		na.w = &_Element{}
	}
	na.ca = 1
	na.w.tag = 1
	na.ca1.w = &na.w.x1
	na.ca1.m = na.m
	return na.ca1.BeginMap(sizeHint)
	return nil, schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "BeginMap called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	if na.w == nil {
		na.w = &_Element{}
	}
	na.ca = 3
	na.w.tag = 3
	na.ca3.w = &na.w.x3
	na.ca3.m = na.m
	return na.ca3.BeginList(sizeHint)
	return nil, schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "BeginList called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignNull() error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignNull called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignBool(v bool) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignBool called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignInt(v int) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignInt called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignFloat(v float64) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignFloat called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignString(v string) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignString called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignBytes(v []byte) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignBytes called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignLink(v ipld.Link) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign into assembler that's already working on a larger structure!")
	}
	if na.w == nil {
		na.w = &_Element{}
	}
	na.ca = 2
	na.w.tag = 2
	na.ca2.w = &na.w.x2
	na.ca2.m = na.m
	return na.ca2.AssignLink(v)
	return schema.ErrNotUnionStructure{TypeName: "hamt.Element.Repr", Detail: "AssignLink called but is not valid for any of the kinds that are valid members of this union"}
}
func (na *_Element__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Element); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	switch v.ReprKind() {
	case ipld.ReprKind_Bool:
		v2, _ := v.AsBool()
		return na.AssignBool(v2)
	case ipld.ReprKind_Int:
		v2, _ := v.AsInt()
		return na.AssignInt(v2)
	case ipld.ReprKind_Float:
		v2, _ := v.AsFloat()
		return na.AssignFloat(v2)
	case ipld.ReprKind_String:
		v2, _ := v.AsString()
		return na.AssignString(v2)
	case ipld.ReprKind_Bytes:
		v2, _ := v.AsBytes()
		return na.AssignBytes(v2)
	case ipld.ReprKind_Map:
		na, err := na.BeginMap(v.Length())
		if err != nil {
			return err
		}
		itr := v.MapIterator()
		for !itr.Done() {
			k, v, err := itr.Next()
			if err != nil {
				return err
			}
			if err := na.AssembleKey().AssignNode(k); err != nil {
				return err
			}
			if err := na.AssembleValue().AssignNode(v); err != nil {
				return err
			}
		}
		return na.Finish()
	case ipld.ReprKind_List:
		na, err := na.BeginList(v.Length())
		if err != nil {
			return err
		}
		itr := v.ListIterator()
		for !itr.Done() {
			_, v, err := itr.Next()
			if err != nil {
				return err
			}
			if err := na.AssembleValue().AssignNode(v); err != nil {
				return err
			}
		}
		return na.Finish()
	case ipld.ReprKind_Link:
		v2, _ := v.AsLink()
		return na.AssignLink(v2)
	default:
		panic("unreachable")
	}
}
func (na *_Element__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Element__ReprPrototype{}
}
