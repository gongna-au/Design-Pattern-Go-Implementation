package flyweight

import (
	"testing"
)

type Data struct {
	want []string
	get  []string
}

func TestForest(t *testing.T) {

	f := NewForest()
	f.AddTree(1, 2, "柳树", "绿色", "高4米")
	f.AddTree(3, 4, "松树", "绿色", "高6米")
	f.AddTree(5, 6, "柏树", "绿色", "高8米")
	treeTypeFactory.OutputAllTrees()
}
