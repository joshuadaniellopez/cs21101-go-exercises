package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Contact struct {
	Id       int
	Last     string // Last Name
	First    string // First Name
	Company  string
	Address  string
	Country  string
	Position string
}

type Database struct {
	nextID int
	mu     sync.Mutex
	recs   []Contact
}

func main() {
	contactsdb := &Database{recs: []Contact{}}
	http.ListenAndServe(":8080", contactsdb.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processId(id, w, r)
		}
	}
}

func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var rec Contact
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.mu.Lock()

		duplicateFound := false
		existingContact := Contact{}
		for _, item := range db.recs {
			if item.First == rec.First && item.Last == rec.Last && item.Company == rec.Company && item.Address == rec.Address && item.Country == rec.Country && item.Position == rec.Position {
				duplicateFound = true
				existingContact = item
				break
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if duplicateFound {
			data, _ := json.Marshal(existingContact)
			http.Error(w, string(data), http.StatusConflict)
			return
		}

		rec.Id = db.nextID
		db.nextID++
		db.recs = append(db.recs, rec)
		db.mu.Unlock()

		data, _ := json.Marshal(rec)
		fmt.Fprintln(w, string(data))
		return

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.recs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "DELETE":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func (db *Database) processId(id int, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		notFound := true
		contact := Contact{}
		for _, item := range db.recs {
			if id == item.Id {
				notFound = false
				contact = item
			}
		}
		if notFound {
			http.Error(w, "Not found!", http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(contact); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "POST":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "PUT":
		notFound := true
		var contact int
		for j, item := range db.recs {
			if id == item.Id {
				notFound = false
				contact = j
			}
		}
		if notFound {
			http.Error(w, "Not found!", http.StatusNotFound)
			return
		}

		var rec Contact
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.mu.Lock()
		db.recs[contact].First = rec.First
		db.recs[contact].Last = rec.Last
		db.recs[contact].Company = rec.Company
		db.recs[contact].Address = rec.Address
		db.recs[contact].Country = rec.Country
		db.recs[contact].Position = rec.Position

		db.mu.Unlock()

		if err := json.NewEncoder(w).Encode(rec); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "DELETE":
		notFound := true
		for _, item := range db.recs {
			if id == item.Id {
				notFound = false
			}
		}
		if notFound {
			http.Error(w, "Not found!", http.StatusNotFound)
			return
		}

		db.mu.Lock()
		for j, item := range db.recs {
			if id == item.Id {
				db.recs = append(db.recs[:j], db.recs[j+1:]...)
				break
			}
		}
		db.mu.Unlock()
	}
}
