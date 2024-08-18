package helper

var version = "09.4.3.5"
var Application = "Golang"

func SayHello(name string) string {
	return "Hello " + name + "!"
}

func Info() string {
	return version
}