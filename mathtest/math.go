package main

import (
	"fmt"
	"strconv"
	"bufio"
)

func add(in1 string, in2 string) int64 {
	num1, _ := strconv.ParseInt(in1, 0, 64)
	num2, _ := strconv.ParseInt(in2, 0, 64)
	result := num1 + num2
	return result
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//reading an integer
    var age int
    fmt.Println("What is your age?")
    _, err: fmt.Scan(&age)

    //reading a string
    reader := bufio.newReader(os.Stdin)
    var name string
    fmt.Println("What is your name?")
    name, _ := reader.readString("\n")

    fmt.Println("Your name is ", name, " and you are age ", age)
	//fmt.Print("Enter text: ")
	//in2, _ := reader.ReadString('\n')
	//num1, _ := strconv.ParseInt(in1, 0, 64)
	//num2, _ := strconv.ParseInt("2", 0, 64)
	//result := num1 + num2
	//result := add(in1, in2)
	//fmt.Println(result)
}
