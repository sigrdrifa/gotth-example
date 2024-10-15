package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sigrdrifa/gotth-example/internal/store"
	"github.com/sigrdrifa/gotth-example/internal/templates"
)

type GuestStore interface {
	AddGuest(guest store.Guest) error
	GetGuests() ([]store.Guest, error)
}

type server struct {
	logger     *log.Logger
	port       int
	httpServer *http.Server
	guestDb    GuestStore
}

// Creat a new server instance with the given logger and port
func NewServer(logger *log.Logger, port int, guestDb GuestStore) (*server, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	if guestDb == nil {
		return nil, fmt.Errorf("guestDb is required")
	}
	return &server{
		logger:  logger,
		port:    port,
		guestDb: guestDb}, nil
}

// Start the server
func (s *server) Start() error {
	s.logger.Printf("Starting server on port %d", s.port)
	var stopChan chan os.Signal

	// define router
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	router.HandleFunc("GET /", s.defaultHandler)
	router.HandleFunc("GET /about", s.aboutHandler)
	router.HandleFunc("GET /health", s.healthCheckHandler)
	router.HandleFunc("POST /guests", s.addGuestHandler)
	router.HandleFunc("GET /guests", s.getGuestsHandler)

	// define server
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router}

	// create channel to listen for signals
	stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error when running server: %s", err)
		}
	}()

	<-stopChan

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error when shutting down server: %v", err)
		return err
	}
	return nil
}

// GET /
func (s *server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	homeTemplate := templates.Home()
	err := templates.Layout(homeTemplate, "Spooktober party", "/").Render(r.Context(), w)
	if err != nil {
		s.logger.Printf("Error when rendering home: %v", err)
	}
}

func (s *server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	aboutTemplate := templates.About()
	err := templates.Layout(aboutTemplate, "About", "/about").Render(r.Context(), w)
	if err != nil {
		s.logger.Printf("Error when rendering about: %v", err)
	}
}

// GET /health - HealthCheckHandler is a simple handler to check the health of the server
func (s *server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// AddGuestHandler is a handler to add a guest to the guest store
func (s *server) addGuestHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Printf("Error when reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	guest, err := store.DecodeGuest(payload)
	if err != nil {
		s.logger.Printf("Error when decoding guest: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.guestDb.AddGuest(guest); err != nil {
		s.logger.Printf("Error when adding guest: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetGuestsHandler is a handler to get all guests from the guest store
func (s *server) getGuestsHandler(w http.ResponseWriter, r *http.Request) {
	guests, err := s.guestDb.GetGuests()
	if err != nil {
		s.logger.Printf("Error when getting guests: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	for _, guest := range guests {
		w.Write([]byte(fmt.Sprintf("Name: %s, Email: %s, Bringing: %s\n", guest.Name, guest.Email, guest.BringingItems)))
	}
}
