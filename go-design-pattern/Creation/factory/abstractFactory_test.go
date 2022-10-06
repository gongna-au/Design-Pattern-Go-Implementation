package factory

import (
	"testing"
)

type AbstractFactoryData struct {
	name string
	args string
	want jsonConfigParserFactory
}

func TestNewJsonConfigParserFactory(t *testing.T) {
	f := NewJsonConfigParserFactory()
	t.Run("RulesConfigParser", func(t *testing.T) {
		want := jsonRulesConfigParser{}
		get := f.CreateRuleParser()
		_, ok := get.(jsonRulesConfigParser)
		if !ok {
			t.Errorf("CreateRuleParser()= %v want = %v", get, want)
		}

	})
	t.Run("SystemConfigParser", func(t *testing.T) {
		want := jsonSystemConfigParser{}
		get := f.CreateSystemParser()
		_, ok := get.(jsonSystemConfigParser)
		if !ok {
			t.Errorf("CreateSystemParser()= %v want = %v", get, want)
		}
	})
}
