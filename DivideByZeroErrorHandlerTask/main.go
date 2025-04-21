package main

import (
	"fmt"
	"sync"
)

type TwoNums struct {
	FirstOp  int
	SecondOp int
}

func main() {
	operationSlice := []TwoNums{
		{10, 2},
		{5, 0},
		{8, 4},
		{3, 0},
	}

	var wg sync.WaitGroup
	wg.Add(len(operationSlice))

	for _, op := range operationSlice {
		go Divide(&wg, op)
	}

	wg.Wait()
}

func Divide(wg *sync.WaitGroup, operands TwoNums) {
	defer func() {
		if err := recover(); err != nil {
			PrintDivide(operands, err)
		}
	}()
	defer wg.Done()

	if operands.SecondOp == 0 {
		panic("ошибка: деление на ноль") // Здесь при помощи fmt.Sprintf() можно было бы че то ещё поделать добавив переменные, но это лишнее. Просто как возможность пишу.
	}
	result := operands.FirstOp / operands.SecondOp
	PrintDivide(operands, result)
}

func PrintDivide(operands TwoNums, v interface{}) {
	switch v.(type) { // Здесь используется так называемый type switch. Это в каких-то моментах удобнее type assertion
	case string:
		fmt.Printf("%d / %d => %v\n", operands.FirstOp, operands.SecondOp, v) // Здесь можно было бы реализовать добавление => в моментах с ошибкой через использование switch-case и проверку состояние типа данных v. Но это лишний код. Хотя нет, добавлю.
	default:
		fmt.Printf("%d / %d = %v\n", operands.FirstOp, operands.SecondOp, v)
	}
}
