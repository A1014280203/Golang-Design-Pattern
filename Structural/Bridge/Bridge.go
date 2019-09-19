package Bridge

import "fmt"

// Button to order some
type Button interface {
	DingLing()
}

// for vegetable
type ForVegetable struct {
}

func (f ForVegetable) DingLing() {
	fmt.Println("Order some vegetables")
}

// for fruit
type ForFruit struct {
}

func (f ForFruit) DingLing() {
	fmt.Println("Order some fruits")
}

type ButtonBox struct {
	button Button
}

func (box *ButtonBox) SetButton(b Button) {
	box.button = b
}

// table which button is on
type Table struct {
	box ButtonBox
}

func (t *Table) SetButton(b Button) {
	t.box.SetButton(b)
}

func (t *Table) OrderSome() {
	t.box.button.DingLing()
}

func Test() {
	table := Table{}
	vBtn := ForVegetable{}
	fBtn := ForFruit{}
	table.SetButton(vBtn)
	table.OrderSome()
	table.SetButton(fBtn)
	table.OrderSome()
}

/*
refer: https://github.com/sevenelevenlee/go-patterns/blob/master/21-bridge-pattern/bridge.go
桌子上有一个可以修改其中按钮类型的按钮箱
*/
