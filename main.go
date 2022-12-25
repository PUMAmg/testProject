package main

import "fmt"

func main() {
	prices := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	fmt.Println("Список деликатесов:")
	for k, v := range prices {
		if v >= 500 {
			fmt.Println(k)
		}
	}
	orderSum := 0
	for _, v := range order {
		for k, vm := range prices {
			if v == k {
				orderSum += vm
			}
		}
	}
	fmt.Printf("Стоимость заказа - %d", orderSum)
}
