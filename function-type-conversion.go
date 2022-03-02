package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func PrintStyle1(message string) {
	fmt.Println("Style 1", message)
}

func PrintStyle2(message string) {
	fmt.Println("Style 2:", message)
}

type UserHandler struct {}

func (u UserHandler) Log(message string) {
	fmt.Println("User Handler:", message)
}

func printHelloWorld(logger Logger) {
	logger.Log("Hello World")
}

func main() {
	// The following is a function type conversion technique
	// Note that LoggerAdapter is a type (not a function) so LoggerAdapter(something) is not a function call. It is type conversion like int(y) (y := 9.5)
	
	// In the example below, Logger concrete types need to implement a single method Log()
	// 1) LoggerAdapter has implemented methods interface Logger required
	// 2) LoggerAdapter also has the same signature func(string) as implementation functions PrintStyle1, PrintStyle2, UserHandler.Log to be able to convert from these types
	// 3) LoggerAdapter also has the same signature func(string) as single required method Log() (because PrintStyle1, PrintStyle2, UserHandler.Log must implement this method)
	// So with these conditions, it can be used to convert something to a type required (in the example something is a function)
	style1 := LoggerAdapter(PrintStyle1)
	style2 := LoggerAdapter(PrintStyle2)
	style3 := LoggerAdapter(UserHandler{}.Log)
	printHelloWorld(style1)
	printHelloWorld(style2)
	printHelloWorld(style3)
}