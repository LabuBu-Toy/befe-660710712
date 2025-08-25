package week5_lab1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type User struct {
	ID string `json:"id"`
	Name string `json:"name"`

}
func main() {
	app := gin.Default()

	app.GET("/users", func (c *gin.Context){
		user := []User{
			{ID: "1", Name: "Thanawit"},
		}
		c.JSON(http.StatusOK, user)
	})
	app.Run(":8080")
}
