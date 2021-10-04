package task2

import (
	"fmt"
	"math/rand"
	"time"
)

// MainTask2
/* Написать генератор случайных чисел.

Итого: В принципе, легкая задача, на базовые знания по асинхронному взаимодействию в Go.
	Для решения я бы использовал небуфферезированный канал. Будем асинхронно писать туда случайные
	числа и закроем его, когда закончим писать.

	Плюс ее можно использовать в немного измененном виде в задаче на слияние N канала.
*/

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()
	return out
}

func MainTask2() {
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}
}
