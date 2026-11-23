package main

import (
	"fmt"
	"math"
)

func main() {
	// http.NewServeMux()
	// http.HandleFunc("GET /", HomeHandler)
	// http.HandleFunc("POST /", ReportSheetHandler)

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	os.Exit(1)
	// }

	// ReportSheet()

	fmt.Println(Quadratic(4, -5, 25))

}

func Quadratic(a, b, c float64) (float64, float64) {
	// discriminant
	d := math.Sqrt(float64((b*b - 4*a*c)))

	// check
	if d > 0 {
		root1, root2 := (-b+d)/(2.0*a), (-b-d)/(2.0*a)
		return root1, root2
	} else if d < 0 {
		return 0.0, 0.0
	} else {
		root1, root2 := (-b+d)/(2.0*a), (-b-d)/(2.0*a)
		return root1, root2
	}
}

// func Quadratic(a, b, c float64) (float64, float64) {
// 	discriminant := b*b - 4*a*c

// 	if discriminant < 0 {
// 		return 0.0, 0.0 // no real roots
// 	}

// 	d := math.Sqrt(discriminant)

// 	root1 := (-b + d) / (2 * a)
// 	root2 := (-b - d) / (2 * a)

// return root1, root2

//}
