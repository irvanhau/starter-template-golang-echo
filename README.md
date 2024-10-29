# Starter Template Golang Echo

This is Starter Template if you want to create REST API project using Golang framework Echo, Already used Dependency Injection

## Created By [Irvan Hauwerich](https://www.linkedin.com/in/irvan-hauwerich-b953a822b/)

## Features

- Login
- Register
- Profile using JWT

## 3rd Party App

1. For Database you can use PostgreSQL or MySQL

## Another Link

[API Documentation POSTMAN](https://documenter.getpostman.com/view/33387055/2sAY4uE4VB)

## REST API Design

Request POST ```{{BASE_URL}}/api/v1/login```

Response

```json
{
  "data": {
    "username": "admin",
    "token": {
      "access_token": "this is access token",
      "refresh_token": "this is refresh token"
    }
  },
  "message": "Success Login"
}
```

## Project Structure

Clean Architecture

## How To Use

- ```git clone https://github.com/irvanhau/starter-template-golang.git``` or you can fork
- ```cp .env.example .env```
- create the database first
- change .env with your config
- ```go mod tidy```
- ```go run main.go wire_gen.go```

## How To Use Variable Secret Github Action
- DOCKERHUB_USERNAME = irvan (Docker Hub Username)
- DOCKERHUB_TOKEN = (Generate Docker Hub Token)
- HOST = 127.0.0.1 (SSH Host)
- USERNAME root (SSH Username)
- KEY = (Using SSH Private Key)
- PORT = 22 (SSH Port)

### Notes

- If you create the module, don't forget to add on ```injector.go``` and run command ```wire injector.go```
- If any question you can ask me via LinkedIn Thank you
- Don't Forget for Follow and star the project