package main

import (
	"cname/client"
	"fmt"
)

func main() {
	sv := client.NewClient()
	arr := sv.GetNameByProbability(10)
	fmt.Println(arr)
}
