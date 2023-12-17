package service

import (
	"fmt"
	"github.com/liu12589/cname/contants"
	"github.com/liu12589/cname/util"
	"math/rand"
	"time"
)

type Client struct {
	Ratio         int
	Range         int
	LastNameRange int
	SurnameRange  []int32
	Surname       map[int32]string
	LastName      map[int32]string
}

type GetName interface {
	GetNameByProbability(sum int) []string
}

func (c Client) GetNameByProbability(sum int) (result []string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sum; i++ {
		var name string
		num1 := rand.Intn(c.Range)
		num2 := rand.Intn(contants.MaxRatio)
		num3 := rand.Intn(c.LastNameRange)
		num4 := rand.Intn(c.LastNameRange)
		surnameSite := util.GetCloseSet(int32(num1), c.SurnameRange)
		surname := c.Surname[surnameSite]
		if num2 <= c.Ratio {
			name = fmt.Sprintf("%s%s%s", surname, c.LastName[int32(num3)], c.LastName[int32(num4)])
		} else {
			name = fmt.Sprintf("%s%s", surname, c.LastName[int32(num3)])
		}
		result = append(result, name)
	}
	return result
}
