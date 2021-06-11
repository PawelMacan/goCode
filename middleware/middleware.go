package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PawelMacan/ticketProvider/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "****"
	dbname   = "ticketProvider"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected successfully!")
	return db
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var event model.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Fatalf("Usale to decode the request boddy. %v", err)
	}
	inserID := insertEvent(event)

	res := response{
		ID:      inserID,
		Message: "Ticket created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	event, err := getEvent(int64(id))
	if err != nil {
		log.Fatalf("Unable to get ticket. %v", err)
	}
	json.NewEncoder(w).Encode(event)
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	event, err := getAllEvents()

	if err != nil {
		log.Fatalf("Unable to get all tickets. %v", err)
	}

	json.NewEncoder(w).Encode(event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var event model.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updateRows := updateEvent(int64(id), event)

	msg := fmt.Sprintf("Ticket updated successfully. Total rows/record affected %v", updateRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows := deleteEvent(int64(id))

	msg := fmt.Sprintf("Ticket updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var ticket model.Ticket

	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Fatalf("Usale to decode the request boddy. %v", err)
	}
	inserID := insertTicket(ticket)

	res := response{
		ID:      inserID,
		Message: "Ticket created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	ticket, err := getTicket(int64(id))
	if err != nil {
		log.Fatalf("Unable to get ticket. %v", err)
	}
	json.NewEncoder(w).Encode(ticket)
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tickets, err := getAllTickets()

	if err != nil {
		log.Fatalf("Unable to get all tickets. %v", err)
	}

	json.NewEncoder(w).Encode(tickets)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var ticket model.Ticket
	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updateRows := updateTicket(int64(id), ticket)

	msg := fmt.Sprintf("Ticket updated successfully. Total rows/record affected %v", updateRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows := deleteTicket(int64(id))

	msg := fmt.Sprintf("Ticket updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

//-------Priveate Functions--------------

func insertTicket(ticket model.Ticket) int64 {
	db := createConnection()
	defer db.Close()
	if ticket.SellingOption.IsValid() == false {
		log.Fatalf("Selling option is not valid: %v", ticket.SellingOption)
	}

	sqlStatement := `INSERT INTO tickets (name, price, sellingOption) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, ticket.Name, ticket.Price, ticket.SellingOption).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getTicket(id int64) (model.Ticket, error) {
	db := createConnection()
	defer db.Close()
	var ticket model.Ticket

	sqlStatement := `SELECT * FROM tickets WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&ticket.Id, &ticket.Name, &ticket.Price, &ticket.SellingOption)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return ticket, nil
	case nil:
		return ticket, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return ticket, err
}

func getAllTickets() ([]model.Ticket, error) {
	db := createConnection()
	defer db.Close()
	var tickets []model.Ticket

	sqlStatement := `SELECT * FROM tickets`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.Id, &ticket.Name, &ticket.Price, &ticket.SellingOption)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, err
}

func updateTicket(id int64, ticket model.Ticket) int64 {
	db := createConnection()
	defer db.Close()
	if ticket.SellingOption.IsValid() == false {
		log.Fatalf("Selling option is not valid: %v", ticket.SellingOption)
	}
	sqlStatement := `UPDATE tickets SET name=$2, price=$3, sellingOption=$4 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, ticket.Name, ticket.Price, ticket.SellingOption)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteTicket(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM tickets WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func insertEvent(event model.Event) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO events (name, date) VALUES ($1, $2) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, event.Name, event.Date).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getEvent(id int64) (model.Event, error) {
	db := createConnection()
	defer db.Close()
	var event model.Event

	sqlStatement := `SELECT * FROM events WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&event.Id, &event.Name, &event.Date)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return event, nil
	case nil:
		return event, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return event, err
}

func getAllEvents() ([]model.Event, error) {
	db := createConnection()
	defer db.Close()
	var events []model.Event

	sqlStatement := `SELECT * FROM events`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var event model.Event
		err := rows.Scan(&event.Id, &event.Name, &event.Date)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		events = append(events, event)
	}
	return events, err
}

func updateEvent(id int64, event model.Event) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE events SET name=$2, date=$3 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, event.Name, event.Date)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteEvent(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM events WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
