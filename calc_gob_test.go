package calc

import (
	"math"
	"testing"

	. "github.com/franela/goblin"
)

func TestAdd(t *testing.T) {
	gob := Goblin(t)
	gob.Describe("Calculator Operations", func() {
		gob.It("it should sum up two integers", func() {
			gob.Assert(Add(1, 2)).Equal(3)
			gob.Assert(Add(4, 150)).Equal(154)
			gob.Assert(Add(-4, 2)).Equal(-2)
			gob.Assert(Add(-5, 2)).Equal(-3)
		})

		gob.It("it should divide two integers", func() {
			gob.Assert(Divide(1, 0)).Equal(math.Inf(1))
			gob.Assert(Divide(6, 2)).Equal(3.0)
			gob.Assert(Divide(12, 5)).Equal(2.4)

		})
	})
}
