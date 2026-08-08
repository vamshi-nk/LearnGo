package main

import (
	"fmt"
)

func main() {
	var student_name string
	var (
		math_marks     int
		physics_marks  int
		computer_marks int
		total_marks    int
		average_marks  float64
		grade          string
	)

	fmt.Println("Please enter you're name :")
	fmt.Scan(&student_name)

	fmt.Println("Enter your Maths marks :")
	fmt.Scan(&math_marks)

	fmt.Println("Enter your Physics marks")
	fmt.Scan(&physics_marks)

	fmt.Println("Enter your Computer marks ")
	fmt.Scan(&computer_marks)

	total_marks = math_marks + physics_marks + computer_marks

	average_marks = float64(total_marks) / 3

	if average_marks >= 90 {
		grade = "A"
	} else if average_marks >= 80 {
		grade = "B"

	} else if average_marks >= 70 {
		grade = "C"

	} else if average_marks >= 60 {
		grade = "D"

	} else {
		grade = "Fail"
	}

	fmt.Println(
		"Student Repoprt \n",
		"Student Name		: ", student_name,
		"\nMathematics		: ", math_marks,
		"\nPhysics			: ", physics_marks,
		"\nComputers		: ", computer_marks,

		"\nTotale marks		:", total_marks,

		"\nAverage marks	:", average_marks,

		"\nGrade			:", grade,
	)
}
