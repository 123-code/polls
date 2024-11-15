/*
package main

import (

	    "fmt"
	    "strconv"
		"pollsbackend/util"

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
		util.CreateWallet("1803047263")
	}
*/
package main

import (
	//"fmt"
	//"github.com/gin-contrib/cors"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"

	//"net/url"
	"pollsbackend/util"
	//"pollsbackend/controllers"
	//"pollsbackend/util"
	//"pollsbackend/util"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB

/*func init() {
	var err error
	encodedPassword := url.QueryEscape("AVNS_LNgImquHJXNIMn4aMTt")

	dsn := fmt.Sprintf("postgres://avnadmin:AVNS_Xv4rbmtFkoxKOGBDks_@pg-382a4bdb-udla-54df.g.aivencloud.com:18022/defaultdb?sslmode=require", encodedPassword)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
}
*/

func main() {
	//util.InitializeUserWallet("0x0506208DC8461d22f964AD7ee223cbD09e10980A","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")
	//util.MintNFTWithExecute("0x3B57EAc775f5D2711572c05DedA51f8D5341202c","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")
	//util.ValidateWallet()
	//util.CreateWalet()
	//util.MintNFT("1804072310")
	util.MintNFTWithExecute("0x858581A5c619bA15f21C23598aB74e1e317ABECc","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")
	/*
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"https://cellariusec-cellarius-web-store.vercel.app",
		"https://cellariusec-cellarius-web-store-icu5c4pzw-cellarius-projects.vercel.app",
		"https://cellariusec-cellarius-web-store-git-main-cellarius-projects.vercel.app",
		"http://localhost:8080",
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Cookie", "Usertype"}

	r := gin.Default()
	r.Use(cors.New(config))
	r.POST("/validateid", controllers.EnterUser)






	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
	

	//util.ValidateWallet();
	//util.InitializeUserWallet("0x858581A5c619bA15f21C23598aB74e1e317ABECc","0xC8ba9fBF6AA9A285D02912a25531B17006039717")
	//util.MintNFT()
	//util.VerifyContract()
	//util.ValidataWallet()
	*/
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	var user User
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	var user User
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	var user User
	if err := DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
