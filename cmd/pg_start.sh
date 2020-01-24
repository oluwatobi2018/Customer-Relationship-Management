#!/bin/bash

# docker run --rm -e POSTGRES_DB="gocrm" -e POSTGRES_USER="postgres" -e POSTGRES_PASSWORD="1" --network=k8s --name=pg -d -p 5432:5432 postgres:12
docker run --rm -e POSTGRES_DB="gocrm" -e POSTGRES_USER="postgres" -e POSTGRES_PASSWORD="1" --name=pg -d -p 5432:5432 postgres:12