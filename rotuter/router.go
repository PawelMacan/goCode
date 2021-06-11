package rotuter

import (
	"github.com/PawelMacan/ticketProvider/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Tickets endpoints
	router.HandleFunc("/api/ticket/{id}", middleware.GetTicket).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/ticket", middleware.GetAllTickets).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newTicket", middleware.CreateTicket).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/ticket/{id}", middleware.UpdateTicket).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/ticket/{id}", middleware.DeleteTicket).Methods("DELETE", "OPTIONS")

	// Events endpoints
	router.HandleFunc("/api/event/{id}", middleware.GetEvent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/event", middleware.GetAllEvents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newEvent", middleware.CreateEvent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/event/{id}", middleware.UpdateEvent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/event/{id}", middleware.DeleteEvent).Methods("DELETE", "OPTIONS")

	return router
}
