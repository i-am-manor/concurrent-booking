package main

import (
	"log"
	"net/http"

	"github.com/i-am-manor/concurrent-booking/internal/adapters/redis"
	"github.com/i-am-manor/concurrent-booking/internal/booking"
	"github.com/i-am-manor/concurrent-booking/internal/utils"
)

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}

var movies = []movieResponse{
	{ID: "inception", Title: "Inception", Rows: 5, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune: Part Two", Rows: 4, SeatsPerRow: 6},
}

func main() {
	r := http.NewServeMux()
	r.Handle("GET /", http.FileServer(http.Dir("static")))
	r.HandleFunc("GET /movies", listMovies)

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)
	bookingHandler := booking.NewHandler(svc)

	r.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)
	r.HandleFunc("POST /movies/{movieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)
	r.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	r.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, movies)
}
