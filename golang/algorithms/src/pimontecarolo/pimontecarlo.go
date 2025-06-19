package pimontecarolo

import (
	"fmt"
	"math"
	"math/rand"
)

const diamterOfCircle = 100
const radiusOfCircle = diamterOfCircle / 2

func estimatePI(numberOfThrowsInCircle int, numberOfThrows int) string {
	// To estimate PI we use a simple calculation
	// (areaOfCircle / areaOfSquare) = (PIr^2 / 4r^2)
	// PI = (4 * areaOfCircle) / areaOfSquare
	// areaOfCircle = numberOfThrowsInCircle, areaOfSquare = numberOfThrows
	PI := float64((4 * numberOfThrowsInCircle) / numberOfThrows)

	return fmt.Sprintf("%.2f", PI)
}

func isThrowInTheCircle(x int, y int) bool {
	// For an easy calculation, we try to re-plot the x and y axis on the positive side of the graph
	distanceX := math.Abs(float64(x - radiusOfCircle))
	distanceY := math.Abs(float64(y - radiusOfCircle))

	// Using pythagoras theorem, we have to evaluate the distance of the throw from the center of the circle
	// i.e d^2 = x^2 + y^2, where d(r) is the hypothenus and thus the distance of the point from the radius of the circle
	distanceOfThrowFromCenter := math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))

	// The throw is only in the circle if the distance from the center is less than the radius of the circle
	return distanceOfThrowFromCenter < radiusOfCircle
}

// We want to use monte carlo simulation to estimate the value of PI
// @see https://dlibin.net/posts/monte-carlo-simulation-javascript
// @see https://www.youtube.com/watch?v=BfS2H1y6tzQ
// Monte carlo simulation is an estimation that tries to derive an output using randomly selected input
// E.g to get the average height of all the people in the world, we could get the height of everyone and divide by the total number of people
// however this is gonna take time, hence we can use monte carlo to randonly select some people, calculate their average
// and use that to deduce the average height for everyone else.
// Monte carlo also plays with the law of large numbers i.e the larger your sample size, the likely you are to getting an accurate estimation
// For our example, we will be throwing like darts
func generateThrows(numberOfThrows int) (estimatedPI string) {
	numberOfThrowsInCircle := 0

	for i := 0; i < numberOfThrows; i++ {
		x, y := rand.Intn(diamterOfCircle), rand.Intn(diamterOfCircle)

		if isThrowInTheCircle(x, y) {
			numberOfThrowsInCircle += 1
		}
	}

	return estimatePI(numberOfThrowsInCircle, numberOfThrows)
}

func Init() {
	var numberOfThrows int

	fmt.Println("Program Begin")
	fmt.Println("=========================")
	fmt.Println("Enter the number of throws")

	fmt.Scan(&numberOfThrows)

	result := generateThrows(numberOfThrows)

	fmt.Printf("The result of estimating PI after %v throws is %v\n", numberOfThrows, result)
}
