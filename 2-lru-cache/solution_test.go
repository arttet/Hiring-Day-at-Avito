package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	in  []string
	out string
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			in:  []string{"A", "B", "C", "D", "A", "E", "D", "Z"},
			out: "C-A-E-D-Z",
		},
		{
			in:  []string{"A", "B", "A", "C", "A", "B"},
			out: "C-A-B",
		},
		{
			in:  []string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"},
			out: "E-D-Q-Z-C",
		},
	}

	for _, q := range questions {
		out, in := q.out, q.in
		ast.Equal(out, LRUCache(in), "Input:%v", in)
	}
}
