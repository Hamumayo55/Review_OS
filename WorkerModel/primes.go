package main

import "fmt"

func main() {

	prime_flag := true
	prime_list := []int{}

	for i := 2; i < 100; i++ {
		prime_flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime_flag = false
				break
			}
		}
		if prime_flag == true {
			prime_list = append(prime_list, i)
		}
	}
	fmt.Println(prime_list)
}
