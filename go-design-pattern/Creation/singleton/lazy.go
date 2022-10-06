package singleton

import "sync"

var once = &sync.Once{}
var lazySingleton *Singleton

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
