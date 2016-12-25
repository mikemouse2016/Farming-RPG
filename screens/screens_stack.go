package screens

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

func (stack *Stack) Push(screen Screener) {
	stack.top = &StackItem{screen: screen, next: stack.top}
	stack.size++
}

func (stack *Stack) Pop() (value Screener) {
	if stack.size > 0 {
		value, stack.top = stack.top.screen, stack.top.next
		stack.size--
	}
	return
}

func (stack *Stack) Peek() (screen Screener) {
	if stack.size > 0 {
		screen = stack.top.screen
	}
	return
}
