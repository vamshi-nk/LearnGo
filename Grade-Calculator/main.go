package main

import (
	"fmt"
)

func validatemark(marks int ) bool {


	if marks < 0 {
		//fmt.Println("Invalid")
		return false
		
	}else if marks > 100{

		return false

	}else{

		return true
	}
}

func main() {
	var student_name string
	var (
		math_marks     int
		physics_marks  int
		computer_marks int
		total_marks    int
		average_marks  float64
		grade          string
		result			bool
	)

	fmt.Println("Please enter you're name :")
	fmt.Scan(&student_name)

	fmt.Println("Enter your Maths marks :")
	fmt.Scan(&math_marks)
	result = validatemark(math_marks)
	if result == false{
		fmt.Println("Invalid Maths marks. Enter a value between 0 and 100.")
		return	
	}

	fmt.Println("Enter your Physics marks")
	fmt.Scan(&physics_marks)
	result = validatemark(physics_marks)
	if result == false{
		fmt.Println("Invalid Physics marks. Enter a value between 0 and 100.")
		return	
	}


	fmt.Println("Enter your Computer marks ")
	fmt.Scan(&computer_marks)
	result = validatemark(computer_marks)
	if result == false{
		fmt.Println("Invalid Computer marks. Enter a value between 0 and 100.")
		return	
	}


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
