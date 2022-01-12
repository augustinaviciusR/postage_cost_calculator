package mutator

import (
	"github.com/zimmski/go-mutesting/test"
	"testing"
)

func TestMutatorCase(t *testing.T) {
	test.Mutator(
		t,
		MutatorReturnNil,
		"testdata/testdata.go",
		1,
	)
}
