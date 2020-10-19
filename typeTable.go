package hamt

// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		hamt.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		hamt.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	Any                     _Any__Prototype
	Any__Repr               _Any__ReprPrototype
	Bool                    _Bool__Prototype
	Bool__Repr              _Bool__ReprPrototype
	Bucket                  _Bucket__Prototype
	Bucket__Repr            _Bucket__ReprPrototype
	BucketEntry             _BucketEntry__Prototype
	BucketEntry__Repr       _BucketEntry__ReprPrototype
	Bytes                   _Bytes__Prototype
	Bytes__Repr             _Bytes__ReprPrototype
	Element                 _Element__Prototype
	Element__Repr           _Element__ReprPrototype
	Float                   _Float__Prototype
	Float__Repr             _Float__ReprPrototype
	HashMapNode             _HashMapNode__Prototype
	HashMapNode__Repr       _HashMapNode__ReprPrototype
	HashMapRoot             _HashMapRoot__Prototype
	HashMapRoot__Repr       _HashMapRoot__ReprPrototype
	Int                     _Int__Prototype
	Int__Repr               _Int__ReprPrototype
	Link                    _Link__Prototype
	Link__Repr              _Link__ReprPrototype
	Link__HashMapNode       _Link__HashMapNode__Prototype
	Link__HashMapNode__Repr _Link__HashMapNode__ReprPrototype
	List                    _List__Prototype
	List__Repr              _List__ReprPrototype
	List__Element           _List__Element__Prototype
	List__Element__Repr     _List__Element__ReprPrototype
	Map                     _Map__Prototype
	Map__Repr               _Map__ReprPrototype
	String                  _String__Prototype
	String__Repr            _String__ReprPrototype
}
