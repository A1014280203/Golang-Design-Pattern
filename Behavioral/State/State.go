package State

import (
	"fmt"
)

type State interface {
	name() string
	update(s *State, delta int)
}

type Role struct {
	roleName  string
	moodState State
	moodValue int
}

func (r Role) printMoodState() {
	fmt.Printf("mood state of %s is %s now\n", r.roleName, r.moodState.name())
}
func (r *Role) updateMood(delta int) {
	r.moodValue += delta
	r.checkMoodValue()
	r.moodState.update(&r.moodState, r.moodValue)
}

func (r *Role) checkMoodValue() {
	if r.moodValue > 100 {
		r.moodValue = 100
	} else if r.moodValue < 0 {
		r.moodValue = 0
	}
}

func NewState(moodName string) (state State) {
	moodNames := []string{"sad", "normal", "happy"}
	mc := MoodChanger{79, 39, moodNames}
	switch moodName {
	case "happy":
		state = &Happy{"happy", mc}
	case "normal":
		state = &Normal{"normal", mc}
	case "sad":
		state = &Sad{"sad", mc}
	default:
		panic("unknown mood name")
	}
	return state
}

type MoodChanger struct {
	upperBound, lowerBound int
	moodNames              []string
}

func (m *MoodChanger) update(state *State, moodValue int) {
	if moodValue < m.lowerBound {
		*state = NewState(m.moodNames[0])
	} else if moodValue > m.upperBound {
		*state = NewState(m.moodNames[2])
	} else {
		*state = NewState(m.moodNames[1])
	}
}

type Happy struct {
	curName string
	MoodChanger
}

func (h *Happy) name() string {
	return h.curName
}

type Normal struct {
	curName string
	MoodChanger
}

func (n *Normal) name() string {
	return n.curName
}

type Sad struct {
	curName string
	MoodChanger
}

func (s *Sad) name() string {
	return s.curName
}

func Test() {
	r := Role{"soldier", NewState("normal"), 60}
	r.printMoodState()
	r.updateMood(20)
	r.printMoodState()
	r.updateMood(-100)
	r.printMoodState()
	r.updateMood(60)
	r.printMoodState()
}

/*
状态模式特点
1. 条件语句放在内部，即结构体的方法里面。自行控制状态，外部调用者无需操心。
2. 状态和行为对应。通过改变状态，改变行为。外部调用者无需在调用方法时进行条件判断。
将角色和心情作为整体看，提高了内聚。分开看，则是增加了耦合。
*/
