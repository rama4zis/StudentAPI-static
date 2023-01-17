This is simple how you make API with golang

I'm using gin-gonic framework for this API

# How to run this API

1. Clone this repository
2. Run this command `go run main.go`
3. You can use this API with this endpoint `http://localhost:8080`
4. Tools for testing API is Postman or you can use curl command in your terminal

# API Usage

## Get all users

```bash
curl -X GET http://localhost:8080/students
```

## Get user by id

```bash
curl -X GET http://localhost:8080/students/{uid}
```

## Create user

```bash
curl --location --request POST 'http://localhost:8080/students' \
--header 'Content-Type: application/json' \
--data-raw  '{
    "uid": "222",
    "Name": "Putri",
    "Age": 22,
    "Gender": "famele",
    "Faculty": "Architecture",
    "Semester": 8,
    "Graduated": false
}'
```

## Update user

```bash
curl --location --request PUT 'http://localhost:8080/students/edit/{uid}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uid": "222",
    "Name": "Putri",
    "Age": 23,
    "Gender": "famele",
    "Faculty": "Architecture",
    "Semester": 10,
    "Graduated": true
}'
```

## Delete user

```bash
curl -X DELETE http://localhost:8080/students/delete/{uid}
```