# Shorten Service

Microservice that generates short URL from long URL.

## Explanation

We take a long URL like "https://kbtu.golang.kz/final/exam" and return "http://localhost:8080/expl1".

## Teck Stack

Go, Gin for REST API

Redis for fast key-value storage

Docker for containerization

## Get Started

1. git clone https://github.com/Edaaail/goproj

2. cd goproj

3. go mod tidy 

4. docker-compose up --build

## Example
Run: 

 curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d "{\"long_url\": \"https://example.com/very/long/url\"}"

Response:

{"short_url":"http://localhost:8080/jfomah"}