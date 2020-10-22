package hamt

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _HashMapNode struct {
	_map _Bytes
	data _List__Element
}
type HashMapNode = *_HashMapNode

func (n _HashMapNode) FieldMap() Bytes {
	return &n._map
}
func (n _HashMapNode) FieldData() List__Element {
	return &n.data
}

type _HashMapNode__Maybe struct {
	m schema.Maybe
	v HashMapNode
}
type MaybeHashMapNode = *_HashMapNode__Maybe

func (m MaybeHashMapNode) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeHashMapNode) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeHashMapNode) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeHashMapNode) AsNode() ipld.Node {
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
func (m MaybeHashMapNode) Must() HashMapNode {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var (
	fieldName__HashMapNode_Map  = _String{"map"}
	fieldName__HashMapNode_Data = _String{"data"}
)
var _ ipld.Node = (HashMapNode)(&_HashMapNode{})
var _ schema.TypedNode = (HashMapNode)(&_HashMapNode{})

func (HashMapNode) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n HashMapNode) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "map":
		return &n._map, nil
	case "data":
		return &n.data, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n HashMapNode) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (HashMapNode) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.HashMapNode"}.LookupByIndex(0)
}
func (n HashMapNode) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n HashMapNode) MapIterator() ipld.MapIterator {
	return &_HashMapNode__MapItr{n, 0}
}

type _HashMapNode__MapItr struct {
	n   HashMapNode
	idx int
}

func (itr *_HashMapNode__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__HashMapNode_Map
		v = &itr.n._map
	case 1:
		k = &fieldName__HashMapNode_Data
		v = &itr.n.data
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_HashMapNode__MapItr) Done() bool {
	return itr.idx >= 2
}

func (HashMapNode) ListIterator() ipld.ListIterator {
	return nil
}
func (HashMapNode) Length() int {
	return 2
}
func (HashMapNode) IsAbsent() bool {
	return false
}
func (HashMapNode) IsNull() bool {
	return false
}
func (HashMapNode) AsBool() (bool, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsBool()
}
func (HashMapNode) AsInt() (int, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsInt()
}
func (HashMapNode) AsFloat() (float64, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsFloat()
}
func (HashMapNode) AsString() (string, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsString()
}
func (HashMapNode) AsBytes() ([]byte, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsBytes()
}
func (HashMapNode) AsLink() (ipld.Link, error) {
	return mixins.Map{"hamt.HashMapNode"}.AsLink()
}
func (HashMapNode) Prototype() ipld.NodePrototype {
	return _HashMapNode__Prototype{}
}

type _HashMapNode__Prototype struct{}

func (_HashMapNode__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _HashMapNode__Builder
	nb.Reset()
	return &nb
}

type _HashMapNode__Builder struct {
	_HashMapNode__Assembler
}

func (nb *_HashMapNode__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_HashMapNode__Builder) Reset() {
	var w _HashMapNode
	var m schema.Maybe
	*nb = _HashMapNode__Builder{_HashMapNode__Assembler{w: &w, m: &m}}
}

type _HashMapNode__Assembler struct {
	w     *_HashMapNode
	m     *schema.Maybe
	state maState
	s     int
	f     int

	cm      schema.Maybe
	ca__map _Bytes__Assembler
	ca_data _List__Element__Assembler
}

func (na *_HashMapNode__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca__map.reset()
	na.ca_data.reset()
}

var (
	fieldBit__HashMapNode_Map         = 1 << 0
	fieldBit__HashMapNode_Data        = 1 << 1
	fieldBits__HashMapNode_sufficient = 0 + 1<<0 + 1<<1
)

func (na *_HashMapNode__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_HashMapNode{}
	}
	return na, nil
}
func (_HashMapNode__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"hamt.HashMapNode"}.BeginList(0)
}
func (na *_HashMapNode__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"hamt.HashMapNode"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_HashMapNode__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignBool(false)
}
func (_HashMapNode__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignInt(0)
}
func (_HashMapNode__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignFloat(0)
}
func (_HashMapNode__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignString("")
}
func (_HashMapNode__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignBytes(nil)
}
func (_HashMapNode__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"hamt.HashMapNode"}.AssignLink(nil)
}
func (na *_HashMapNode__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_HashMapNode); ok {
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
		return ipld.ErrWrongKind{TypeName: "hamt.HashMapNode", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_HashMapNode__Assembler) Prototype() ipld.NodePrototype {
	return _HashMapNode__Prototype{}
}
func (ma *_HashMapNode__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca__map.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_data.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_HashMapNode__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	switch k {
	case "map":
		if ma.s&fieldBit__HashMapNode_Map != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__HashMapNode_Map}
		}
		ma.s += fieldBit__HashMapNode_Map
		ma.state = maState_midValue
		ma.f = 0
		ma.ca__map.w = &ma.w._map
		ma.ca__map.m = &ma.cm
		return &ma.ca__map, nil
	case "data":
		if ma.s&fieldBit__HashMapNode_Data != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__HashMapNode_Data}
		}
		ma.s += fieldBit__HashMapNode_Data
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_data.w = &ma.w.data
		ma.ca_data.m = &ma.cm
		return &ma.ca_data, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName: "hamt.HashMapNode", Key: &_String{k}}
	}
}
func (ma *_HashMapNode__Assembler) AssembleKey() ipld.NodeAssembler {
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
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_HashMapNode__KeyAssembler)(ma)
}
func (ma *_HashMapNode__Assembler) AssembleValue() ipld.NodeAssembler {
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
	switch ma.f {
	case 0:
		ma.ca__map.w = &ma.w._map
		ma.ca__map.m = &ma.cm
		return &ma.ca__map
	case 1:
		ma.ca_data.w = &ma.w.data
		ma.ca_data.m = &ma.cm
		return &ma.ca_data
	default:
		panic("unreachable")
	}
}
func (ma *_HashMapNode__Assembler) Finish() error {
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
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_HashMapNode__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_HashMapNode__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}

type _HashMapNode__KeyAssembler _HashMapNode__Assembler

func (_HashMapNode__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.BeginMap(0)
}
func (_HashMapNode__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.BeginList(0)
}
func (na *_HashMapNode__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignNull()
}
func (_HashMapNode__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignBool(false)
}
func (_HashMapNode__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignInt(0)
}
func (_HashMapNode__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignFloat(0)
}
func (ka *_HashMapNode__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "map":
		if ka.s&fieldBit__HashMapNode_Map != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__HashMapNode_Map}
		}
		ka.s += fieldBit__HashMapNode_Map
		ka.state = maState_expectValue
		ka.f = 0
	case "data":
		if ka.s&fieldBit__HashMapNode_Data != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__HashMapNode_Data}
		}
		ka.s += fieldBit__HashMapNode_Data
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName: "hamt.HashMapNode", Key: &_String{k}}
	}
	return nil
}
func (_HashMapNode__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignBytes(nil)
}
func (_HashMapNode__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"hamt.HashMapNode.KeyAssembler"}.AssignLink(nil)
}
func (ka *_HashMapNode__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_HashMapNode__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (HashMapNode) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n HashMapNode) Representation() ipld.Node {
	return (*_HashMapNode__Repr)(n)
}

type _HashMapNode__Repr _HashMapNode

var _ ipld.Node = &_HashMapNode__Repr{}

func (_HashMapNode__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_HashMapNode__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.LookupByString("")
}
func (n *_HashMapNode__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_HashMapNode__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n._map.Representation(), nil
	case 1:
		return n.data.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _HashMapNode__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "hamt.HashMapNode.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_HashMapNode__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_HashMapNode__Repr) ListIterator() ipld.ListIterator {
	return &_HashMapNode__ReprListItr{n, 0}
}

type _HashMapNode__ReprListItr struct {
	n   *_HashMapNode__Repr
	idx int
}

func (itr *_HashMapNode__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 2 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n._map.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.data.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_HashMapNode__ReprListItr) Done() bool {
	return itr.idx >= 2
}

func (rn *_HashMapNode__Repr) Length() int {
	l := 2
	return l
}
func (_HashMapNode__Repr) IsAbsent() bool {
	return false
}
func (_HashMapNode__Repr) IsNull() bool {
	return false
}
func (_HashMapNode__Repr) AsBool() (bool, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsBool()
}
func (_HashMapNode__Repr) AsInt() (int, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsInt()
}
func (_HashMapNode__Repr) AsFloat() (float64, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsFloat()
}
func (_HashMapNode__Repr) AsString() (string, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsString()
}
func (_HashMapNode__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsBytes()
}
func (_HashMapNode__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"hamt.HashMapNode.Repr"}.AsLink()
}
func (_HashMapNode__Repr) Prototype() ipld.NodePrototype {
	return _HashMapNode__ReprPrototype{}
}

type _HashMapNode__ReprPrototype struct{}

func (_HashMapNode__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _HashMapNode__ReprBuilder
	nb.Reset()
	return &nb
}

type _HashMapNode__ReprBuilder struct {
	_HashMapNode__ReprAssembler
}

func (nb *_HashMapNode__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_HashMapNode__ReprBuilder) Reset() {
	var w _HashMapNode
	var m schema.Maybe
	*nb = _HashMapNode__ReprBuilder{_HashMapNode__ReprAssembler{w: &w, m: &m}}
}

type _HashMapNode__ReprAssembler struct {
	w     *_HashMapNode
	m     *schema.Maybe
	state laState
	f     int

	cm      schema.Maybe
	ca__map _Bytes__ReprAssembler
	ca_data _List__Element__ReprAssembler
}

func (na *_HashMapNode__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca__map.reset()
	na.ca_data.reset()
}
func (_HashMapNode__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.BeginMap(0)
}
func (na *_HashMapNode__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_HashMapNode{}
	}
	return na, nil
}
func (na *_HashMapNode__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"hamt.HashMapNode.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_HashMapNode__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignBool(false)
}
func (_HashMapNode__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignInt(0)
}
func (_HashMapNode__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignFloat(0)
}
func (_HashMapNode__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignString("")
}
func (_HashMapNode__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignBytes(nil)
}
func (_HashMapNode__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"hamt.HashMapNode.Repr"}.AssignLink(nil)
}
func (na *_HashMapNode__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_HashMapNode); ok {
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
	if v.ReprKind() != ipld.ReprKind_List {
		return ipld.ErrWrongKind{TypeName: "hamt.HashMapNode.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
}
func (_HashMapNode__ReprAssembler) Prototype() ipld.NodePrototype {
	return _HashMapNode__ReprPrototype{}
}
func (la *_HashMapNode__ReprAssembler) valueFinishTidy() bool {
	switch la.f {
	case 0:
		switch la.cm {
		case schema.Maybe_Value:
			la.cm = schema.Maybe_Absent
			la.state = laState_initial
			la.f++
			return true
		default:
			return false
		}
	case 1:
		switch la.cm {
		case schema.Maybe_Value:
			la.cm = schema.Maybe_Absent
			la.state = laState_initial
			la.f++
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (la *_HashMapNode__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when still in the middle of assembling the previous value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	if la.f >= 2 {
		return nil // schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(2)} // FIXME: need an error thunking assembler!  it has returned.  sigh.
	}
	la.state = laState_midValue
	switch la.f {
	case 0:
		la.ca__map.w = &la.w._map
		la.ca__map.m = &la.cm
		return &la.ca__map
	case 1:
		la.ca_data.w = &la.w.data
		la.ca_data.m = &la.cm
		return &la.ca_data
	default:
		panic("unreachable")
	}
}
func (la *_HashMapNode__ReprAssembler) Finish() error {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	la.state = laState_finished
	*la.m = schema.Maybe_Value
	return nil
}
func (la *_HashMapNode__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}
