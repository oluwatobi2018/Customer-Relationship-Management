#!/bin/bash

DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=1 DB_NAME=gocrm DB_SSLMODE=disable \
go run .
