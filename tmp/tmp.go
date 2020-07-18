package tmp

import (
	complex2 "github.com/go-demo/complex"
	"sync"
)

var lock = &sync.Mutex{}

var zeroComplex complex2.Complex
var initComplex complex2.Complex

func init() {
	initComplex = complex2.NewComplex(0, 0)
}

func ZeroComplex() complex2.Complex {
	lock.Lock()
	defer lock.Unlock()
	if zeroComplex == initComplex {
		zeroComplex = complex2.NewComplex(1, 1)
	}
	return zeroComplex
}
