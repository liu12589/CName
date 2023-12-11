package main

import (
	"CName/client"
	"fmt"
)

func main() {
	sv := client.NewClient()
	arr := sv.GetNameByProbability(10)
	fmt.Println(arr)
}
