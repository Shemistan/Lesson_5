package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("operation nit valid", func(t *testing.T) {
		//exp := 0
		_, err := Sum(1, 2, "test")

		assert.Equal(t, err.Error(), "unknown operation")

	})

	tests := []test{
		{
			name:      "test sum",
			a:         1,
			b:         2,
			operation: "+",
			exp:       3,
		},
		{
			name:      "test sub",
			a:         1,
			b:         2,
			operation: "-",
			exp:       -1,
		},
	}

	for _, test := range tests {
		res, err := Sum(test.a, test.b, test.operation)
		if err != nil {
			log.Println("expected error, got nil", test.name)
		}
		assert.Equal(t, res, test.exp)

	}

}

type test struct {
	name      string
	a         int
	b         int
	operation string
	exp       int
}
