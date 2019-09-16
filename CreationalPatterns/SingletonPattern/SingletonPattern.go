package SingletonPattern

import (
	"fmt"
	"sync"
)

// 创建结构体
type singleton struct {
}

// 创建需要维护的package内部的全局变量
var instance *singleton
var once sync.Once

// 外部调用的实例化接口
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func Test() {
	ins1 := GetInstance()
	fmt.Printf("%p\n", ins1)
	ins2 := GetInstance()
	fmt.Printf("%p\n", ins2)

}

/*
refer: http://marcio.io/2015/07/singleton-pattern-in-go/

1. 与其它可操作类静态变量或者类属性的语言不同，Go通过命名空间（package）维护单例模式
2. 设置singleton结构体私有是为了控制实例化行为，否则单例模式就没有意义
3. 在其他语言中，例如Python，可以通过类属性配合__new__或者__init__方法控制单例，也就不需要私有化类
4. 大写字母开头表示能被其它包（命名空间）访问或调用，所以
*/
