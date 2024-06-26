package main

import (
	"georeportapi/config"
	"georeportapi/controller"
	"georeportapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	defer config.CloseDb()

	router := gin.Default()
	v1 := router.Group("/georeport")
	{
		homepage := v1.Group("/homepage")
		{
			homepage.GET("/", controller.Homepage) // SEM AUTH
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login) //
		}

		report := v1.Group("/report")
		{
			report.GET("/", middleware.Authorized(), controller.GetMyReports)                                       // SEM AUTH
			report.GET("/:id", controller.GetReport)                                       // SEM AUTH
			report.POST("/reporting", middleware.Authorized(), controller.InsertReport)    // DTO IN -> AUTH -> DTO OUT
			report.PUT("/update/:id", middleware.Authorized(), controller.UpdateReport)    // DTO IN -> AUTH + OWNER -> DTO OUT
			report.DELETE("/delete/:id", middleware.Authorized(), controller.DeleteReport) // AUTH + OWNER
		}

		user := v1.Group("/user")
		{
			user.GET("/:id", middleware.Authorized(), controller.Profile)                 // AUTH - OWNER -> DTO RESPONSE (ID, NAME, EMAIL, PROFILE PICTURE)
			user.POST("/registration", controller.Register)                               // SEM AUTH -> DTO IN -> DTO OUT
			user.PUT("/update/:id", middleware.Authorized(), controller.UpdateProfile)    // AUTH - OWNER
			user.DELETE("/delete/:id", middleware.Authorized(), controller.DeleteAccount) // AUTH - OWNER
		}

		admin := v1.Group("/admin")
		{
			admin.GET("/reports", middleware.Authorized(), controller.GetAllReports) // AUTH -> DTO RESPONSE (ID, NAME, EMAIL)
		}
		router.Run(":3000")
	}
}
