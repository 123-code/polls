package util

import (
    "fmt"
    "strconv"
)

func ValidateID(id uint) bool{

    idStr := strconv.FormatUint(uint64(id), 10)
    coefficients := []float64{2, 1, 2, 1, 2, 1, 2, 1, 2}

    if len(idStr) < 2 {
        fmt.Println("ID must have at least 2 digits.")
        return false
    }
    firstTwoDigitsStr := idStr[:2]
    firstTwoDigits, err := strconv.Atoi(firstTwoDigitsStr)
    if err != nil {
        fmt.Println("Error converting first two digits to integer:", err)
        return false
    }
    if firstTwoDigits <= 24 {
        fmt.Printf("The first two digits (%d) are less than or equal to 24.\n", firstTwoDigits)
    } else {

        fmt.Printf("The first two digits (%d) are greater than 24.\n", firstTwoDigits)
        return false
    }

    third_digitstr := idStr[2:3]


    third_digit,err := strconv.Atoi(third_digitstr)

    if err != nil{
        fmt.Println("error",err)
        return false
    }

    if third_digit >= 1 || third_digit <= 5{
        fmt.Println("correcto")
    } else{
        fmt.Println("cedula invalida digitos")
        return false
    }


    results := make([]float64,9)

    for i:=0;i<9;i++{
        digit,err := strconv.Atoi(string(idStr[i]))
        if err != nil{
            fmt.Println("error de conversion",err)
            return false
        }
        results[i] = float64(digit) * coefficients[i]

        if results[i] >= 10 {
            results[i] = results[i]-9
        }
    }
    totalsum := 0.0

    for _, result := range results{
        totalsum += result
    }
    fmt.Printf("Total sum of results: %.1f\n", totalsum)
    nextHighestTen := float64(int(totalsum/10) + 1) * 10
    finalResult := nextHighestTen - totalsum
    fmt.Println(finalResult)
    lastdigitstr := idStr[len(idStr)-1:]
    lastdigit,err := strconv.Atoi(lastdigitstr)

    if err != nil{
        fmt.Println("error de conversion",err)
        return false

    }
    fmt.Println("last digot",lastdigit)
    fmt.Println("final reult",finalResult)
    if float64(lastdigit) == finalResult{
        fmt.Println("verificacion de digitos exitosa")
    } else{
        if (finalResult == 10 && lastdigit == 0){
            return true
        } else {return false}
    }
    return true

}