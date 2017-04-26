package utils

import (
	"math"
	"math/rand"
)

type Key struct {
	value float64
	MinValue0 float64
	MinValue1 float64
	MaxValue0 float64
	MaxValue1 float64
	MaxValue2 float64
}
func NewKey() *Key {
	k := new(Key)
	k.value = 0
	k.MinValue0 = -math.MaxFloat64
	k.MinValue1 = -0.95*math.MaxFloat64
	k.MaxValue0 = math.MaxFloat64
	k.MaxValue1 = 0.95*math.MaxFloat64
	k.MaxValue2 = 0.9*math.MaxFloat64
	return k
}
func NewKeyValue(value float64) *Key {
	k := new(Key)
	k.value = value
	k.MinValue0 = -math.MaxFloat64
	k.MinValue1 = -0.95*math.MaxFloat64
	k.MaxValue0 = math.MaxFloat64
	k.MaxValue1 = 0.95*math.MaxFloat64
	k.MaxValue2 = 0.9*math.MaxFloat64
	return k
}
func (k *Key) CompareTo(o *Key) bool {
	return k.value < o.value
}
func (k *Key) Equals(o *Key) bool {
	return k.value == o.value
}
func (k *Key) getValue() float64 {
	return k.value
}
func (k *Key) simpleNew(b *Key) *Key {
	if k.value < b.value {
		return b
	} else {
		return k
	}
}
func (k *Key) randPoint(b *Key) *Key {
	nk := new (Key)
	if k.value < b.value {
		nk.value = (rand.Float64()*k.value) + b.value		
	} else {
		nk.value = (rand.Float64()*b.value) + k.value
	}
	return nk
}
func (k *Key) newPoint(b *Key) *Key {
	nk := new (Key)
	nk.value = k.value*0.5 + b.value*0.5
	return nk
}
