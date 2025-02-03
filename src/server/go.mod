module example.com/server

go 1.23.0

replace example.com/backend => ../backend

require example.com/backend v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
)
