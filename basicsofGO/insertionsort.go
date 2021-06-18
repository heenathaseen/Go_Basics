package main



import (

    "fmt"

    "math/rand"

    "time"

)



func random(min int64, max int64) int64 {

    return rand.Int63n(max-min) + min

}



func InsertionSort(numbers [100000]int64) [100000]int64 {

    for i := 0; i < len(numbers); i++ {

        for j := 0; j < i+1; j++ {

            //compare element present at index i with every element present

            //  left of it place it in right place so that array on the

            //left remains   sorted

            if numbers[j] > numbers[i] {

                intermediate := numbers[j]

                numbers[j] = numbers[i]

                numbers[i] = intermediate

            }

        }

        //fmt.Println(numbers)

    }

    return numbers

}

func main() {



    var ar [100000]int64



    for i := 0; i < len(ar); i++ {

        //rand.Seed(time.Now().UnixNano())

        randomNum := random(-10000, 10000)

        ar[i] = randomNum

        //fmt.Printf(" %d ", ar[i])

    }



    t0 := time.Now()

    InsertionSort(ar)

    fmt.Println(time.Since(t0))

}

