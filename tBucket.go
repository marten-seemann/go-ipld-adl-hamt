package hamt

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Bucket struct {
	x []_BucketEntry
}
type Bucket = *_Bucket
type _Bucket__Maybe struct {
	m schema.Maybe
	v Bucket
}
type MaybeBucket = *_Bucket__Maybe

func (m MaybeBucket) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeBucket) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeBucket) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeBucket) AsNode() ipld.Node {
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
func (m MaybeBucket) Must() Bucket {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Bucket)(&_Bucket{})
var _ schema.TypedNode = (Bucket)(&_Bucket{})

func (Bucket) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (Bucket) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"hamt.Bucket"}.LookupByString("")
}
func (n Bucket) LookupByNode(k ipld.Node) (ipld.Node, error) {
	idx, err := k.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(idx)
}
func (n Bucket) LookupByIndex(idx int) (ipld.Node, error) {
	if n.Length() <= idx {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(idx)}
	}
	v := &n.x[idx]
	return v, nil
}
func (n Bucket) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "hamt.Bucket", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (Bucket) MapIterator() ipld.MapIterator {
	return nil
}
func (n Bucket) ListIterator() ipld.ListIterator {
	return &_Bucket__ListItr{n, 0}
}

type _Bucket__ListItr struct {
	n   Bucket
	idx int
}

func (itr *_Bucket__ListItr) Next() (idx int, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.x) {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	x := &itr.n.x[itr.idx]
	v = x
	itr.idx++
	return
}
func (itr *_Bucket__ListItr) Done() bool {
	return itr.idx >= len(itr.n.x)
}

func (n Bucket) Length() int {
	return len(n.x)
}
func (Bucket) IsAbsent() bool {
	return false
}
func (Bucket) IsNull() bool {
	return false
}
func (Bucket) AsBool() (bool, error) {
	return mixins.List{"hamt.Bucket"}.AsBool()
}
func (Bucket) AsInt() (int, error) {
	return mixins.List{"hamt.Bucket"}.AsInt()
}
func (Bucket) AsFloat() (float64, error) {
	return mixins.List{"hamt.Bucket"}.AsFloat()
}
func (Bucket) AsString() (string, error) {
	return mixins.List{"hamt.Bucket"}.AsString()
}
func (Bucket) AsBytes() ([]byte, error) {
	return mixins.List{"hamt.Bucket"}.AsBytes()
}
func (Bucket) AsLink() (ipld.Link, error) {
	return mixins.List{"hamt.Bucket"}.AsLink()
}
func (Bucket) Prototype() ipld.NodePrototype {
	return _Bucket__Prototype{}
}

type _Bucket__Prototype struct{}

func (_Bucket__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Bucket__Builder
	nb.Reset()
	return &nb
}

type _Bucket__Builder struct {
	_Bucket__Assembler
}

func (nb *_Bucket__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Bucket__Builder) Reset() {
	var w _Bucket
	var m schema.Maybe
	*nb = _Bucket__Builder{_Bucket__Assembler{w: &w, m: &m}}
}

type _Bucket__Assembler struct {
	w     *_Bucket
	m     *schema.Maybe
	state laState

	cm schema.Maybe
	va _BucketEntry__Assembler
}

func (na *_Bucket__Assembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_Bucket__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"hamt.Bucket"}.BeginMap(0)
}
func (na *_Bucket__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_Bucket{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_BucketEntry, 0, sizeHint)
	}
	return na, nil
}
func (na *_Bucket__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"hamt.Bucket"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Bucket__Assembler) AssignBool(bool) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignBool(false)
}
func (_Bucket__Assembler) AssignInt(int) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignInt(0)
}
func (_Bucket__Assembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignFloat(0)
}
func (_Bucket__Assembler) AssignString(string) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignString("")
}
func (_Bucket__Assembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignBytes(nil)
}
func (_Bucket__Assembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"hamt.Bucket"}.AssignLink(nil)
}
func (na *_Bucket__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Bucket); ok {
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
		return ipld.ErrWrongKind{TypeName: "hamt.Bucket", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_Bucket__Assembler) Prototype() ipld.NodePrototype {
	return _Bucket__Prototype{}
}
func (la *_Bucket__Assembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_Bucket__Assembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _BucketEntry{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_Bucket__Assembler) Finish() error {
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
func (la *_Bucket__Assembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _BucketEntry__Prototype{}
}
func (Bucket) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Bucket) Representation() ipld.Node {
	return (*_Bucket__Repr)(n)
}

type _Bucket__Repr _Bucket

var _ ipld.Node = &_Bucket__Repr{}

func (_Bucket__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_Bucket__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"hamt.Bucket.Repr"}.LookupByString("")
}
func (nr *_Bucket__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Bucket)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(BucketEntry).Representation(), nil
}
func (nr *_Bucket__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	v, err := (Bucket)(nr).LookupByIndex(idx)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(BucketEntry).Representation(), nil
}
func (n _Bucket__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "hamt.Bucket.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_Bucket__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (nr *_Bucket__Repr) ListIterator() ipld.ListIterator {
	return &_Bucket__ReprListItr{(Bucket)(nr), 0}
}

type _Bucket__ReprListItr _Bucket__ListItr

func (itr *_Bucket__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	idx, v, err = (*_Bucket__ListItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return idx, v.(BucketEntry).Representation(), nil
}
func (itr *_Bucket__ReprListItr) Done() bool {
	return (*_Bucket__ListItr)(itr).Done()
}

func (rn *_Bucket__Repr) Length() int {
	return len(rn.x)
}
func (_Bucket__Repr) IsAbsent() bool {
	return false
}
func (_Bucket__Repr) IsNull() bool {
	return false
}
func (_Bucket__Repr) AsBool() (bool, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsBool()
}
func (_Bucket__Repr) AsInt() (int, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsInt()
}
func (_Bucket__Repr) AsFloat() (float64, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsFloat()
}
func (_Bucket__Repr) AsString() (string, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsString()
}
func (_Bucket__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsBytes()
}
func (_Bucket__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"hamt.Bucket.Repr"}.AsLink()
}
func (_Bucket__Repr) Prototype() ipld.NodePrototype {
	return _Bucket__ReprPrototype{}
}

type _Bucket__ReprPrototype struct{}

func (_Bucket__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Bucket__ReprBuilder
	nb.Reset()
	return &nb
}

type _Bucket__ReprBuilder struct {
	_Bucket__ReprAssembler
}

func (nb *_Bucket__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Bucket__ReprBuilder) Reset() {
	var w _Bucket
	var m schema.Maybe
	*nb = _Bucket__ReprBuilder{_Bucket__ReprAssembler{w: &w, m: &m}}
}

type _Bucket__ReprAssembler struct {
	w     *_Bucket
	m     *schema.Maybe
	state laState

	cm schema.Maybe
	va _BucketEntry__ReprAssembler
}

func (na *_Bucket__ReprAssembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_Bucket__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.BeginMap(0)
}
func (na *_Bucket__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_Bucket{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_BucketEntry, 0, sizeHint)
	}
	return na, nil
}
func (na *_Bucket__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"hamt.Bucket.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Bucket__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignBool(false)
}
func (_Bucket__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignInt(0)
}
func (_Bucket__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignFloat(0)
}
func (_Bucket__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignString("")
}
func (_Bucket__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignBytes(nil)
}
func (_Bucket__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"hamt.Bucket.Repr"}.AssignLink(nil)
}
func (na *_Bucket__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Bucket); ok {
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
		return ipld.ErrWrongKind{TypeName: "hamt.Bucket.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_Bucket__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Bucket__ReprPrototype{}
}
func (la *_Bucket__ReprAssembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_Bucket__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _BucketEntry{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_Bucket__ReprAssembler) Finish() error {
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
func (la *_Bucket__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _BucketEntry__ReprPrototype{}
}
