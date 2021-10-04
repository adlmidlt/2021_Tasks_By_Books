package ex12

import (
	"log"
	"todo/todo/ex12/my_directory"
)

func FoodMain() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		if err := fridge.FindFood(food); err != nil {
			log.Fatal(err)
		}
	}
}

func DirMain() {
	my_directory.DirMain()
}
