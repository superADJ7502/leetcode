package offer

type MinStack struct {
	Slice    []int
	MinSlice []int
}

func (m *MinStack) Push(x int) {
	m.Slice = append(m.Slice, x)
	if len(m.MinSlice) == 0 || x <= m.MinSlice[len(m.MinSlice)-1] {
		m.MinSlice = append(m.MinSlice, x)
		return
	}
}

func (m *MinStack) Pop() {
	popVal := m.Slice[len(m.Slice)-1]
	m.Slice = m.Slice[:len(m.Slice)-1]
	if popVal == m.MinSlice[len(m.MinSlice)-1] {
		m.MinSlice = m.MinSlice[:len(m.MinSlice)-1]
	}
}

func (m *MinStack) Top() int {
	return m.Slice[len(m.Slice)-1]
}

func (m *MinStack) Min() int {
	return m.MinSlice[len(m.MinSlice)-1]
}
