package main

// constraint
type SumResponse int

type Number interface {
	~int | float64
}



// func SomaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// T pode utilizar int | float64 ou a constraint Number
func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable -> comparada se a e b sao do msm tipo, se nao for, nao pode comparar devido a dif de tipagem 
func Compara[T comparable](a, b T) bool{
	return a==b
}

func main() {
	intObj := map[string]int{"andre": 100, "luiz": 200, "alice": 50}
	floatObj := map[string]float64{"andre": 100.64, "luiz": 200.11, "alice": 50.8}

	sum := SomaGenerics(intObj)
	sum2 := SomaGenerics(floatObj)
	
	
	// como o type SumResponse nao eh type int, mesmo implementando ele por tras, precisa user "~" dentro da interface Number
	//  pra dizer que eles tem a "msm" tipagem
	intObj2 := map[string]SumResponse{"andre": 100, "luiz": 200, "alice": 50}
	sum3 := SomaGenerics(intObj2)
	println(sum)
	println(sum2)
	println(sum3)
	println(Compara(10,10))
	println(Compara("10","lala"))
}