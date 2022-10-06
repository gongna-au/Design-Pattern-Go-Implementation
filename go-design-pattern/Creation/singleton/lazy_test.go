package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
