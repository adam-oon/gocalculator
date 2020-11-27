package calc

import (
	"math"
	"testing"

	. "github.com/franela/goblin"
)

func TestAdd(t *testing.T) {
	gob := Goblin(t)
	gob.Describe("Calculator Add Operations", func() {
		gob.It("it should sum up two integers", func() {
			gob.Assert(Add(1, 2)).Equal(3)
			gob.Assert(Add(4, 150)).Equal(154)
			gob.Assert(Add(-4, 2)).Equal(-2)
			gob.Assert(Add(-5, 2)).Equal(-3)
		})
	})

	gob.Describe("Calculator Subtract Operations", func() {
		gob.It("it should subract second integer from first integer", func() {
			gob.Assert(Subtract(20, 1)).Equal(19)
			gob.Assert(Subtract(49, 150)).Equal(-101)
			gob.Assert(Subtract(-4, 2)).Equal(-6)
		})
	})

	gob.Describe("Calculator Multiply Operations", func() {
		gob.It("it should multiply two integers", func() {
			gob.Assert(Multiply(1, 2)).Equal(2)
			gob.Assert(Multiply(0, 150)).Equal(0)
			gob.Assert(Multiply(-4, 2)).Equal(-8)
			gob.Assert(Multiply(-5, -2)).Equal(10)
		})
	})

	gob.Describe("Calculator Division Operations", func() {
		gob.It("it should divide two integers", func() {
			gob.Assert(Divide(1, 0)).Equal(math.Inf(1))
			gob.Assert(Divide(6, 2)).Equal(3.0)
			gob.Assert(Divide(12, 5)).Equal(2.4)
		})
	})
}
