package main

import (
    "testing"

    "internal/utils"
)

func TestMax(t *testing.T) {
    // arrange
    testTable := []struct{
        numbers []int
        expected int
    }{
        {
            numbers : []int{0, -1, 1, 3, -5},
            expected : 3,
        },
        {
            numbers : []int{0, -1, -10, -5},
            expected : 0,
        },
        {
            numbers : []int{},
            expected : 0,
        },
    }

    // act
    for _, testCase := range(testTable) {
        result := utils.Max(testCase.numbers)
        t.Logf(
            "Calling Max(%v)\n expected: %d\n result: %d\n\n",
            testCase.numbers,
            testCase.expected,
            result,
        )

        //assert
        if result != testCase.expected {
            t.Errorf(
                "Incorrect result. Expected %d, got %d",
                testCase.expected,
                result,
            )
        }
    }
}

func TestMaxIdx(t *testing.T) {
    // arrange
    numbers := []int{0, -1, 3, 1, -5}
    expected := 2

    // act
    result := utils.MaxIdx(numbers)

    //assert
    if result != expected {
        t.Errorf("Incorrect result. Expected %d, got %d", expected, result)
    }
}
