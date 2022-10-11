package bridge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBridge(t *testing.T) {

	remote := NewRemoteControl(NewTv())
	err := remote.TogglePower()
	require.Nil(t, err)
}
