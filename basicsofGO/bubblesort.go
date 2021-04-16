package main

import "fmt"


func main() {
    bubbleSort(toBeSorted)
}

var toBeSorted [10]int = [10]int{1,7,2,4,8,6,7,2,3,9}

func bubbleSort(input [10]int) {
    n := 10
    swapped := true
    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
           
            if input[i-1] > input[i] {
              
                input[i], input[i-1] = input[i-1], input[i]
              
                swapped = true
            }
        }
    }
    fmt.Println(input)
}

