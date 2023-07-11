package main

import "fmt"

// todo 未通过
func main() {
	e := Constructor(10)
	for i := 0; i < 4; i++ {
		fmt.Println(e.Seat())
	}
	e.Leave(4)
	fmt.Println(e.Seat())
}

type ExamRoom struct {
	seats []int
	n     int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		seats: make([]int, n),
		n:     n,
	}
}

func (this *ExamRoom) Seat() int {
	max, pos := 0, 0
	for i := 0; i < this.n; i++ {
		if i == 0 && this.seats[i] == 0 {
			this.seats[i] = 1
		} else if i == 0 {

		} else if this.seats[i] == 0 && this.seats[i-1] == -1 {
			this.seats[i] = 1
		} else if this.seats[i] == 0 {
			this.seats[i] = this.seats[i-1] + 1
		}
	}

	for i := 0; i < this.n; i++ {
		if max < this.seats[i] {
			max = this.seats[i]
			pos = i
		}
	}

	if max == this.n {
		pos = 0
	} else if pos+1 == this.n {

	} else {
		pos = pos - (max+1)/2
	}

	this.seats[pos] = -1
	fmt.Println(this.seats)

	return pos
}

func (this *ExamRoom) Leave(p int) {
	if p < this.n {
		this.seats[p] = 0
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
