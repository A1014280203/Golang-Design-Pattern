package Command

import "fmt"

type Command interface {
	Do()
}

// receiver
type Pen struct {
}

func (p Pen) Write(text string) {
	fmt.Println("Pen is writing", text)
}

// command
type PenWriteCommand struct {
	pen  *Pen
	text string
}

func (pwc PenWriteCommand) Do() {
	pwc.pen.Write(pwc.text)
}

// invoker
type Computer struct {
}

func (c *Computer) Execute(command Command) {
	command.Do()
}

func Test() {
	pen := &Pen{}
	pwc := PenWriteCommand{pen: pen, text: "message"}
	var computer Computer
	computer.Execute(pwc)
}

/*
refer: https://en.wikipedia.org/wiki/Command_pattern#Go
用结构体的方式包装需要执行的行为，以为其提供更多操作
*/
