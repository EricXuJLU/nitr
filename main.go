package main

import (
	rice "github.com/GeertJohan/go.rice"
	db "github.com/bitcav/nitr/database"
	"github.com/bitcav/nitr/handlers"
	service2 "github.com/bitcav/nitr/service"
	"github.com/bitcav/nitr/utils"
	"github.com/kardianos/service"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/embed"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/websocket"
)

func main() {
	interval := 30
	if len(os.Args) > 1 {
		atoi, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("invalid params")
		}
		interval = atoi
	}
	for i := 1; ; i++ {
		hostName := service2.GetHostName()
		service2.UpdateCPUStatus(hostName)
		service2.UpdateDiskStatus(hostName)
		service2.UpdateGPUStatus(hostName)
		service2.UpdateHostStatus(hostName)
		service2.UpdateIspStatus(hostName)
		service2.UpdateMemoryStatus(hostName)
		service2.UpdateNetworkStatus(hostName)
		service2.UpdateProcessStatus(hostName)
		service2.UpdateRAMStatus(hostName)
		log.Println("第"+strconv.Itoa(i)+"记录：", time.Now())
		time.Sleep(time.Second * time.Duration(interval))
	}
}

func server() {
	//Set Config.ini Default Values
	utils.ConfigFileSetup()

	//Set API Server default Data
	db.SetAPIData()

	//App Config
	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
	})

	//In Memory Static Assets
	app.Use("/assets", embed.New(embed.Config{
		Root: rice.MustFindBox("app/assets").HTTPBox(),
	}))

	//Checks if logs saving is activated
	utils.Logs(app)

	app.Use(recover.New(recover.Config{
		Handler: handlers.Recover,
	}))

	//API Config
	api := app.Group("/api")
	v1 := api.Group("/v1")

	//API Key auth middleware
	v1.Use(handlers.AuthAPI)

	//nitr API Endpoints
	v1.Get("/", handlers.Overview)
	v1.Get("/cpu", handlers.CPU)
	v1.Get("/bios", handlers.Bios)
	v1.Get("/bandwidth", handlers.Bandwidth)
	v1.Get("/chassis", handlers.Chassis)
	v1.Get("/disks", handlers.Disk)
	v1.Get("/drives", handlers.Drive)
	v1.Get("/devices", handlers.Devices)
	v1.Get("/gpu", handlers.GPU)
	v1.Get("/host", handlers.Host)
	v1.Get("/isp", handlers.ISP)
	v1.Get("/network", handlers.Network)
	v1.Get("/processes", handlers.Process)
	v1.Get("/ram", handlers.RAM)
	v1.Get("/baseboard", handlers.Baseboard)
	v1.Get("/product", handlers.Product)
	v1.Get("/memory", handlers.Memory)

	//Login View
	handlers.ViewsBox = rice.MustFindBox("app/views")
	app.Get("/", handlers.Login)

	//Login Submit
	app.Post("/", handlers.LoginSubmit)

	//Auth middleware
	app.Use(handlers.Auth)

	//Panel View
	app.Get("/panel", handlers.Panel)

	//Panel JSON Data
	app.Get("/content", handlers.PanelContent)

	//Panel Logout
	app.Post("/logout", handlers.Logout)

	//Generate new API Key
	app.Post("/generate", handlers.GenerateApiKey)

	//Change Password View
	app.Get("/password", handlers.Password)

	//New Password Submit
	app.Post("/password", handlers.PasswordSubmit)

	app.Get("/status", websocket.New(handlers.SocketReader))

	//Server startup
	utils.StartServer(app)
}

type program struct{}

var logger service.Logger

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	server()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

//
//func main() {
//	if len(os.Args) > 1 {
//		cmd.Execute()
//		return
//	}
//
//	svcConfig := &service.Config{
//		Name:        "NitrService",
//		Description: "A Remote Monitoring Tool for system information gathering, making it available through a JSON API.",
//	}
//
//	prg := &program{}
//	s, err := service.New(prg, svcConfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//	logger, err = s.Logger(nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = s.Run()
//	if err != nil {
//		logger.Error(err)
//	}
//}
