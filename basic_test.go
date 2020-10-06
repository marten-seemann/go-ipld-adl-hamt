package hamt

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBasic(t *testing.T) {
	t.Parallel()

	builder := Prototype{}.NewBuilder()
	assembler, err := builder.BeginMap(0)
	qt.Assert(t, err, qt.IsNil)

	assembler.AssembleKey().AssignString("foo")
	assembler.AssembleValue().AssignString("bar")
	assembler.Finish()

	node := builder.Build()

	qt.Assert(t, node.Length(), qt.Equals, 1)
}
