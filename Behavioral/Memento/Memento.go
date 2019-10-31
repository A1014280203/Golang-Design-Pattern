package Memento

import (
	"errors"
	"fmt"
)

type Number struct {
	value   int
	manager *Manager
}

func (number *Number) getValue() int {
	return number.value
}

func (number *Number) add(n int) {
	number.value += n
}

func (number *Number) bindManager(manager *Manager) {
	number.manager = manager
}

func (number *Number) recordNumber() {
	number.manager.set(number.value)
}

func (number *Number) recoverNumber() {
	numberStored, err := number.manager.get()
	if err == nil {
		number.value = numberStored
	}
}

type Manager struct {
	memento Memento
}

func (manager *Manager) get() (int, error) {
	return manager.memento.getState()
}

func (manager *Manager) set(data int) {
	manager.memento.setSate(data)
}

type Memento struct {
	states []int
}

func (memento *Memento) getState() (int, error) {
	index := len(memento.states) - 1
	if index < 0 {
		return 0, errors.New("no state kept now")
	}
	defer func() { memento.states = append(memento.states[:index], memento.states[index+1:]...) }()
	return memento.states[index], nil
}

func (memento *Memento) setSate(state int) {
	memento.states = append(memento.states, state)
}

func Test() {
	manager := Manager{}
	number := Number{value: 1}
	number.bindManager(&manager)
	number.add(2)
	fmt.Println(number.getValue())
	number.recordNumber()
	number.add(2)
	fmt.Println(number.getValue())
	number.recoverNumber()
	fmt.Println(number.getValue())
}
