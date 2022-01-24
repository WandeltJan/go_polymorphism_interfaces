package main

import "fmt"

//interfaces
type error interface {
	Error() string
}

//structs
type networkProblem struct {
	message string
	code    int
}

type SyntaxError struct {
	Line int
	Col  int
}

//interface struct functions
func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message: %s, code: %v", np.message, np.code)
}

func (se SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", se.Line, se.Col)
}

//"normal" functions
func ErrorHandler(err error) {
	fmt.Println(err.Error())
}

// main function
func main() {

	er1 := networkProblem{
		message: "Unauthorized",
		code:    401,
	}

	er2 := networkProblem{
		message: "Resource not found",
		code:    404,
	}

	se1 := SyntaxError{
		Line: 5,
		Col: 4,
	}
	ErrorHandler(er1)
	ErrorHandler(er2)
	ErrorHandler(se1)
}