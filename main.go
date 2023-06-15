package main

import "fmt"

func main() {
	ser, err := InitializeService()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ser.GetData())
}
