package main

func main() {
	forever := make(chan bool)
	
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	// Caso não existe uma chance de encher o canal, acontece um deadlock
	<-forever
}
