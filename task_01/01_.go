package main

import (
	"fmt"
	"math"
)

type Point struct{
	x1, y1, x2, y2 int
}

func newOffice(new Point, old Point) int {
	var result int
	var res Point

	if new.x2 < old.x2 { //rigth x point
		res.x2 = old.x2	
	} else {
		res.x2 = new.x2
	}

	if new.x1 < old.x1 { //left x point
		res.x1 = new.x1
	} else {
		res.x1 = old.x1
	}

	if new.y2 < old.y2 { //rigth y point
		res.y2 = old.y2
	} else {
		res.y2 = new.y2
	}

	if new.y1 < old.y1 { //left y point
		res.y1 = new.y1
	} else {
		res.y1 = old.y1
	}

	x_res := math.Abs(float64(res.x2 - res.x1))
	y_res := math.Abs(float64(res.y2 - res.y1))

	if x_res > y_res {
		result = int(x_res * x_res)
	} else {
		result =int(y_res * y_res)
	}
	return result
}

func main(){
	var new, old Point

	fmt.Scan(&old.x1, &old.y1, &old.x2, &old.y2)
	fmt.Scan(&new.x1, &new.y1, &new.x2, &new.y2)
	fmt.Println(newOffice(new, old))
}