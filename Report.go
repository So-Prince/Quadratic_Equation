package main

import "fmt"

type StudentBio struct {
	name       string
	subjects   []string
	TestScore  float32
	ExamScore  float32
	TotalScore float32
	Grade      float32
}

func ReportSheet() {

	var Mary StudentBio
	{
		Mary.name = "Mary Abiodun"
		Mary.subjects = []string{
			"Math",
			"English",
			"Science",
			"Agric",
			"Literature",
			"Civic"}
		Mary.TestScore = 38.8
		Mary.ExamScore = 42.7
		Mary.TotalScore = Mary.TestScore + Mary.ExamScore
		Mary.Grade = Mary.TotalScore

		Peter := StudentBio{
			name:      "Peter Mark",
			subjects:  []string{"Economics", "Geography", "Government", "Arts"},
			TestScore: 37.5,
			ExamScore: 28.0,
		}
		Peter.TotalScore = Peter.TestScore + Peter.ExamScore
		Peter.Grade = Peter.TotalScore

		fmt.Println(Mary.name)
		fmt.Println("Subjects offered: ", Mary.subjects)
		fmt.Println("Your Test Score: ", Mary.TestScore)
		fmt.Println("Your Exam Score: ", Mary.ExamScore)
		fmt.Println("Your Total Score is: ", Mary.Grade)
		if Mary.Grade >= 80 {
			fmt.Println("Your grade is: Excellent - A")
		} else if Mary.Grade >= 60 {
			fmt.Println("Your grade is: Very Good - B")
		} else if Mary.Grade >= 40 {
			fmt.Println("Your grade is: Good - C")
		} else if Mary.Grade >= 20 {
			fmt.Println("Your grade is: Pass - D")
		} else {
			fmt.Println("Your grade is: Failed")
		}

		fmt.Println("\n")

		fmt.Println(Peter.name)
		fmt.Println("Subjects offered: ", Peter.subjects)
		fmt.Println("Your Test Score: ", Peter.TestScore)
		fmt.Println("Your Exam Score: ", Peter.ExamScore)
		fmt.Println("Your Total Score is: ", Peter.Grade)
		if Peter.Grade >= 80 {
			fmt.Println("Your grade is: Excellent - A")
		} else if Peter.Grade >= 60 {
			fmt.Println("Your grade is: Very Good - B")
		} else if Peter.Grade >= 40 {
			fmt.Println("Your grade is: Good - C")
		} else if Peter.Grade >= 20 {
			fmt.Println("Your grade is: Pass - D")
		} else {
			fmt.Println("Your grade is: Failed")
		}
	}

}
