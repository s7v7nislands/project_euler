package main

import "fmt"

func fibonacci(end int) chan int {
	c := make(chan int)

	go func() {
		i, k := 1, 2
		for {
			if i > end {
				close(c)
				break
			}
			c <- i
			t := i + k
			i, k = k, t
		}

	}()

	return c
}

func main() {

    sum := 0
	for i := range fibonacci(4000000) {
        if i % 2 == 0 {
            sum += i
        }
	}

    fmt.Println(sum)

}
