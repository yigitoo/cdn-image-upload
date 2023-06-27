package lib

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var config Config = NewConfig()
var iPORT uint16 = config.GetPort()
var PORT = fmt.Sprintf(":%s", strconv.Itoa(int(iPORT)))

func SetupApi() *gin.Engine {
	InitializeLogger()
	InfoLogger.Println("===PROGRAM_STARTED===")

	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})
	r.Static("/public/", "./public")
	r.LoadHTMLGlob("./templates/*.html")

	InfoLogger.Println("Routes are setting into application.")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"content": "This is an login page.",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var request User

		if err := ctx.BindJSON(&request); err != nil {
			panic(err)
		}

		user, err := ValidateLogin(request.Username, request.Password)
		LogError(err)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"user_id":  user.UserID,
				"username": user.Username,
			})
			// i use return because of
			return
		}
		// this context
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "We aren't able to login to you because of credentials.",
		})

	})

	r.POST("/imageUpload", func(ctx *gin.Context) {
		var image_upload ImageUpload
		if err := ctx.ShouldBind(&image_upload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		InfoLogger.Println("Image is being uploaded.")
		InfoLogger.Println(image_upload)

		can_pass, assoc := FileNameAnalyzer(image_upload.Image.Filename)
		if can_pass {
			InfoLogger.Println("File name is valid.")
			ctx.SaveUploadedFile(image_upload.Image, "public/uploads/"+RandomIDGenerator()+assoc)
			ctx.Redirect(http.StatusMovedPermanently, "/")
		}
	})
	r.GET("/id/:userid", func(ctx *gin.Context) {
		user_id := ctx.Params.ByName("userid")
		user, err := GetUserByID(user_id)
		LogError(err)

		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"username": user.Username,
		})
	})

	return r
}
