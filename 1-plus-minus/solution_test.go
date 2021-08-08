package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	in  int
	out string
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			in:  199,
			out: "not possible",
		},
		{
			in:  26712,
			out: "-+--",
		},

		{
			in:  0,
			out: "not possible",
		},
		{
			in:  9,
			out: "not possible",
		},

		{
			in:  211,
			out: "--",
		},
		{
			in:  3111,
			out: "---",
		},

		{
			in:  112,
			out: "+-",
		},
		{
			in:  1113,
			out: "++-",
		},
	}

	for _, q := range questions {
		out, in := q.out, q.in
		ast.Equal(out, PlusMinus(in), "Input:%v", in)
	}
}
