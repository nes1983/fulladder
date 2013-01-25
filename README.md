fulladder
=========

A parallel fulladder in go.

This is just a little finger exercise for myself, inspired by SICP, on how to implement a full adder in go. Written by an idiot, signifying nothing, since any go implementation probably will require a full adder somewhere in the machine already.

It's implemented entirely using channels. There's no other kind of object around. There's a RunProbe method that isn't being used, but that's useful to see what's in a channel.