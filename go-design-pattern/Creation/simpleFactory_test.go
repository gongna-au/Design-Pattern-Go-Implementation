package creation

import (
	"reflect"
	"testing"
)

type SimpleFactoryDatas struct {
	name string
	args string
	want IRuleConfigParser
}

func TestNewIRuleConfigParser(t *testing.T) {
	tests := []SimpleFactoryDatas{
		{
			name: "test Json",
			args: "json",
			want: jsonRuleConfigParser{},
		},
		{
			name: "test Yaml",
			args: "yaml",
			want: yamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := NewIRuleConfigParser(tt.args)
			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("NewIRuleConfigParser()= %v want = %v", get, tt.want)
			}
		})
	}

}
