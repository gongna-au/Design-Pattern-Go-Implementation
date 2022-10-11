package structural

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDynamicPorxy(t *testing.T) {
	_, err := MockMemoryUserStore().GetUser(context.Background(), 1)
	require.Nil(t, err)
}
