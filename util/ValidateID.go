package util

import (
    "fmt"
    "strconv"
)

func ValidateID(id uint) {

    idStr := strconv.FormatUint(uint64(id), 10)

    if len(idStr) < 2 {
        fmt.Println("ID must have at least 2 digits.")
        return
    }


    firstTwoDigitsStr := idStr[:2]
    firstTwoDigits, err := strconv.Atoi(firstTwoDigitsStr)
    if err != nil {
        fmt.Println("Error converting first two digits to integer:", err)
        return
    }

    // Check if the first two digits are less than or equal to 24
    if firstTwoDigits <= 24 {
        fmt.Printf("The first two digits (%d) are less than or equal to 24.\n", firstTwoDigits)
    } else {
        fmt.Printf("The first two digits (%d) are greater than 24.\n", firstTwoDigits)
    }
}


func main(){
	ValidateID(1804072310)
}