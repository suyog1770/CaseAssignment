package main

import "fmt"

func main() {
	//hardcoded map created as per the data given in problem statement
	Score := map[string]int{"A":90,"B":48,"C":63,"D":75,"E":70}

	//result map to print first name of student and then grade
	Result := map[string]string{}

	//iteration on the score map for assigning the grade to students
	for name, marks := range Score {
		switch {
		case marks >= 0 && marks <= 49:
			Result[name] = "F"
		case marks >= 50 && marks <= 59:
			Result[name] = "E"
		case marks >= 60 && marks <= 69:
			Result[name] = "D"
		case marks >= 70 && marks <= 79:
			Result[name] = "C"
		case marks >= 80 && marks <= 89:
			Result[name] = "B"
		case marks >= 90 && marks <= 100:
			Result[name] = "A"
		}
	}

	fmt.Println(Result)

}

