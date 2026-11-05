package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const radius = 20

var robots = [][2]float64{
	{37.30, 6.60}, {60.41, 20.88}, {4.38, 38.32}, {81.29, 38.44}, {38.30, 64.64},
	{73.56, 21.79}, {36.17, 12.12}, {66.40, 46.62}, {71.24, 71.59}, {94.17, 97.08},
	{11.93, 34.02}, {28.66, 22.71}, {65.26, 4.20}, {92.68, 13.00}, {45.08, 21.11},
	{36.73, 72.17}, {89.26, 99.73}, {7.03, 97.35}, {93.67, 85.59}, {15.16, 62.72},
	{74.32, 75.60}, {25.52, 5.68}, {70.45, 57.18}, {83.46, 54.36}, {55.68, 78.90},
	{18.69, 17.66}, {93.83, 79.34}, {53.28, 23.39}, {45.32, 28.50}, {3.91, 25.89},
	{91.37, 68.49}, {16.76, 72.46}, {6.06, 29.82}, {30.54, 57.78}, {89.79, 86.68},
	{25.28, 69.86}, {45.37, 89.03}, {63.78, 54.09}, {44.84, 9.85}, {48.06, 87.43},
	{91.47, 96.40}, {27.30, 73.46}, {44.33, 70.69}, {81.88, 15.63}, {21.41, 79.40},
	{27.34, 99.36}, {85.26, 19.60}, {67.26, 52.91}, {33.04, 95.21}, {41.35, 57.92},
}

var step = 1.0

func moveStep(x float64, y float64) (float64, float64) {
	var medx = 0.0
	var medy = 0.0
	var cnt = 0.0
	for ind := range 50 {
		var dx = robots[ind][0] - x
		var dy = robots[ind][1] - y
		if math.Sqrt(dx*dx+dy*dy) <= radius {
			medx = medx + robots[ind][0]
			medy = medy + robots[ind][1]
			cnt += 1
		}
	}

	medx /= cnt
	medy /= cnt

	dx := medx - x
	dy := medy - y
	dist := math.Sqrt(dx*dx + dy*dy)

	if dist <= step {
		return medx, medy
	}

	return x + dx/dist*step, y + dy/dist*step
}

func robot(i int, wg *sync.WaitGroup) {
	var x = robots[i][0]
	var y = robots[i][1]

	for {
		time.Sleep(100 * time.Millisecond)

		// var oldx = x
		// var oldy = y
		x, y = moveStep(x, y)
		robots[i][0] = x
		robots[i][1] = y
		fmt.Println(x, y, i)
		// if oldx == x && oldy == y {
		// 	break
		// }
	}

	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go robot(i, &wg)
	}

	// for i := range robots {
	// 	fmt.Println(robots[i][0], robots[i][1], i)
	// }
	wg.Wait()
}
