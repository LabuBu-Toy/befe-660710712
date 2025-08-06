package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("Name is required")
	}

	if s.Year < 1 || s.Year > 4 {
		return errors.New("Year must be between 1-4")
	}

	if s.GPA < 1 || s.GPA > 4 {
		return errors.New("GPA must be between 1-4")
	}
	return nil
}

func main() {
	var st Student = Student{
		ID:    "1",
		Name:  "Thanawit",
		Email: "udompol_t@silpakorn.edu",
		Year:  3,
		GPA:   2.99,
	}

	students := []Student{
		{
			ID:    "1",
			Name:  "Thanawit",
			Email: "udompol_t@silpakorn.edu",
			Year:  3,
			GPA:   2.99,
		},
		{
			ID:    "2",
			Name:  "alice",
			Email: "udompol_t@silpakorn.edu",
			Year:  3,
			GPA:   1.99,
		},
	}

	newStudent := Student{
		ID:    "3",
		Name:  "test",
		Email: "udompol_t@silpakorn.edu",
		Year:  3,
		GPA:   3.6,
	}

	students = append(students, newStudent)

	for i, student := range students {
		fmt.Printf("%d Honor %v\n", i, student.IsHonor())
		fmt.Printf("%d Validation %v\n", i, student.Validate())

	}

}
