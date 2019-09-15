package main

import (
	"Golang-Design-Pattern/CreationalPatterns/SingletonPattern"
)

func main() {
	//FactoryPattern.Test()

	ins := SingletonPattern.GetInstance()
	SingletonPattern.Test(ins)
	ins2 := SingletonPattern.GetInstance()
	SingletonPattern.Test(ins2)
}
