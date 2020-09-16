package hamt

import "github.com/ipld/go-ipld-prime"

// TODO : the synthetic Node that acts like a single unified map goes here, and so does its Prototype and whatever Config that entails.

// var _ ipld.Node = (*Node)(nil)

type Node struct {
	// TODO(mvdan): shouldn't these be in the schema?
	hashAlg    string
	bucketSize int

	_InteriorNode
}

var _ ipld.NodePrototype = (*Prototype)(nil)

type Prototype struct {
}

func (Prototype) NewBuilder() ipld.NodeBuilder {
	return &builder{}
}
