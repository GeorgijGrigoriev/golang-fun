package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input int
	var star string
	l := int(len(`Осторожно, число больше 1 млн. может сильно загрузить процессор.`))
	for i := 0; i <= l; i++ {
		star += `*`
	}
	fmt.Println(star)
	fmt.Println(" Проверка случайности по числу иттераций.\n Осторожно, число больше 1 млн. может сильно загрузить процессор.")
	fmt.Println(" Введите число иттераций: ")
	fmt.Println(star)
	fmt.Scanf("%d", &input)
	test(input)
}

func test(r int) {
	start := time.Now()
	var (
		a = 0
		b = 0
	)
	for i := 0; i <= r; i++ {
		x := rand.Intn(2)
		if x == 0 {
			a++
		} else {
			b++
		}
	}
	aa := (float64(a) / float64(r)) * 100
	bb := (float64(b) / float64(r)) * 100
	elapsed := time.Since(start)
	fmt.Println("Вероятность 0: ", aa, "%, вероятность 1: ", bb, "%")
	fmt.Println("Выполнение программы заняло: ", elapsed)
}

func currentTime() {
	for range time.Tick(time.Second*1){
	fmt.Println(time.Now())
	}
}
