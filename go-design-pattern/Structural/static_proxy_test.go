package structural

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStaticProxy(t *testing.T) {
	p := NewUserProxy()
	p.SetUser(NewUser())
	err := p.Login("test", "12345")
	require.Nil(t, err)
}
