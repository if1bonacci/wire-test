package main

import "fmt"

func main() {
	ser := InitializeService()

	fmt.Println(ser.GetData())
}
