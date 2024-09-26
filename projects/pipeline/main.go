package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)
	var prev string
	for str := range inputStream {
		if str != prev {
			outputStream <- str
			prev = str
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		for _, r := range "112334456" {
			inputStream <- string(r)
		}
		close(inputStream)
	}()

	for result := range outputStream {
		fmt.Println(result)
	}
}
