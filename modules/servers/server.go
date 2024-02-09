package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/MarkTBSS/go-monitorModuleRefactor/config"
	"github.com/gofiber/fiber/v2"
)

func NewServer(cfg config.IConfig) IServer {
	return &server{
		cfg: cfg,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

type IServer interface {
	Start()
}

type server struct {
	cfg config.IConfig
	app *fiber.App
}

func (s *server) Start() {
	// Modules
	v1 := s.app.Group("v1")
	modules := InitModule(v1, s)
	modules.MonitorModule()

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
