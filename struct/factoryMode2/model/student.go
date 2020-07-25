package model

type student struct {
	Name  string
	score float64
}

// 因为 student 这个结构体的首字母是小写
// 相当于private,只能在自己所在的包内使用
// 因此可以通过工厂模式来解决这个问题

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

// 如果score首字母小写，则在其它包不可以直接访问
// 因此可以直接提供一个方法
func (student *student) GetScore() float64 {
	return (*student).score
}
