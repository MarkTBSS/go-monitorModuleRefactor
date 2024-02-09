package servers

import (
	monitorhandlers "github.com/MarkTBSS/go-monitorModuleRefactor/modules/monitor/monitorHandlers"
	"github.com/gofiber/fiber/v2"
)

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorhandlers.MonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}
