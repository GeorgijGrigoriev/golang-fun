package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    var input int
    fmt.Println("Проверка случайности по числу иттераций.\n Осторожно, число больше 1 млн. может сильно загрузить процессор.")
    fmt.Println("Введите число иттераций: ")
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
        if (x == 0){
            a++
        } else {
            b++
        }
    }
    aa := (float64(a) / float64(r)) * 100
    bb := (float64(b) / float64(r)) * 100
    elapsed := time.Since(start)
    fmt.Println("Вероятность 0: ", aa ,"%, вероятность 1: ",bb,"%")
    fmt.Println("Выполнение программы заняло: ", elapsed)
}
