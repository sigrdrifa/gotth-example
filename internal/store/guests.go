// Description: This file contains the guestStore struct and its methods
// that are used to store and retrieve guests from the store.
// This is for example purposes only and just uses an in-memory Map
package store

import (
	"encoding/json"
	"fmt"
	"log"
)

type Guest struct {
	Name  string
	Email string
}

type guestStore struct {
	logger *log.Logger
	guests map[string]Guest
}

func NewGuestStore(logger *log.Logger) *guestStore {
	return &guestStore{
		logger: logger,
		guests: make(map[string]Guest),
	}
}

func (gs *guestStore) AddGuest(guest Guest) error {
	if guest.Email == "" {
		return fmt.Errorf("email is required")
	}
	// check if the key exists
	if _, ok := gs.guests[guest.Email]; ok {
		return fmt.Errorf("A spooky guest with email %s already exists", guest.Email)
	}
	gs.guests[guest.Email] = guest
	fmt.Printf("Added guest: %v \n", guest)
	return nil
}

func (gs *guestStore) GetGuests() ([]Guest, error) {
	if gs.guests == nil {
		return nil, fmt.Errorf("no guests found")
	}

	// create guests slice
	guests := make([]Guest, 0, len(gs.guests))
	for _, guest := range gs.guests {
		guests = append(guests, guest)
	}
	return guests, nil
}

func DecodeGuest(payload []byte) (Guest, error) {
	var guest Guest
	if err := json.Unmarshal(payload, &guest); err != nil {
		return Guest{}, fmt.Errorf("error decoding guest: %w", err)
	}
	return guest, nil
}
