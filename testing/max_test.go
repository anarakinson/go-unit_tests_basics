package main

import (
    "testing"

    "internal/utils"
)

func TestMax(t *testing.T) {
    // arrange
    numbers := []int{0, -1, 1, 3, -5}
    expected := 3

    // act
    result := utils.Max(numbers)

    //assert
    if result != expected {
        t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
    }
}

func TestMaxIdx(t *testing.T) {
    // arrange
    numbers := []int{0, -1, 3, -5}
    expected := 2

    // act
    result := utils.MaxIdx(numbers)

    //assert
    if result != expected {
        t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
    }
}
