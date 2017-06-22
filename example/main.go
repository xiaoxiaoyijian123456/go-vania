package main

import (
	"fmt"
	vania "github.com/xiaoxiaoyijian123456/go-vania"
)

func main() {
	targets := []interface{}{
		"Team A", "Team B", "Team C",
	}
	objects := []interface{}{
		"Front-end Development",
		"Back-end Development",
		"Testing",
		"Documentation",
	}
	weights := [][]float64{
		{1, 2, 3, 2},
		{3, 1, 4, 2},
		{3, 4, 1, 1},
	}

	ret, v, err := vania.FairDistributor(targets, objects, weights, false)
	if err != nil {
		fmt.Println("Err: ", err.Error())
		return
	}
	fmt.Println(ret)
	fmt.Println(v)
}
