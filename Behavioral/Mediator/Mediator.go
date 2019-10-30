package Mediator

import (
	"fmt"
	"time"
)

type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) broadcastMessage(sender *User, msg string) {
	msg = formatMessage(msg, sender.getName())
	for _, receiver := range c.users {
		if receiver != sender {
			receiver.msgBox <- msg
		}
	}
}

type User struct {
	name     string
	msgBox   chan string
	chatRoom *ChatRoom
}

func (u *User) bindChatRoom(c *ChatRoom, messageBoxSize int) {
	u.msgBox = make(chan string, messageBoxSize)
	u.chatRoom = c
	c.users = append(c.users, u)
}

func (u User) getName() string {
	return u.name
}

func (u *User) sendMessage(msg string) {
	u.msgBox <- formatMessage(msg, "Me")
	u.chatRoom.broadcastMessage(u, msg)
}

func (u User) displayMessage() {
	for {
		select {
		case msg := <-u.msgBox:
			fmt.Println(msg)
		}
	}
}

func formatMessage(msg, userName string) string {
	return fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05") + " " + userName + ": " + msg)
}

func Test() {
	var chatRoom ChatRoom
	var John = User{name: "John"}
	var Mike = User{name: "Mike"}
	John.bindChatRoom(&chatRoom, 10)
	Mike.bindChatRoom(&chatRoom, 10)
	John.sendMessage("hello, I'm John.")
	Mike.sendMessage("hi, I'm Mike.")

	go John.displayMessage()

	time.Sleep(time.Second * 2)
}

/*
通过
*/
