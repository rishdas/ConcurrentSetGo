package harrisll

import (
	"math"
	"math/rand"
)

type key struct {
	value float64
	minValue0 float64
	minValue1 float64
	maxValue0 float64
	maxValue1 float64
	maxValue2 float64
}
func newKey() *key {
	k := new(key)
	k.value = 0
	k.minValue0 = -math.MaxFloat64
	k.minValue1 = -0.95*math.MaxFloat64
	k.maxValue0 = math.MaxFloat64
	k.maxValue1 = 0.95*math.MaxFloat64
	k.maxValue2 = 0.9*math.MaxFloat64
	return k
}
func NewKeyValue(value float64) *key {
	k := new(key)
	k.value = value
	k.minValue0 = -math.MaxFloat64
	k.minValue1 = -0.95*math.MaxFloat64
	k.maxValue0 = math.MaxFloat64
	k.maxValue1 = 0.95*math.MaxFloat64
	k.maxValue2 = 0.9*math.MaxFloat64
	return k
}
func (k *key) compareTo(o *key) bool {
	return k.value < o.value
}
func (k *key) equals(o *key) bool {
	return k.value == o.value
}
func (k *key) getValue() float64 {
	return k.value
}
func (k *key) simpleNew(b *key) *key {
	if k.value < b.value {
		return b
	} else {
		return k
	}
}
func (k *key) randPoint(b *key) *key {
	nk := new (key)
	if k.value < b.value {
		nk.value = (rand.Float64()*k.value) + b.value		
	} else {
		nk.value = (rand.Float64()*b.value) + k.value
	}
	return nk
}
func (k *key) newPoint(b *key) *key {
	nk := new (key)
	nk.value = k.value*0.5 + b.value*0.5
	return nk
}
