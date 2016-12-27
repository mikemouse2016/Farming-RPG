package screens

import "github.com/phestek/farming-rpg/eng"

type StackItem struct {
	screen Screener
	next   *StackItem
}

type Stack struct {
	top  *StackItem
	size int
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Push(window *eng.Window, screen Screener) {
	stack.top = &StackItem{screen: screen, next: stack.top}
	stack.size++
	stack.top.screen.Init(window)
}

func (stack *Stack) Pop() {
	if stack.size > 0 {
		stack.top.screen.Dispose()
		stack.top = stack.top.next
		stack.size--
	}
}

func (stack *Stack) Peek() (screen Screener) {
	if stack.size > 0 {
		screen = stack.top.screen
	}
	return
}
