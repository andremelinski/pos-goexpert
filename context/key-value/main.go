package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// pega o valor desas chaves
	ctx = context.WithValue(ctx, "token", "senha")
}

func bookHotel(ctx context.Context, name string){
	token := ctx.Value("token")
	fmt.Println(token)
}
