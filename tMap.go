package hamt

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Map struct {
	m map[_String]MaybeAny
	t []_Map__entry
}
type Map = *_Map
type _Map__entry struct {
	k _String
	v _Any__Maybe
}

func (n *_Map) LookupMaybe(k String) MaybeAny {
	v, ok := n.m[*k]
	if !ok {
		return &_Map__valueAbsent
	}
	return v
}

var _Map__valueAbsent = _Any__Maybe{m: schema.Maybe_Absent}

// TODO generate also a plain Lookup method that doesn't box and alloc if this type contains non-nullable values!
type _Map__Maybe struct {
	m schema.Maybe
	v Map
}
type MaybeMap = *_Map__Maybe

func (m MaybeMap) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMap) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMap) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMap) AsNode() ipld.Node {
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
func (m MaybeMap) Must() Map {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Map)(&_Map{})
var _ schema.TypedNode = (Map)(&_Map{})

func (Map) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Map) LookupByString(k string) (ipld.Node, error) {
	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	v, exists := n.m[k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k)}
	}
	if v.m == schema.Maybe_Null {
		return ipld.Null, nil
	}
	return v.v, nil
}
func (n Map) LookupByNode(k ipld.Node) (ipld.Node, error) {
	k2, ok := k.(String)
	if !ok {
		panic("todo invalid key type error")
		// 'ipld.ErrInvalidKey{TypeName:"hamt.Map", Key:&_String{k}}' doesn't quite cut it: need room to explain the type, and it's not guaranteed k can be turned into a string at all
	}
	v, exists := n.m[*k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k2.String())}
	}
	if v.m == schema.Maybe_Null {
		return ipld.Null, nil
	}
	return v.v, nil
}
func (Map) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.Map"}.LookupByIndex(0)
}
func (n Map) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n Map) MapIterator() ipld.MapIterator {
	return &_Map__MapItr{n, 0}
}

type _Map__MapItr struct {
	n   Map
	idx int
}

func (itr *_Map__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.t) {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	x := &itr.n.t[itr.idx]
	k = &x.k
	switch x.v.m {
	case schema.Maybe_Null:
		v = ipld.Null
	case schema.Maybe_Value:
		v = x.v.v
	}
	itr.idx++
	return
}
func (itr *_Map__MapItr) Done() bool {
	return itr.idx >= len(itr.n.t)
}

func (Map) ListIterator() ipld.ListIterator {
	return nil
}
func (n Map) Length() int {
	return len(n.t)
}
func (Map) IsAbsent() bool {
	return false
}
func (Map) IsNull() bool {
	return false
}
func (Map) AsBool() (bool, error) {
	return mixins.Map{"hamt.Map"}.AsBool()
}
func (Map) AsInt() (int, error) {
	return mixins.Map{"hamt.Map"}.AsInt()
}
func (Map) AsFloat() (float64, error) {
	return mixins.Map{"hamt.Map"}.AsFloat()
}
func (Map) AsString() (string, error) {
	return mixins.Map{"hamt.Map"}.AsString()
}
func (Map) AsBytes() ([]byte, error) {
	return mixins.Map{"hamt.Map"}.AsBytes()
}
func (Map) AsLink() (ipld.Link, error) {
	return mixins.Map{"hamt.Map"}.AsLink()
}
func (Map) Prototype() ipld.NodePrototype {
	return _Map__Prototype{}
}

type _Map__Prototype struct{}

func (_Map__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__Builder
	nb.Reset()
	return &nb
}

type _Map__Builder struct {
	_Map__Assembler
}

func (nb *_Map__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__Builder) Reset() {
	var w _Map
	var m schema.Maybe
	*nb = _Map__Builder{_Map__Assembler{w: &w, m: &m}}
}

type _Map__Assembler struct {
	w     *_Map
	m     *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__Assembler
	va _Any__Assembler
}

func (na *_Map__Assembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map{}
	}
	na.w.m = make(map[_String]MaybeAny, sizeHint)
	na.w.t = make([]_Map__entry, 0, sizeHint)
	return na, nil
}
func (_Map__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"hamt.Map"}.BeginList(0)
}
func (na *_Map__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"hamt.Map"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignBool(false)
}
func (_Map__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignInt(0)
}
func (_Map__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignFloat(0)
}
func (_Map__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignString("")
}
func (_Map__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignBytes(nil)
}
func (_Map__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"hamt.Map"}.AssignLink(nil)
}
func (na *_Map__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map); ok {
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
		return ipld.ErrWrongKind{TypeName: "hamt.Map", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_Map__Assembler) Prototype() ipld.NodePrototype {
	return _Map__Prototype{}
}
func (ma *_Map__Assembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.m = &tz.v.m
		tz.v.m = allowNull
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__Assembler) valueFinishTidy() bool {
	tz := &ma.w.t[len(ma.w.t)-1]
	switch tz.v.m {
	case schema.Maybe_Null:
		ma.state = maState_initial
		ma.va.reset()
		return true
	case schema.Maybe_Value:
		tz.v.v = ma.va.w
		ma.va.w = nil
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.m = &tz.v.m
	tz.v.m = allowNull
	return &ma.va, nil
}
func (ma *_Map__Assembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__Assembler) Finish() error {
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
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_Map__Assembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Any__Prototype{}
}
func (Map) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Map) Representation() ipld.Node {
	return (*_Map__Repr)(n)
}

type _Map__Repr _Map

var _ ipld.Node = &_Map__Repr{}

func (_Map__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (nr *_Map__Repr) LookupByString(k string) (ipld.Node, error) {
	v, err := (Map)(nr).LookupByString(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Any).Representation(), nil
}
func (nr *_Map__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Map)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Any).Representation(), nil
}
func (_Map__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.Map.Repr"}.LookupByIndex(0)
}
func (n _Map__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (nr *_Map__Repr) MapIterator() ipld.MapIterator {
	return &_Map__ReprMapItr{(Map)(nr), 0}
}

type _Map__ReprMapItr _Map__MapItr

func (itr *_Map__ReprMapItr) Next() (k ipld.Node, v ipld.Node, err error) {
	k, v, err = (*_Map__MapItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return k, v.(Any).Representation(), nil
}
func (itr *_Map__ReprMapItr) Done() bool {
	return (*_Map__MapItr)(itr).Done()
}

func (_Map__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Map__Repr) Length() int {
	return len(rn.t)
}
func (_Map__Repr) IsAbsent() bool {
	return false
}
func (_Map__Repr) IsNull() bool {
	return false
}
func (_Map__Repr) AsBool() (bool, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsBool()
}
func (_Map__Repr) AsInt() (int, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsInt()
}
func (_Map__Repr) AsFloat() (float64, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsFloat()
}
func (_Map__Repr) AsString() (string, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsString()
}
func (_Map__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsBytes()
}
func (_Map__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"hamt.Map.Repr"}.AsLink()
}
func (_Map__Repr) Prototype() ipld.NodePrototype {
	return _Map__ReprPrototype{}
}

type _Map__ReprPrototype struct{}

func (_Map__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__ReprBuilder
	nb.Reset()
	return &nb
}

type _Map__ReprBuilder struct {
	_Map__ReprAssembler
}

func (nb *_Map__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__ReprBuilder) Reset() {
	var w _Map
	var m schema.Maybe
	*nb = _Map__ReprBuilder{_Map__ReprAssembler{w: &w, m: &m}}
}

type _Map__ReprAssembler struct {
	w     *_Map
	m     *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__ReprAssembler
	va _Any__ReprAssembler
}

func (na *_Map__ReprAssembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map{}
	}
	na.w.m = make(map[_String]MaybeAny, sizeHint)
	na.w.t = make([]_Map__entry, 0, sizeHint)
	return na, nil
}
func (_Map__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"hamt.Map.Repr"}.BeginList(0)
}
func (na *_Map__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"hamt.Map.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignBool(false)
}
func (_Map__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignInt(0)
}
func (_Map__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignFloat(0)
}
func (_Map__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignString("")
}
func (_Map__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignBytes(nil)
}
func (_Map__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"hamt.Map.Repr"}.AssignLink(nil)
}
func (na *_Map__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map); ok {
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
		return ipld.ErrWrongKind{TypeName: "hamt.Map.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_Map__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Map__ReprPrototype{}
}
func (ma *_Map__ReprAssembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.m = &tz.v.m
		tz.v.m = allowNull
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ReprAssembler) valueFinishTidy() bool {
	tz := &ma.w.t[len(ma.w.t)-1]
	switch tz.v.m {
	case schema.Maybe_Null:
		ma.state = maState_initial
		ma.va.reset()
		return true
	case schema.Maybe_Value:
		tz.v.v = ma.va.w
		ma.va.w = nil
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _String
	if err := (_String__ReprPrototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.m = &tz.v.m
	tz.v.m = allowNull
	return &ma.va, nil
}
func (ma *_Map__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__ReprAssembler) Finish() error {
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
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__ReprPrototype{}
}
func (ma *_Map__ReprAssembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Any__ReprPrototype{}
}
