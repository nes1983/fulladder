package fulladder

import (
	"testing"
)

func TestOr(t *testing.T) {
	c1 := make(chan bool)
	c2 := make(chan bool)
	o := make(chan bool)
	RunOr(c1, c2, o)
	go func() { c1 <- true; c2 <- true }()
	if !<-o {
		t.Error("true and true should be true")
	}

	go func() { c1 <- true; c2 <- false }()
	if !<-o {
		t.Error("true and false should be true")
	}

	go func() { c1 <- false; c2 <- false }()
	if <-o {
		t.Error("false and false should be false")
	}
}

func TestHalfAdder(t *testing.T) {
	inA, inB, out, carry := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	RunHalfAdder(inA, inB, out, carry)
	go func() { inA <- false; inB <- false }()
	if <-out != false {
		t.Error("0 + 0 is 0, but was not.")
	}
	if <-carry != false {
		t.Error("0 + 0 should not make a carry over, but did.")
	}
}

func TestHalfAdderCarry(t *testing.T) {
	inA, inB, out, carry := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	RunHalfAdder(inA, inB, out, carry)
	go func() { inA <- true; inB <- true }()
	if <-out != false {
		t.Error("1 + 1 is 0, but was not.")
	}
	if <-carry != true {
		t.Error("1 + 1 should carry over, but did not.")
	}
}

func TestHalfAdderNoCarry(t *testing.T) {
	inA, inB, out, carry := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	RunHalfAdder(inA, inB, out, carry)
	go func() { inA <- true; inB <- false }()
	if <-out != true {
		t.Error("1 + 0 is 1, but was not.")
	}
	if <-carry != false {
		t.Error("1 + 0 should not carry over, but did.")
	}
}

func TestFullAdder(t *testing.T) {
	inA, inB, carryIn, out, carry := make(chan bool), make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	RunFullAdder(inA, inB, carryIn, out, carry)
	go func() { inA <- true; inB <- true; carryIn <- true }()
	if <-out != true {
		t.Error("1 + 1 + 1 should be 1, but was not.")
	}
	if <-carry != true {
		t.Error("1 + 1 + 1 should carry over, but did not")
	}
}
