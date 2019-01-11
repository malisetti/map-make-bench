package mapb

func x(c <-chan int) {
	m := make(map[int]struct{})
	for i := range c {
		m[i] = struct{}{}
	}
}

func y(s int, c <-chan int) {
	m := make(map[int]struct{}, s)
	for i := range c {
		m[i] = struct{}{}
	}
}
