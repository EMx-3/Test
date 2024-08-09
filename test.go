package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func arabic_to_roman(x int) string { //поразрядный перевод арабского числа 1-100 в римскую систему
	first_digit := map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC", 10: "C"}
	last_digit := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 0: ""}
	return (first_digit[x/10] + last_digit[x%10])
}

func calc(x, y int, z string) int { //сам калькулятор для арабских чисел, учитывая, что операнд введен верно
	switch z {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	default:
		return x / y

	}
}

func main() {
	arabic := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}          //создаем мапу арабских чисел 1-10, чтобы убедиться что число принадлежит нужному диапазону и для дальнейшего перевод в целочисленный тип данных
	roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10} // аналогично для римских
	fmt.Println("Введите выражние")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)    //удаляем ненужные пробелы
	result := strings.Split(input, " ") // разбиваем ввод на строки пробелом
	if len(result) != 3 {               // проверяем введены ли 2 операнда и 1 оператор
		panic("Ввод неверен")
	}
	_, flaga := arabic[result[0]] //ищем первый операнд в мапе арабских чисел
	if flaga {
		_, flag2 := arabic[result[2]]
		if flag2 {
			fmt.Println(+calc(arabic[result[0]], arabic[result[2]], result[1])) //мы сюда зашли, значит все ок, считаем результат и выводим его
		} else {
			panic("Ввод неверен") // нашли первый операнд в арабской мапе, но не нашли второй
		}
	} else {
		_, flagr := roman[result[0]] //ищем первый операнд в мапе римских чисел
		if flagr {
			_, flag2 := roman[result[2]]
			if flag2 {
				answer := calc(roman[result[0]], roman[result[2]], result[1]) //оба операнда есть в мапе римских чисел, считаем ответ
				if answer < 1 {
					panic("В римской системы нет чисел меньше 1")
				}
				fmt.Println(arabic_to_roman(answer)) //выводим результат в римской системе
			} else {
				panic("Ввод неверен") //нашли первый операнд в римской мапе, но не нашли второй
			}
		} else {
			panic("Ввод неверен") // не нашли первый операнд ни в римской, ни в арабской мапах
		}
	}
}
