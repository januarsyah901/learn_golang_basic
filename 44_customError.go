package main

import "fmt"

type ValitadionError struct {
	Message string
}

func (e *ValitadionError) Error() string {
	return e.Message
}

type UserNotFoundError struct {
	Message string
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}

func login(username string, data any) error {
	if username == "" {
		return &ValitadionError{Message: "Username is required"}
	}
	if username != "rodot" {
		return &UserNotFoundError{Message: "User not found"}
	}
	return nil
}
func main() {
	err := login("root",nil)
	// menggunakan if
	if err != nil {
	// 	if _, ok := err.(*ValitadionError); ok {
	// 		fmt.Println("Error:", err)
	// 	} else if _, ok := err.(*UserNotFoundError); ok {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		fmt.Println("Error:",err.Error())
	// 	}
	// } else {
	// 	fmt.Println("Login successful")
	// menggunakan swift
	switch finalError := err.(type){
	case *ValitadionError:
		fmt.Println("Error, required input:", finalError.Message)
	case *UserNotFoundError:
		fmt.Println("Error, user not found:", finalError.Message)
	default:
		fmt.Println("Error, unknown error:", finalError.Error())
	}
	} else {
		fmt.Println("Login berhasil")
	}

}