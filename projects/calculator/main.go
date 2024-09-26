package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		select {
		case x := <-firstChan:
			resultChan <- x * x
		case x := <-secondChan:
			resultChan <- x * 3
		case <-stopChan:
			return
		}
	}()

	return resultChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		firstChan <- 4
	}()
	fmt.Println("Результат из firstChan (4^2):", <-resultChan)

	go func() {
		secondChan <- 5
	}()
	fmt.Println("Результат из secondChan (5 * 3):", <-resultChan)

	go func() {
		stopChan <- struct{}{}
	}()

	// time.Sleep(1 * time.Second)
	fmt.Println("Тест завершён. Функция calculator остановлена.")
}
