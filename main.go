package main

import "fmt"

func main() {
	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	// fromTuesdayToThursDaySlice[2] = 0   изменение элементов слайса

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice))
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))
	fmt.Println(weekTempArr)

	/* Уменьшение размера слайсла
	weekTempSlice = weekTempSlice[:len(weekTempSlice)-1]
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))
	*/

	/*Увеличение размера слайса с созданием нового базового массива
	newSlice := append(weekTempSlice, 8)
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	*/

	/*Увеличение размера слайса с тем же базовы массивом
	newSlice := append(workDaysSlice, 8)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice[2] = 0
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	fmt.Println(weekTempArr)
	*/

	//Задание 8
	//Создайте слайс и заполните его числами от 1 до 100.
	//Оставьте в слайсе первые и последние 10 элементов и разверните слайс в обратном порядке.
	hundredSlice := make([]int, 100)
	for i := 0; i < 100; i++ {
		hundredSlice[i] = i + 1
	}
	fmt.Println(hundredSlice)
	if len(hundredSlice) != 0 {
		hundredSlice = append(hundredSlice[:10], hundredSlice[len(hundredSlice)-10:]...)
	}
	fmt.Println(hundredSlice)
	for i := range hundredSlice[:len(hundredSlice)/2] {
		hundredSlice[i], hundredSlice[len(hundredSlice)-i-1] = hundredSlice[len(hundredSlice)-i-1], hundredSlice[i]
	}
	fmt.Println(hundredSlice)

}
