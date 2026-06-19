type MinStack struct {
	stack []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min: []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)

	if len(m.min) == 0 {
		m.min = append(m.min, val)
		return
	}
	curMin := m.min[len(m.min) -1]
	if curMin > val {
		m.min = append(m.min, val)
	} else {
		m.min = append(m.min, curMin)
	}
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack) - 1]
	m.min = m.min[:len(m.min) - 1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack) - 1]

}

func (m *MinStack) GetMin() int {
	return m.min[len(m.min) - 1]
}
