package main

func main(){
	forever := make(chan bool)

	go func(){
		for i := 0; i < 2; i++ {
			println(i)
		}
		// nao da dead lock se vc enche o canal com alguma coisa.
		// valor precisa estar dentro de outra go routine para nao dar erro de dead lock
		forever <- true
		}()
		
	<- forever
}