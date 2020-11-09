package hamt

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
	"math/bits"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/twmb/murmur3"
)

var _ ipld.Node = (*Node)(nil)

type Node struct {
	_HashMapRoot

	linkBuilder ipld.LinkBuilder
	linkLoader  ipld.Loader
	linkStorer  ipld.Storer
}

func (n Node) WithLinking(builder ipld.LinkBuilder, loader ipld.Loader, storer ipld.Storer) *Node {
	n.linkBuilder = builder
	n.linkLoader = loader
	n.linkStorer = storer
	return &n
}

func (n *Node) bitWidth() int {
	// bitWidth is inferred from the map length via the equation:
	//
	//     log2(byteLength(map) x 8)
	//
	// Since byteLength(map) is a power of 2, we don't need the expensive
	// float-based math.Log2; we can simply count the trailing zero bits.
	return bits.TrailingZeros32(uint32(len(n.hamt._map.x))) + 3
}

func (*Node) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}

func (n *Node) LookupByString(s string) (ipld.Node, error) {
	key := []byte(s)
	hash := n.hashKey(key)
	return lookupValue(&n.hamt, n.bitWidth(), 0, hash, key)
}

func (*Node) LookupByNode(ipld.Node) (ipld.Node, error) {
	panic("TODO")
}

func (*Node) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	panic("TODO")
}

func (*Node) MapIterator() ipld.MapIterator {
	panic("TODO")
}

func (n *Node) Length() int {
	count := 0
	for _, element := range n.hamt.data.x {
		switch element := element.x.(type) {
		case _Bucket:
			count += len(element.x)
		default:
			panic("todo")
		}
	}
	return count
}

func (*Node) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"hamt.Node"}.LookupByIndex(0)
}

func (*Node) ListIterator() ipld.ListIterator {
	return mixins.Map{"hamt.Node"}.ListIterator()
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

func (n *Node) hashKey(b []byte) []byte {
	var hasher hash.Hash
	switch n.hashAlg.x {
	case Identity:
		return b
	case Sha2_256:
		hasher = sha256.New()
	case Murmur3_128:
		hasher = murmur3.New128()
	default:
		// TODO: could we reach this? the builder already handles this
		// case, but other entry points like Reify don't.
		panic(fmt.Sprintf("unsupported hash algorithm: %x", n.hashAlg.x))
	}
	hasher.Write(b)
	return hasher.Sum(nil)
}

func insertEntry(node *_HashMapNode, bitWidth, bucketSize, depth int, hash []byte, entry _BucketEntry) error {
	from := depth * bitWidth
	index := rangedInt(hash, from, from+bitWidth)

	dataIndex := onesCountRange(node._map.x, index)
	exists := bitsetGet(node._map.x, index)
	if !exists {
		// Insert a new bucket at dataIndex.
		bucket := _Bucket{[]_BucketEntry{entry}}
		node.data.x = append(node.data.x[:dataIndex],
			append([]_Element{{bucket}}, node.data.x[dataIndex:]...)...)
		bitsetSet(node._map.x, index)
		return nil
	}
	switch element := node.data.x[dataIndex].x.(type) {
	case _Bucket:
		if len(element.x) < bucketSize {
			i, _ := lookupBucketEntry(element.x, entry.key.x)
			if i >= 0 {
				// Replace an existing key.
				element.x[i] = entry
			} else {
				// Add a new key.
				// TODO: keep the list sorted
				element.x = append(element.x, entry)
			}
		} else {
			panic("TODO")
		}
		node.data.x[dataIndex].x = element
	case _Link__HashMapNode:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unexpected element type: %T", element))
	}
	return nil
}

func lookupBucketEntry(entries []_BucketEntry, key []byte) (idx int, value _Any) {
	// TODO: to better support large buckets, should this be a
	// binary search?
	for i, entry := range entries {
		if bytes.Equal(entry.key.x, key) {
			return i, entry.value
		}
	}
	return -1, _Any{}
}

func lookupValue(node *_HashMapNode, bitWidth, depth int, hash, key []byte) (ipld.Node, error) {
	from := depth * bitWidth
	index := rangedInt(hash, from, from+bitWidth)

	exists := bitsetGet(node._map.x, index)
	if !exists {
		return nil, nil
	}
	dataIndex := onesCountRange(node._map.x, index)
	switch element := node.data.x[dataIndex].x.(type) {
	case _Bucket:
		i, value := lookupBucketEntry(element.x, key)
		if i >= 0 {
			return value.Representation(), nil
		}
	default:
		panic("TODO")
	}
	return nil, nil
}
