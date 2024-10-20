#! /bin/bash

migrate -path database/migrations -database "postgres://db-user:db-pass@localhost:5432/objsto?sslmode=disable" up

