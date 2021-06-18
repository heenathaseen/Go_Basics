package main



import (

    "fmt"

    "math/rand"

    "time"

)



func random(min int64, max int64) int64 {

    return rand.Int63n(max-min) + min

}



func selection_sort(arr [100000]int64, size int) {

    for i := 0; i < size; i++ {

        min := i

        for j := i + 1; j < size; j++ {

            if arr[min] > arr[j] {

                min = j

            }



            temp := arr[min]

            arr[min] = arr[i]

            arr[i] = temp

        }



    }

}



func main() {



    var ar [100000]int64

    var size int

    fmt.Println("Enter the size of array :")

    fmt.Scan(&size)



    for i := 0; i < len(ar); i++ {

        //rand.Seed(time.Now().UnixNano())

        randomNum := random(-10000, 10000)

        ar[i] = randomNum

        //fmt.Printf(" %d ", ar[i])

    }



    t0 := time.Now()

    selection_sort(ar, size)

    fmt.Println(time.Since(t0))

}

