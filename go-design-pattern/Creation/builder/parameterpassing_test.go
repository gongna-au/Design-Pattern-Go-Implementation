package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 其实可以看到，绝大多数情况下直接使用后面的这种方式就可以了，
// 并且在编写公共库的时候，强烈建议入口的参数都可以这么传递，
// 这样可以最大程度的保证我们公共库的兼容性，避免在后续的更新的时候出现破坏性的更新的情况。

func TestNewResourcePoolConfig(t *testing.T) {
	type args struct {
		name string
		opts []ResourcePoolConfigOptFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: "test",
				opts: []ResourcePoolConfigOptFunc{
					func(option *ResourcePoolConfigOption) {
						option.minIdle = 2
					},
				},
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: 10,
				maxIdle:  9,
				minIdle:  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(tt.args.name, tt.args.opts...)
			require.Equalf(t, tt.wantErr, err != nil, "error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
