package main



import (

    "fmt"

    "math/rand"

    "time"

)



func random(min int64, max int64) int64 {

    return rand.Int63n(max-min) + min

}



func BubbleSort(arr [100000]int64) [100000]int64 {

    for i := 0; i < 100000-1; i++ {

        for j := 0; j < 100000-i-1; j++ {

            if arr[j] > arr[j+1] {

                arr[j], arr[j+1] = arr[j+1], arr[j]

            }

        }

    }

    for i := 0; i < 100000-1; i++ {

        fmt.Printf(" %d ", arr[i])

    }

    return arr

}



func main() {

    var ar [100000]int64



    for i := 0; i < len(ar); i++ {



        randomNum := random(-100000, 100000)

        ar[i] = randomNum

        //fmt.Printf(" %d ", ar[i])

    }

    t0 := time.Now()

    BubbleSort(ar)

    fmt.Println(time.Since(t0))



}

