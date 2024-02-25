package main

import (
	"fmt"

	"github.com/andremelinski/pos-goexpert/6-packaging/modulo/math"
)

func main() {
// result := math.MathStruct{} // inicia a struct sem nada
	result := math.MathStruct{A: 1,B: 2}

	fmt.Println(result.Add())

	result2 := math.NewMath(1,2)
	fmt.Println(result2.AddByStruct())
}