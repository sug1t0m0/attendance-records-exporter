package sample

type Num int

func (n Num) Add(number Num) Num {
	return n + number
}
