package main

import (
	"fmt"

	"github.com/andremelinski/pos-goexpert/6-packaging/workspaces/math"
	"github.com/google/uuid"
)

func main() {
// result := math.MathStruct{} // inicia a struct sem nada
uuidv6, _ := uuid.NewV6()
	fmt.Println(uuidv6.String())
	result2 := math.NewMath(1,2)
	fmt.Println(result2.AddByStruct())
}