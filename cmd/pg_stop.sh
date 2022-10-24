#!/bin/bash

docker ps --format "{{.Image}};{{.ID}}" | grep "postgres:12"| awk -F \; '{print $2}' | xargs docker stop
