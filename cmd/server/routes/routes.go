package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine  *gin.Engine
	Tickets []domain.Ticket
}

func NewRouter(engine *gin.Engine, tickets []domain.Ticket) *Router {
	return &Router{
		Engine:  engine,
		Tickets: tickets}
}

func (router *Router) MapRoutes() {

	repository := tickets.NewRepository(router.Tickets)
	service := tickets.NewService(repository)
	ticketHandler := handler.NewService(service)

	router.Engine.GET("/tickets/getByCountry/:dest", ticketHandler.GetTicketsByCountry())
	router.Engine.GET("/tickets/getAverage/:dest", ticketHandler.AverageDestination())

}
