package fulladder

import (
	"fmt"
)

func RunXor(inA chan bool, inB chan bool, out chan bool) {
	RunGate(inA, inB, out, func(a bool, b bool) bool { return a != b })
}

func RunAnd(inA chan bool, inB chan bool, out chan bool) {
	RunGate(inA, inB, out, func(a bool, b bool) bool { return a && b })
}

func RunOr(inA chan bool, inB chan bool, out chan bool) {
	RunGate(inA, inB, out, func(a bool, b bool) bool { return a || b })
}

func RunGate(inA chan bool, inB chan bool, out chan bool, compare func(bool, bool) bool) {
	go func() {
		for {
			out <- compare(<-inA, <-inB)
		}
	}()
}

func RunSplitter(in chan bool, outA chan bool, outB chan bool) {
	go func() {
		for {
			post := <-in
			outA <- post
			outB <- post
		}
	}()
}

func RunHalfAdder(inA chan bool, inB chan bool, mainOut chan bool, carryOver chan bool) {
	a1, a2, b1, b2 := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	RunSplitter(inA, a1, a2)
	RunSplitter(inB, b1, b2)
	RunXor(a1, b1, mainOut)
	RunAnd(a2, b2, carryOver)
}

func RunFullAdder(inA chan bool, inB chan bool, carryIn chan bool, out chan bool, carryOut chan bool) {
	c1, o1, c2 := make(chan bool), make(chan bool), make(chan bool)
	RunHalfAdder(inA, inB, o1, c1)
	RunHalfAdder(carryIn, o1, out, c2)
	RunOr(c1, c2, carryOut)
}

func RunProbe(in chan bool, out chan bool) {
	go func() {
		for {
			post := <-in
			fmt.Println(post)
			out <- post
		}
	}()
}
