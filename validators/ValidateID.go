package validators

import (
	"fmt"
	"pollsbackend/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateID(id uint) (bool,error){

    idStr := strconv.FormatUint(uint64(id), 10)
    coefficients := []float64{2, 1, 2, 1, 2, 1, 2, 1, 2}

    if len(idStr) < 2 {
        fmt.Println("ID must have at least 2 digits.")
        return false, fmt.Errorf("ID must have at least 2 digits")
    }
    firstTwoDigitsStr := idStr[:2]
    firstTwoDigits, err := strconv.Atoi(firstTwoDigitsStr)
    if err != nil {
        fmt.Println("Error converting first two digits to integer:", err)
        return false,fmt.Errorf("error converting two digits to integer")
    }
    if firstTwoDigits <= 24 {
        fmt.Printf("The first two digits (%d) are less than or equal to 24.\n", firstTwoDigits)
    } else {

        fmt.Printf("The first two digits (%d) are greater than 24.\n", firstTwoDigits)
        return false,nil
    }

    third_digitstr := idStr[2:3]


    third_digit,err := strconv.Atoi(third_digitstr)

    if err != nil{
        fmt.Println("error",err)
        return false,fmt.Errorf("error: %w", err)
    }

    if third_digit >= 1 || third_digit <= 5{
        fmt.Println("correcto")
    } else{
        fmt.Println("cedula invalida digitos")
        return false,fmt.Errorf("error: %w", err)
    }


    results := make([]float64,9)

    for i:=0;i<9;i++{
        digit,err := strconv.Atoi(string(idStr[i]))
        if err != nil{
            fmt.Println("error de conversion",err)
            return false,fmt.Errorf("error: %w", err)
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
        return false,fmt.Errorf("error: %w", err)

    }
    fmt.Println("last digit",lastdigit)
    fmt.Println("final reult",finalResult)
    if float64(lastdigit) == finalResult{
        fmt.Println("verificacion de digitos exitosa")
    } else{
        if (finalResult == 10 && lastdigit == 0){
            return true,fmt.Errorf("error: %w", err)
        } else {return false,fmt.Errorf("error: %w", err)}
    }
    return true,fmt.Errorf("error: %w", err)

}


type CreateWalletRequest struct {
    UserID string `json:"userID"`
}

func CreateUserWallet(c *gin.Context) error {

    
    var request CreateWalletRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        fmt.Println("error binding JSON", err)
        //return "", fmt.Errorf("invalid request body: %v", err)
    }
    id := request.UserID
    uintid, err := strconv.ParseUint(id, 10, 0) // Base 10, bit size 0 (auto)
    if err != nil {
        return fmt.Errorf("failed to convert ID to uint: %v", err)
    }
    valid,err := ValidateID(uint(uintid))
    if err != nil{
        return err
    }
    if valid{
        util.CreateWallet(uint(uintid))
        return nil
    } else {
        return fmt.Errorf("ID validation failed")

    }

}