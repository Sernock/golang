package main

import "fmt"

func main() {
    var count int
    fmt.Scan(&count)

    numbers := make([]int, count)
    for i := 0; i < count; i++ {
        fmt.Scan(&numbers[i])
    }

    flag := 1 
    i, j := 0, count-1

    for i < j {
        if numbers[i] != numbers[j] {
            flag = 0
            break
        }
        i++
        j--
    }

    if flag == 1 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
