package main

import (
	"backend-webUE/api"
	"backend-webUE/config"
	"backend-webUE/database"
	"backend-webUE/models"
	"backend-webUE/router"
	"backend-webUE/services"
	"backend-webUE/utils"
	"context"
	"fmt"
	"log"
)

func main() {
	// Đọc cấu hình
	mongoConfig, serverConfig := config.LoadConfig()

	// Kết nối đến MongoDB
	db, err := database.Connect(mongoConfig)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = db.Client().Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Initialize OperatorConfig
	operatorConfig := &utils.OperatorConfig{
		PlmnId: models.PlmnId{
			Mcc: "208",
			Mnc: "93",
		},
		Amf: "8000",
		UeConfiguredNssai: []models.Snssai{
			{
				Sst: 1,
				Sd:  "010203",
			},
		},
		UeDefaultNssai: []models.Snssai{
			{
				Sst: 1,
				Sd:  "010203",
			},
		},
		// Add other necessary configuration
	}
	// Create Operator
	operator := utils.NewOperator(operatorConfig)

	// Khởi tạo services
	ueProfileService := services.NewUeProfileService(db, operator)

	// Khởi tạo API
	ueProfileAPI := api.NewUeProfileAPI(ueProfileService)

	// Khởi tạo router
	router := router.SetupRouter(ueProfileAPI, serverConfig) // Truyền ueProfileAPI và serverConfig vào SetupRouter

	// Chạy web server
	err = router.Run(fmt.Sprintf(":%d", serverConfig.Port))
	if err != nil {
		log.Fatalf("failed to run web server: %v", err)
	}
}
