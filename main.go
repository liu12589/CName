package main

import (
	"fmt"
	"github.com/liu12589/cname/client"
)

func main() {
	sv := client.NewClient()
	arr := sv.GetNameByProbability(10)
	fmt.Println(arr)
}
