# Go Training Project
## Personal Finance Management Tool

This tool allows users to create user accounts, register bank accounts, record their expenses or income and create buckets for grouping expense entries.

This application is written on Golang, and runs on Docker.

<br>

## Required Tools
1. Docker
1. Golang Binary


<br>

## Server
The server is written on Golang, and uses PostgreSQL Database.
There are 4 main entities:

1. UserAccount
1. BankAccount
1. Bucket
1. LineItem

### User Account
This entity hosts all user information: name, unique username (identifier) and PIN.

### Bank Account
This entity hosts the information of the bank. This bank record is tied to a user account.

### Bucket
This entity refers to the name of a group of expenses/income. This group is also associated to a user account.

### Line Item
This entity refers to an expense or income entry. This object is associated to a user, and can be linked to a Bank or a Bucket.

<br>

## Running the Server
To run the server, simply go to the directory where the docker-compose.yml file is located. Run the following command in the same directory:

```
docker compose up
```

The program will run on port 9000.

<br>

## Running the client
The client application is also written in Golang. Simply go to the client directory and run the following command.

```
go run main.go
```

You will be prompted to create an account. To create an account, simply respond 'Y' and supply your username and PIN.

The user can perform the following operations:

* Create Bank
* View Bank
* Update Bank
* Delete Bank
* Create Bucket
* View Bucket
* Update Bucket
* Delete Bucket
* Create Line Item
* View Line Item
* Delete Line Item

Depending on the selected operation, the user may need to supply information needed per entity.

