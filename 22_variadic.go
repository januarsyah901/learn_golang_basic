package main


import "fmt"

// variadic function
func sumAll(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
	// sum varargs
	sum := sumAll(10,10)
	fmt.Println(sum)

	// sum slice
	// numbers := []int{1,2,3,4,5,6}
	// sum2 := sum(numbers...)
	// fmt.Println(sum2)
	// aneh kok error
}