package main

import "fmt"

func main() {
	var nama = []string{"andi", "Budi", "cacing"}
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}

}
