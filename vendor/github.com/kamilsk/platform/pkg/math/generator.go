package math

import "sync/atomic"

// Generator provides functionality to produce increasing sequence of numbers.
//
//  generator, sequence := new(Generator).At(10), make([]ID, 0, 10)
//
//  for range Sequence(10) {
//  	sequence = append(sequence, ID(generator.Next()))
//  }
//
type Generator uint64

// At sets the Generator to the new position.
func (generator *Generator) At(position uint64) *Generator {
	atomic.StoreUint64((*uint64)(generator), position)
	return generator
}

// Current returns a current value of the Generator.
func (generator *Generator) Current() uint64 {
	return atomic.LoadUint64((*uint64)(generator))
}

// Jump moves the Generator forward at the specified distance.
func (generator *Generator) Jump(distance uint64) uint64 {
	return atomic.AddUint64((*uint64)(generator), distance)
}

// Next moves the Generator one step forward.
func (generator *Generator) Next() uint64 {
	return atomic.AddUint64((*uint64)(generator), 1)
}
