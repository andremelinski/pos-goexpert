package main

import (
	"fmt"
	"os"
)

func main(){
	for i := 0; i < 501; i++ {
		fileName := fmt.Sprintf("../tmp/file_%d.txt",i)
		f, err := os.Create( fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString("Hello World")
		
	}
}