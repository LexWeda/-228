package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addition() {
	var numberA int64
	var numberB int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Введи первую переменную: ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseInt(strings.TrimSpace(inputA), 10, 32)
	}

	fmt.Print("- Введи вторую переменную: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseInt(strings.TrimSpace(inputB), 10, 32)
	}

	fmt.Println("Ответ: ", numberA+numberB)
}

func substraction() {
	var numberA int64
	var numberB int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Введи первую переменную: ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseInt(strings.TrimSpace(inputA), 10, 32)
	}

	fmt.Print("- Введи вторую переменную: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseInt(strings.TrimSpace(inputB), 10, 32)
	}

	fmt.Println("Ответ: ", numberA-numberB)
}

func multiplication() {
	var numberA float64
	var numberB float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Введи первую переменную: ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseFloat(strings.TrimSpace(inputA), 32)
	}

	fmt.Print("- Введи вторую переменную: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseFloat(strings.TrimSpace(inputB), 32)
	}

	fmt.Println("Ответ : ", numberA*numberB)

}

func division() {
	var numberA float64
	var numberB float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("- Введи первую переменную: ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberA, _ = strconv.ParseFloat(strings.TrimSpace(inputA), 32)
	}

	fmt.Print("- Введи вторую переменную: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		numberB, _ = strconv.ParseFloat(strings.TrimSpace(inputB), 32)
	}

	fmt.Println("Ответ: ", numberA/numberB)

}

func main() {
	fmt.Println("Привет это мой проект calculator")
	fmt.Println("- Нажми 1 это сложение +")
	fmt.Println("- Нажми 2 это вычетание -")
	fmt.Println("- Нажми 3 это умножение *")
	fmt.Println("- Нажми 4 это деление /")
	fmt.Println("- Нажми 5 это выход")

	

	reader := bufio.NewReader(os.Stdin)

	inputOption, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	} else {
		options, _ := strconv.ParseInt(strings.TrimSpace(inputOption), 10, 64)
		if options == 1 {
			addition()
		}
		if options == 2 {
			substraction()
		}
		if options == 3 {
			multiplication()
		}
		if options == 4 {
			division()
		}
	}
}




