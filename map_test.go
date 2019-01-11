package mapb

import "testing"

func BenchmarkX(b *testing.B) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := 0; n < b.N; n++ {
			ch <- n
		}
	}()

	x(ch)
}

func BenchmarkX2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < n; i++ {
				ch <- i
			}
		}()

		x(ch)
	}
}

func BenchmarkY(b *testing.B) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := 0; n < b.N; n++ {
			ch <- n
		}
	}()

	y(b.N, ch)
}

func BenchmarkY2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < n; i++ {
				ch <- i
			}
		}()

		y(n, ch)
	}
}
