package task3

import (
	"fmt"
	"sync"
)

// MainTask3
/* 	Слить N каналов в один.

Даны n каналов типа chan int. Надо написать функцию, которая смержит все данные из этих каналов
в один и вернет его.

Мы хотим, чтобы результат работы функции выглядел примерно так:
for num := range joinChannels(a, b, c) {
	fmt.Println(num)
}

Для этого напишем функцию, которая будет асинхронно читать из исходных каналов, которые ей
передадут в качестве аргументов, и писать в результирующий канал, который вернется из функции.

1) Создаем канал, куда будем сливать все данные. Он будет небуфферезированный, потому что мы
не знаем, сколько данных придет из каналов.
2) Дальше асинхронно прочитаем из исходных каналов и закроем результирующий канал для мержа,
когда все чтение закончится.
3) Чтобы дождаться конца чтения, просто обернем этот цикл по каналам в wait group.
*/

func joinChannel(chs ...<-chan int) <-chan int {
	mergeCh := make(chan int)

	go func() {
		/*WaitGroup ожидает завершения набора горутин. Основная горутина вызывает Add, чтобы
		установить количество ожидаемых горутин. Затем каждая из горутин запускается и по
		завершении вызывает Done. В то же время Wait можно использовать для блокировки до
		завершения всех горутин. WaitGroup нельзя копировать после первого использования.*/
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergeCh <- id
				}
			}(ch, wg)
		}
		wg.Wait()
		close(mergeCh)
	}()
	return mergeCh
}

func MainTask3() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{10, 20, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannel(a, b, c) {
		fmt.Println(num)
	}
}
