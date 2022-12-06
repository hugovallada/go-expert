package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	i := 0
	for i < 10 {
		fmt.Printf("%d: Task: %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		i++
		wg.Done()
	}
}

/**
*
* WaitGroups:
* 	Adicionar qtd de tarefas
*   Terminar a operação
*   Esperar pelas operações
 */
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)
	go task("A", &waitGroup)
	go func() {
		i := 0
		for i < 10 {
			fmt.Printf("%d: Task: %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			i++
			waitGroup.Done()
		}
	}()
	go task("B", &waitGroup)
	waitGroup.Wait()
}
