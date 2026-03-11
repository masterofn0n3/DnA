package main

import "sort"

/*
target = 12, position = [10,8,5,3,0], speed = [2,4,1,3,1]
time = 1, 1, 7, 3, 12
*/
type Car struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := []*Car{}

	for i := range position {
		cars = append(cars, &Car{pos: position[i], speed: speed[i]})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	var count int
	var maxTime float64
	for _, v := range cars {
		time := float64(target-v.pos) / float64(v.speed)
		if time > maxTime {
			maxTime = time
			count++
		}
	}
	return count
}
