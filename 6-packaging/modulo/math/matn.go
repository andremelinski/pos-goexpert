package math

type MathStruct struct {
	A int
	B int
	c int
}

func (m MathStruct) Add() int {
	return m.A + m.B
}

type PrivateMathStruct struct {
	a int
	b int
}

// "injeta" os valores pelo NewMath na struct e nao diretamente pela struct.
//
//	Nao posso mais fucar mudando os valores de A e B diretamente, o q protege a struct
func NewMath(a, b int) PrivateMathStruct {
	return PrivateMathStruct{a, b}
}

func (m *PrivateMathStruct) AddByStruct() int {
	return m.a + m.b
}
