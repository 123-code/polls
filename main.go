
package main

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pollsbackend/controllers"
	"pollsbackend/initializers"
//	"pollsbackend/util"

)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB



func init(){
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	//util.InitializeUserWallet("0x0506208DC8461d22f964AD7ee223cbD09e10980A","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")
	//util.MintNFTWithExecute("0x3B57EAc775f5D2711572c05DedA51f8D5341202c","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")
	//util.ValidateWallet()
	//util.CreateWalet()
	//util.MintNFT("1804072310")
	//util.MintNFTWithExecute("0x858581A5c619bA15f21C23598aB74e1e317ABECc","0xbD6AAD0E7B72cFD2f7338b39d9047B1c3837266b")\
	//util.AddCandidate("daniel noboa","526938daf3a62f82fc13d7abe8d063104160bfd869ddbc25e3feb6a2f8a8042e")
	//util.GetLeader()
	//util.GetAllCandidates()
	
	voteController := controllers.NewVoteController(DB)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://localhost:3000/vote",
		"https://cellariusec-cellarius-web-store.vercel.app",
		"https://cellariusec-cellarius-web-store-icu5c4pzw-cellarius-projects.vercel.app",
		"https://cellariusec-cellarius-web-store-git-main-cellarius-projects.vercel.app",
		"http://localhost:8080",
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Cookie", "Usertype", "X-Forwarded-For"}

	r := gin.Default()
	r.Use(cors.New(config))
	r.POST("/validateid", controllers.EnterUser)





//curl -X GET "http://localhost:8080/sort?candidate_id=1&sort=oldest" -H "Content-Type: application/json"
	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	r.POST("/vote/:id",voteController.CastVote)
	r.POST("/registercandidate",voteController.RegisterCandidate)
	r.GET("/votes",controllers.AnalyzeVotersByProvince)
    r.GET("/vote_counts", voteController.GetVoteCounts)
	r.GET("/vote_data", voteController.GetAllVotes)
	r.GET("/sort",voteController.SortVotes)
	r.GET("/vote_data/all", voteController.GetAllVotes)
	r.GET("/obras/ratio", controllers.GetObrasRatio)
    r.POST("/obras/initialize", controllers.InitializeObras)
	r.GET("/obras_analysis", controllers.AnalyzeObrasByProvince)
	r.POST("/obras", controllers.CreateObra)
	//r.POST("/events", controllers.CreateEvent)
    r.GET("/events/impact", controllers.AnalyzeEventImpactByProvince)
	r.Run(":8080")
	

	//util.ValidateWallet();
	//util.InitializeUserWallet("0x858581A5c619bA15f21C23598aB74e1e317ABECc","0xC8ba9fBF6AA9A285D02912a25531B17006039717")
	//util.MintNFT()
	//util.VerifyContract()
	//util.ValidataWallet()

	
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
