package main

import (
    "fmt"

    "internal/utils"
)

func main() {
    m := utils.Max([]int{2, 4, 4, 5, 1, 0})

    fmt.Println(m)
    idx := utils.MaxIdx([]int{2, 4, 4, 5, 1, 0})

    fmt.Println(idx)
}
