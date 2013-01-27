package fulladder

import (
	"testing"
)

func TestOr(t *testing.T) {
	c1, c2, o := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)
	RunOr(c1, c2, o)

	c1 <- true
	c2 <- true
	if <-o != true {
		t.Error("true and true should be true")
	}

	c1 <- true
	c2 <- false
	if !<-o {
		t.Error("true and false should be true")
	}

	c1 <- false
	c2 <- false
	if <-o {
		t.Error("false and false should be false")
	}
}

func TestHalfAdder(t *testing.T) {
	inA, inB, out, carry := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1),
		make(chan bool, 1)
	RunHalfAdder(inA, inB, out, carry)

	inA <- false
	inB <- false
	if <-out != false {
		t.Error("0 + 0 is 0, but was not.")
	}
	if <-carry != false {
		t.Error("0 + 0 should not make a carry over, but did.")
	}
}

func TestHalfAdderCarry(t *testing.T) {
	inA, inB, out, carry := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1),
		make(chan bool, 1)
	RunHalfAdder(inA, inB, out, carry)
	inA <- true
	inB <- true
	if <-out != false {
		t.Error("1 + 1 is 0, but was not.")
	}
	if <-carry != true {
		t.Error("1 + 1 should carry over, but did not.")
	}
}

func TestHalfAdderNoCarry(t *testing.T) {
	inA, inB, out, carry := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1),
		make(chan bool, 1)
	RunHalfAdder(inA, inB, out, carry)
	inA <- true
	inB <- false
	if <-out != true {
		t.Error("1 + 0 is 1, but was not.")
	}
	if <-carry != false {
		t.Error("1 + 0 should not carry over, but did.")
	}
}

func TestFullAdder(t *testing.T) {
	inA, inB, carryIn, out, carry := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1),
		make(chan bool, 1), make(chan bool, 1)
	RunFullAdder(inA, inB, carryIn, out, carry)
	inA <- true
	inB <- true
	carryIn <- true
	if <-out != true {
		t.Error("1 + 1 + 1 should be 1, but was not.")
	}
	if <-carry != true {
		t.Error("1 + 1 + 1 should carry over, but did not")
	}
}
