package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type UserAccount struct {
	Id       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Pin      int    `json:"pin" bson:"pin"`
}

type BankAccount struct {
	Id    int    `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Owner int    `json:"ownerid" bson:"ownerid"`
}

type Bucket struct {
	Id    int    `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Owner int    `json:"ownerid" bson:"ownerid"`
}

type LineItem struct {
	Id          int     `json:"id" bson:"id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
	Bucket      int     `json:"bucket" bson:"bucket"`
	Bank        int     `json:"bank" bson:"bank"`
	Owner       int     `json:"owner" bson:"pwner"`
}

const (
	DB_USER     = "admin"
	DB_PASSWORD = "admin"
	DB_NAME     = "goproject"
)

func db_init() *sql.DB {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}

	fmt.Println("Connected to Database!")

	return db
}

func main() {
	http.ListenAndServe(":9000", handler())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/users" {
			userProcess(w, r)
		} else if r.URL.Path == "/banks" {
			bankProcess(w, r)
		} else if r.URL.Path == "/buckets" {
			bucketProcess(w, r)
		} else if r.URL.Path == "/lineitems" {
			lineitemProcess(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/user/%d", &id); n == 1 {
			userProcessId(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/bank/%d", &id); n == 1 {
			bankProcessId(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/bucket/%d", &id); n == 1 {
			bucketProcessId(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/lineitem/%d", &id); n == 1 {
			lineitemProcessId(id, w, r)
		}
	}
}

func userProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		db := db_init()
		defer db.Close()

		var user UserAccount
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var newUserId int
		err := db.QueryRow(
			"INSERT INTO public.useraccount (username, \"name\", pin) VALUES($1, $2, $3) RETURNING id;",
			user.Username,
			user.Name,
			user.Pin,
		).Scan(&newUserId)

		//checkError(err)
		if err != nil {
			res := strings.Contains(string(err.Error()), "duplicate key value violates unique constraint")

			if res {
				http.Error(w, "Username already in use.", http.StatusForbidden)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Id = newUserId

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, username, \"name\", pin FROM public.useraccount;")

		checkError(err)

		var users []UserAccount
		for rows.Next() {
			var id, pin int
			var username, name string

			err = rows.Scan(&id, &username, &name, &pin)
			checkError(err)

			users = append(users, UserAccount{
				Id:       id,
				Username: username,
				Name:     name,
				Pin:      pin,
			})
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
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

func userProcessId(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, username, \"name\", pin FROM public.useraccount WHERE id=$1;", id)

		checkError(err)

		var users []UserAccount
		for rows.Next() {
			var id, pin int
			var username, name string

			err = rows.Scan(&id, &username, &name, &pin)
			checkError(err)

			users = append(users, UserAccount{
				Id:       id,
				Username: username,
				Name:     name,
				Pin:      pin,
			})
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "POST":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "PUT":
		db := db_init()
		defer db.Close()

		var user UserAccount
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var updatedId int
		err := db.QueryRow(
			"UPDATE public.useraccount SET username=$1, \"name\"=$2, pin=$3 WHERE id=$4 RETURNING id;",
			user.Username,
			user.Name,
			user.Pin,
			id,
		).Scan(&updatedId)

		checkError(err)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		db := db_init()
		defer db.Close()

		var user UserAccount
		err := db.QueryRow("DELETE FROM public.useraccount where id = $1 RETURNING id, username, \"name\", pin;", id).Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.Pin,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func bankProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		db := db_init()
		defer db.Close()

		var bank BankAccount
		if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var newBankId int
		err := db.QueryRow(
			"INSERT INTO public.bankaccount (\"name\", ownerid) VALUES($1, $2) RETURNING id;",
			bank.Name,
			bank.Owner,
		).Scan(&newBankId)

		checkError(err)

		bank.Id = newBankId

		if err := json.NewEncoder(w).Encode(bank); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, \"name\", ownerid FROM public.bankaccount;")

		checkError(err)

		var accounts []BankAccount
		for rows.Next() {
			var id, owner int
			var name string

			err = rows.Scan(&id, &name, &owner)
			checkError(err)

			accounts = append(accounts, BankAccount{
				Id:    id,
				Name:  name,
				Owner: owner,
			})
		}
		if err := json.NewEncoder(w).Encode(accounts); err != nil {
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

func bankProcessId(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, \"name\", ownerid FROM public.bankaccount WHERE id=$1;", id)

		checkError(err)

		var banks []BankAccount
		for rows.Next() {
			var id, ownerid int
			var name string

			err = rows.Scan(&id, &name, &ownerid)
			checkError(err)

			banks = append(banks, BankAccount{
				Id:    id,
				Name:  name,
				Owner: ownerid,
			})
		}
		if err := json.NewEncoder(w).Encode(banks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "POST":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "PUT":
		db := db_init()
		defer db.Close()

		var bank BankAccount
		if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var updatedId int
		err := db.QueryRow(
			"UPDATE public.bankaccount SET \"name\"=$1 WHERE id=$2 RETURNING id;",
			bank.Name,
			id,
		).Scan(&updatedId)

		checkError(err)

		if err := json.NewEncoder(w).Encode(bank); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		db := db_init()
		defer db.Close()

		var bank BankAccount
		err := db.QueryRow("DELETE FROM public.bankaccount where id = $1 RETURNING id,\"name\", ownerid;", id).Scan(
			&bank.Id,
			&bank.Name,
			&bank.Owner,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(bank); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func bucketProcess(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, \"name\" FROM public.bucket;")

		checkError(err)

		var buckets []Bucket
		for rows.Next() {
			var id, owner int
			var name string

			err = rows.Scan(&id, &name, &owner)
			checkError(err)

			buckets = append(buckets, Bucket{
				Id:    id,
				Name:  name,
				Owner: owner,
			})
		}
		if err := json.NewEncoder(w).Encode(buckets); err != nil {
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

func bucketProcessId(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "PUT":
	case "DELETE":
	}
}

func lineitemProcess(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	case "GET":
		db := db_init()
		defer db.Close()

		rows, err := db.Query("SELECT id, title, description, amount, bucket, bank, \"owner\" FROM public.lineitem;")

		checkError(err)

		var lineitem []LineItem
		for rows.Next() {
			var id, bucket, bank, owner int
			var title, description string
			var amount float64

			err = rows.Scan(&id, &title, &description, &amount, &bucket, &bank, &owner)
			checkError(err)

			lineitem = append(lineitem, LineItem{
				Id:          id,
				Title:       title,
				Description: description,
				Amount:      amount,
				Bucket:      bucket,
				Bank:        bank,
				Owner:       owner,
			})
		}
		if err := json.NewEncoder(w).Encode(lineitem); err != nil {
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

func lineitemProcessId(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
		http.Error(w, "Not allowed!", http.StatusMethodNotAllowed)
		return
	case "PUT":
	case "DELETE":
	}
}
