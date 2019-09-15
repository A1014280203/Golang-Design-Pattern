package SingletonPattern

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("run")
		instance = &singleton{}
	})
	return instance
}

func Test(ins *singleton) {
	fmt.Printf("%p\n", ins)
}

/*
refer: http://marcio.io/2015/07/singleton-pattern-in-go/

与其它可操作类静态变量或者类属性的语言不同，
Go通过命名空间（package）维护单例模式
*/
