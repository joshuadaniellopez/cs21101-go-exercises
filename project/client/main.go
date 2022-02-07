package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const ROOT_URL = "http://localhost:9000"

// Data Models - same with server
// User Account, Bank Account, Bucket, and Line Item
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
	Owner       int     `json:"ownerid" bson:"ownerid"`
}

func main() {
	// Client runs forever, until EOF
	for {
		welcome()

		var username string
		var name string
		var pin int

		// Ask user if want to create account
		// If yes, ask for User Account Properties
		var createAccount string
		fmt.Print("Would you like to create an account? (Y/N) ")
		fmt.Scan(&createAccount)

		if createAccount == "Y" {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("What's your name: ")
			if scanner.Scan() {
				name = scanner.Text()
			}

			fmt.Print("What's your desired name: ")
			if scanner.Scan() {
				username = scanner.Text()
			}

			fmt.Print("What's your desired pin: ")
			fmt.Scan(&pin)

			success := createUser(username, name, pin)
			if success {
				fmt.Println("User Account Created!")
				continue
			} else {
				fmt.Println("Invalid Request! Make sure to use an unused username and valid PIN.")
			}
		}

		authorized := false

		// Authorization: Check if valid user account
		fmt.Print("Enter your username: ")
		fmt.Scanf("%s", &username)

		fmt.Print("Enter your PIN: ")
		fmt.Scanf("%d", &pin)

		authorized = authorize(username, pin)

		if !authorized {
			unauthorizedDetected()
			authorized = false
			continue
		}

		authorizedUser := getUser(username, pin)
		fmt.Println(authorizedUser)

		var banks []BankAccount
		var buckets []Bucket
		var lineitems []LineItem

		// Retrieve all data related to user
		banks = getBanks(authorizedUser.Id)
		buckets = getBuckets(authorizedUser.Id)
		lineitems = getLineItems(authorizedUser.Id)

		// Run function for supported CRUD processes
		for {
			fmt.Println("-----------------------")
			process(authorizedUser.Id, &banks, &buckets, &lineitems)
			fmt.Println("-----------------------")
		}
	}

}

// Print Welcome Message
func welcome() {
	fmt.Println("-----------------------")
	fmt.Println("Welcome to Monefy	")
	fmt.Println("-----------------------")
}

// Print Unauthorized Message
func unauthorizedDetected() {
	fmt.Println("---------------------------")
	fmt.Println("Unauthorized. Try Again	")
	fmt.Println("---------------------------")
}

// Fire request to /authorize endpoint with user payload
// Returns boolean (true for success, false otherwise)
func authorize(username string, pin int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"username\": \"%s\", \"pin\": %d}", username, pin)
	payload := bytes.NewBuffer([]byte(body))

	// Fire request to /authorize endpoint with user payload
	response, err := client.Post(ROOT_URL+"/authorize", "application/json", payload)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	// if failed, return unsuccessful (false), else true
	return response.StatusCode < 400
}

// Fire request to /authorize endpoint with user payload
// Returns UserAccount object
func getUser(username string, pin int) UserAccount {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"username\": \"%s\", \"pin\": %d}", username, pin)
	payload := bytes.NewBuffer([]byte(body))

	response, err := client.Post(ROOT_URL+"/authorize", "application/json", payload)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	var user UserAccount
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	return user
}

// Process Function
// Hosts all supported operations [Create, Read, Update, Delete]
func process(id int, banks *[]BankAccount, buckets *[]Bucket, lineitems *[]LineItem) {

	entities := []string{"BANK", "BUCKET", "LINEITEM"}
	methods := []string{"CREATE", "VIEW", "UPDATE", "DELETE"}

	var method string
	for {
		fmt.Print("What would you like to do? ")
		fmt.Println(methods)
		fmt.Scan(&method)

		// Check if valid operation, if not - retry
		validMethod := false
		for i := 0; i < len(methods); i++ {
			if method == methods[i] {
				validMethod = true
			}
		}
		if !validMethod {
			fmt.Println("Invalid Method, Try Again!")
		} else {
			break
		}
	}

	var entity string
	for {
		fmt.Print("What record would you like to see? ")
		fmt.Println(entities)
		fmt.Scan(&entity)

		// Check if valid resource, if not - retry
		validEntity := false
		for i := 0; i < len(entities); i++ {
			if entity == entities[i] {
				validEntity = true
			}
		}
		if !validEntity {
			fmt.Println("Invalid Record Type, Try Again!")
		} else {
			break
		}
	}

	switch method {
	case "CREATE":
		switch entity {
		case "BANK": // Operation for creating bank account
			var name string
			fmt.Print("Creating Bank. \nName: ")
			fmt.Scan(&name)

			success := false
			success = createBank(name, id)
			if success {
				fmt.Println("Bank created!")
				fmt.Println("[Bank Id, Bank Name, Bank Owner Id]")
				fmt.Println("Your Banks: ")
				*banks = getBanks(id)
				fmt.Println(*banks)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}
		case "BUCKET": // Operation for creating bucket
			var name string
			fmt.Print("Creating Bucket. \nName: ")
			fmt.Scan(&name)

			success := false
			success = createBucket(name, id)
			if success {
				fmt.Println("Bucket created!")
				fmt.Println("[Bucket Id, Bucket Name, Bucket Owner Id]")
				fmt.Println("Your Buckets: ")
				*buckets = getBuckets(id)
				fmt.Println(*buckets)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}
		case "LINEITEM": // Operation for creating line item/expense entry
			var title, description string
			var amount float64
			var bucket, bank int
			fmt.Println("Creating Line Item Entry.")

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Title: ")
			if scanner.Scan() {
				title = scanner.Text()
			}

			fmt.Print("Description: ")
			if scanner.Scan() {
				description = scanner.Text()
			}

			fmt.Print("Amount: ")
			fmt.Scan(&amount)

			// User can map line item entry to bucket
			for {
				fmt.Print("Available Buckets: ")
				fmt.Println(*buckets)
				fmt.Print("Choose the Bucket Id: (0 for no bucket) ")
				fmt.Scan(&bucket)

				validBucket := false
				for _, buc := range *buckets {
					if bucket == buc.Id || bucket == 0 {
						validBucket = true
						break
					}
				}
				if !validBucket {
					fmt.Println("Invalid Bucket. Try again!")
					continue
				} else {
					break
				}
			}

			// User can map line item entry to bank
			for {
				fmt.Print("Available Banks: ")
				fmt.Println(*banks)
				fmt.Print("Choose the Bank Id: (0 for no bank) ")
				fmt.Scan(&bank)

				validBank := false
				for _, ban := range *banks {
					if bank == ban.Id || bank == 0 {
						validBank = true
						break
					}
				}
				if !validBank {
					fmt.Println("Invalid Bank. Try again!")
					continue
				} else {
					break
				}
			}

			success := false
			success = createLineItem(title, description, amount, bucket, bank, id)
			if success {
				fmt.Println("Bucket created!")
				fmt.Println("[Bucket Id, Bucket Name, Bucket Owner Id]")
				fmt.Println("Your Buckets: ")
				fmt.Println(getBuckets(id))
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}
		}
	case "VIEW":
		switch entity {
		case "BANK": // Operation for retrieving bank records
			fmt.Println("Your current banks: [Bank Id, Bank Name, Your ID]")
			*banks = getBanks(id)
			fmt.Println(*banks)
		case "BUCKET": // Operation for retrieving bucket records
			fmt.Println("Your current buckets: [Bucket Id, Bucket Name, Your ID]")
			*buckets = getBuckets(id)
			fmt.Println(*buckets)
		case "LINEITEM": // Operation for retrieving line item/expense entries
			fmt.Println("Your current items: [Line Item Id, Name, Description, Amount, Bucket, Bank, Your ID]")
			*lineitems = getLineItems(id)
			fmt.Println(*lineitems)
		}
	case "UPDATE":
		switch entity {
		case "BANK": // Operation for updating bank records
			fmt.Println("Your current banks: [Bank Id, Bank Name, Your ID]")
			*banks = getBanks(id)
			fmt.Println(*banks)

			var bankId int
			var bank BankAccount

			fmt.Print("Enter the Bank Id for update: ")
			fmt.Scan(&bankId)

			for _, val := range *banks {
				if val.Id == bankId {
					bank = val
					break
				} else {
					fmt.Println("Invalid ID. Try again.")
				}
			}

			var bankName string
			fmt.Printf("Your Bank Name is %s, change to: \n", bank.Name)

			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				bankName = scanner.Text()
			}

			fmt.Println(bankId)
			fmt.Println(bankName)

			success := false
			success = updateBank(bankId, bankName, id)

			if success {
				fmt.Println("Bank updated!")
				fmt.Println("[Bank Id, Bank Name, Bank Owner Id]")
				fmt.Println("Your Banks: ")
				*banks = getBanks(id)
				fmt.Println(*banks)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}

		case "BUCKET": // Operation for updating bucket records
			fmt.Println("Your current buckets: [Bucket Id, Bucket Name, Your ID]")
			*buckets = getBuckets(id)
			fmt.Println(*buckets)

			var bucketId int
			var bucket Bucket

			fmt.Print("Enter the Bucket Id for update: ")
			fmt.Scan(&bucketId)

			for _, val := range *buckets {
				if val.Id == bucketId {
					bucket = val
					break
				} else {
					fmt.Println("Invalid ID. Try again.")
				}
			}

			var bucketName string
			fmt.Printf("Your Bucket Name is %s, change to: \n", bucket.Name)

			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				bucketName = scanner.Text()
			}

			success := false
			success = updateBucket(bucketId, bucketName, id)

			if success {
				fmt.Println("Bucket updated!")
				fmt.Println("[Bucket Id, Bucket Name, Your ID]")
				fmt.Println("Your Buckets: ")
				*buckets = getBuckets(id)
				fmt.Println(*buckets)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}

		case "LINEITEM": // Updating Line Item will not be supported. Suggest to do DELETE then ADD
			fmt.Println("Not supported! Please perform delete then add operation instead.")
			fmt.Println("Returning to main menu...")
		}
	case "DELETE":
		switch entity {
		case "BANK": // Operation for deleting bank record
			fmt.Println("Your current banks: [Bank Id, Bank Name, Your ID]")
			*banks = getBanks(id)
			fmt.Println(*banks)

			var bank int
			fmt.Print("Enter the Bank Id for deletion: ")
			fmt.Scan(&bank)

			success := false
			success = deleteBank(id, bank)
			if success {
				fmt.Println("Bank deleted!")
				fmt.Println("[Bank Id, Bank Name, Bank Owner Id]")
				fmt.Println("Your Banks: ")
				*banks = getBanks(id)
				fmt.Println(*banks)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}

		case "BUCKET": // Operation for deleting bucket record
			fmt.Println("Your current buckets: [Bucket Id, Bucket Name, Your ID]")
			*buckets = getBuckets(id)
			fmt.Println(*buckets)

			var bucket int
			fmt.Print("Enter the Bucket Id for deletion: ")
			fmt.Scan(&bucket)

			success := false
			success = deleteBucket(id, bucket)
			if success {
				fmt.Println("Bucket deleted!")
				fmt.Println("[Bucket Id, Bucket Name, Bucket Owner Id]")
				fmt.Println("Your Buckets: ")
				*buckets = getBuckets(id)
				fmt.Println(*buckets)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}

		case "LINEITEM": // Operation for deleting line item/expense entry
			fmt.Println("Your current items: [Line Item Id, Name, Description, Amount, Bucket, Bank, Your ID]")
			*lineitems = getLineItems(id)
			fmt.Println(*lineitems)

			var lineitem int
			fmt.Print("Enter the Line Item Id for deletion: ")
			fmt.Scan(&lineitem)

			success := false
			success = deleteLineItem(id, lineitem)
			if success {
				fmt.Println("LineItem deleted!")
				fmt.Println("[LineItem Id, LineItem Name, LineItem Owner Id]")
				fmt.Println("Your LineItems: ")
				*lineitems = getLineItems(id)
				fmt.Println(*lineitems)
			} else {
				fmt.Println("Unexpected error occured. Try again!")
			}
		}
	}
}

// Get Bank Records from Server HTTP API
func getBanks(ownerid int) []BankAccount {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	response, err := client.Get(ROOT_URL + "/banks")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var banks, filtered []BankAccount

	if err := json.NewDecoder(response.Body).Decode(&banks); err != nil {
		log.Fatal(err)
	}

	for _, bank := range banks {
		if bank.Owner == ownerid {
			filtered = append(filtered, bank)
		}
	}

	return filtered
}

// Get Bucket Records from Server HTTP API
func getBuckets(ownerid int) []Bucket {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	response, err := client.Get(ROOT_URL + "/buckets")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var buckets, filtered []Bucket

	if err := json.NewDecoder(response.Body).Decode(&buckets); err != nil {
		log.Fatal(err)
	}

	for _, bucket := range buckets {
		if bucket.Owner == ownerid {
			filtered = append(filtered, bucket)
		}
	}

	return filtered
}

// Get Line Item/Expense Entries from Server HTTP API
func getLineItems(ownerid int) []LineItem {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	response, err := client.Get(ROOT_URL + "/lineitems")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var lineitems, filtered []LineItem

	if err := json.NewDecoder(response.Body).Decode(&lineitems); err != nil {
		log.Fatal(err)
	}

	for _, lineitem := range lineitems {
		if lineitem.Owner == ownerid {
			filtered = append(filtered, lineitem)
		}
	}

	return filtered
}

// Delete Bank Record via Server HTTP API
func deleteBank(ownerid int, id int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/bank/%d", ROOT_URL, id)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err2 := client.Do(request)
	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	return true
}

// Delete Bucket Record via Server HTTP API
func deleteBucket(ownerid int, id int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/bucket/%d", ROOT_URL, id)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err2 := client.Do(request)
	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	return true
}

// Delete Line Item/Expense Entry via Server HTTP API
func deleteLineItem(ownerid int, id int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/lineitem/%d", ROOT_URL, id)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err2 := client.Do(request)
	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	return true
}

// Create User Account via Server HTTP API
func createUser(username string, name string, pin int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"username\": \"%s\", \"name\": \"%s\", \"pin\": %d}", username, name, pin)
	payload := bytes.NewBuffer([]byte(body))

	response, err := client.Post(ROOT_URL+"/users", "application/json", payload)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response.StatusCode < 400
}

// Create Bank Account via Server HTTP API
func createBank(name string, ownerid int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"name\": \"%s\", \"ownerid\": %d}", name, ownerid)
	payload := bytes.NewBuffer([]byte(body))

	response, err := client.Post(ROOT_URL+"/banks", "application/json", payload)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response.StatusCode < 400
}

// Create Bucket via Server HTTP API
func createBucket(name string, ownerid int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"name\": \"%s\", \"ownerid\": %d}", name, ownerid)
	payload := bytes.NewBuffer([]byte(body))

	response, err := client.Post(ROOT_URL+"/buckets", "application/json", payload)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response.StatusCode < 400
}

// Create Line Item/Expense Entry via Server HTTP API
func createLineItem(title string, description string, amount float64, bucket int, bank int, ownerid int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}

	body := ""
	body += fmt.Sprintf("\"title\": \"%s\", \"description\": \"%s\", \"amount\": %f,", title, description, amount)
	if bucket == 0 {
		body += "\"bucket\": null,"
	} else {
		body += fmt.Sprintf("\"bucket\": %d,", bucket)
	}
	if bank == 0 {
		body += "\"bank\": null,"
	} else {
		body += fmt.Sprintf("\"bank\": %d,", bank)
	}
	body = "{" + body + fmt.Sprintf("\"ownerid\": %d}", ownerid)
	fmt.Println(body)
	payload := bytes.NewBuffer([]byte(body))

	response, err := client.Post(ROOT_URL+"/lineitems", "application/json", payload)
	fmt.Println(response)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response.StatusCode < 400
}

// Update Bank Account via Server HTTP API
func updateBank(id int, name string, ownerid int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"name\": \"%s\", \"ownerid\": %d}", name, ownerid)
	payload := bytes.NewBuffer([]byte(body))

	request, err := http.NewRequest("PUT", ROOT_URL+"/bank/"+fmt.Sprint(id), payload)
	if err != nil {
		log.Fatal(err)
	}

	response, err2 := client.Do(request)
	if err2 != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response.StatusCode < 400
}

// Update Bucket via Server HTTP API
func updateBucket(id int, name string, ownerid int) bool {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	body := fmt.Sprintf("{\"name\": \"%s\", \"ownerid\": %d}", name, ownerid)
	payload := bytes.NewBuffer([]byte(body))

	request, err := http.NewRequest("PUT", ROOT_URL+"/bucket/"+fmt.Sprint(id), payload)
	if err != nil {
		log.Fatal(err)
	}
	response, err2 := client.Do(request)
	if err2 != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	return response.StatusCode < 400
}

// Update Line Item/Expense Item via Server HTTP API
// func updateLineItem(id int, title string, description string, amount float64, bucket int, bank int, ownerid int) bool {
// 	client := http.Client{Timeout: time.Duration(1) * time.Second}
// 	body := fmt.Sprintf("{\"title\": \"%s\", \"description\": \"%s\", \"amount\": %f, \"bucket\": \"%d\", \"bank\": \"%d\", \"ownerid\": \"%d\"}", title, description, amount, bucket, bank, ownerid)
// 	payload := bytes.NewBuffer([]byte(body))

// 	request, err := http.NewRequest("PUT", ROOT_URL+"/lineitem/"+fmt.Sprint(id), payload)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	response, err2 := client.Do(request)
// 	if err2 != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()

// 	return response.StatusCode < 400
// }
