package main

import (
	"bioskop/controllers"
	database "bioskop/db"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type Bioskop struct {
// 	ID     int     `json:"id"`
// 	Nama   string  `json:"nama"`
// 	Lokasi string  `json:"lokasi"`
// 	Rating float64 `json:"rating"`
// }

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	host, port, user, password, dbname,
	// )

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()
	router.GET("/bioskop", controllers.GetAllBioskop)
	router.POST("/bioskop", controllers.InsertBioskop)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)
	// db.Connect()
	// router := gin.Default()

	// router.POST("/bioskop", func(ctx *gin.Context) {
	// 	// var newBioskop Bioskop

	// 	// if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
	// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	// 	return
	// 	// }

	// 	// if newBioskop.Nama == "" || newBioskop.Lokasi == "" {
	// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 	// 		"error": "Name and Location cant be empty",
	// 	// 	})
	// 	// 	return
	// 	// }

	// 	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	// 	var lastInsertID int
	// 	err := db.DB.QueryRow(query, newBioskop.Nama, newBioskop.Lokasi, newBioskop.Rating).Scan(&lastInsertID)

	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to get ID"})
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		}
	// 		return
	// 	}

	// 	newBioskop.ID = lastInsertID
	// 	ctx.JSON(http.StatusCreated, gin.H{
	// 		"message": "Bioskop added",
	// 		"data":    newBioskop,
	// 	})
	// })

	router.Run(":8080")
}
