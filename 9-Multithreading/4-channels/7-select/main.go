package main

import "time"

func main(){
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		time.Sleep(time.Second*2)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second*3)
		c2 <- 2
	}()
	// select -> utilizado como swith case para channels. 
	// o valor que chegar primeiro no canal sera utilizado. Caso tenha mais de um, o 1 canal populado aparecera no print
	// utilizado quando vc pode pegar a msm info de 2 maneiras diferentes ou para dar timeout se o canal demorar mt para popular
	select{
	case msg1 := <- c1:
		println("msg 1 came first ", msg1)
	case msg2 := <- c2:
		println("msg 2 came first ",msg2)
	// regra para timeout se os canais demorarem para ser populados
	case <- time.After(time.Second*1):
		println("timeout")
	// default aparecera a nao ser que algum canal receba o valor antes de cair nessa condicao
	default:
		println("default")
	}

}

