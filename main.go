
package main

import (
    "fmt"
    "strconv"
)

func ValidateID(id uint) {

    idStr := strconv.FormatUint(uint64(id), 10)
	coefficients := []float64{2, 1, 2, 1, 2, 1, 2, 1, 2}

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
    if firstTwoDigits <= 24 {
        fmt.Printf("The first two digits (%d) are less than or equal to 24.\n", firstTwoDigits)
    } else {
        fmt.Printf("The first two digits (%d) are greater than 24.\n", firstTwoDigits)
    }

	third_digitstr := idStr[2:3]


	third_digit,err := strconv.Atoi(third_digitstr)

	if err != nil{
		fmt.Println("error",err)
		return
	}

	if third_digit >= 1 || third_digit <= 5{
		fmt.Println("correcto")
	} else{
		fmt.Println("cedula invalida digitos")
	}

	//multiply and store results( first 9 id numbers, and coefficients)
	results := make([]float64,9)

	for i:=0;i<9;i++{
		digit,err := strconv.Atoi(string(idStr[i]))
		if err != nil{
			fmt.Println("error de conversion",err)
			return
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
		return 

	}
	if float64(lastdigit) == finalResult{
		fmt.Println("verificacion de digitos exitosa")
	} else{
		fmt.Println("verificacion de digitos fallida")
	}
	
}


func main(){
	ValidateID(1803047263)
}

/*package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func init() {
	os.Setenv("DB_CONNECTION_STRING", "postgresql://peso_core_db_user:0rU71sCZhy2lcK0uK8YhpKFfBhY83iOZ@dpg-crk9csm8ii6s73eju4m0-a.oregon-postgres.render.com/peso_core_db")
	os.Setenv("ISSUER", "http://localhost:8080")
	os.Setenv("SECRET", "secret")
	os.Setenv("JWT_SECRET", "secret")
	os.Setenv("AUDIENCE", "http://localhost:5000")
	/*
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
	*/

