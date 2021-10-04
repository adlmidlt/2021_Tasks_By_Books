package task2

import (
	"fmt"
	"strconv"
	"strings"
)

/*	Напишите функцию, которая принимает массив из 10 целых чисел (от 0 до 9),
	который возвращает строку этих чисел в форме номера телефона.
	Формат: (123) 456-7890
*/

// Мой вариант

func CreatePhoneNumber(numbers [10]uint) string {

	var strs string
	for _, num := range numbers {
		str := strconv.Itoa(int(num))
		strs += str
	}
	/*	var joinOne string
		var joinTwo string
		var joinThree string*/

	strSplit := strings.Split(strs, "")

	//for i := 0; i < len(strSplit); i++ {
	joinOne := strings.Join(strSplit[:3], "")
	joinTwo := strings.Join(strSplit[3:6], "")
	joinThree := strings.Join(strSplit[6:10], "")
	//}
	result := fmt.Sprintf("(%s) %s-%s", joinOne, joinTwo, joinThree)
	fmt.Println(result)
	return result
}

// Этот вариант был взять с codewars для изучения.

func ArrayToString(numbers interface{}) string {
	return strings.Trim(
		strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber2(numbers [10]uint) string {
	strSplit := ArrayToString(numbers)
	return fmt.Sprintf("(%s) %s-%s", strSplit[0:3], strSplit[3:6], strSplit[6:10])
}
