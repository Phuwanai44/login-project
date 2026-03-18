# Login System (Vue + Go + MongoDB)

A full-stack authentication system built with **Vue 3, TypeScript, Go (Gin), and MongoDB**.

This project demonstrates a complete login system with authentication, authorization, and user management.

---

## Features

* User Register
* User Login
* JWT Authentication
* Protected Routes
* Admin Middleware
* Users CRUD
* Dashboard
* Profile Page

---

## Tech Stack

### Frontend

* Vue 3
* TypeScript
* Bootstrap
* Axios
* Vue Router

### Backend

* Go
* Gin Framework
* MongoDB
* JWT Authentication

---

## Project Structure

```
login-project
│
├── backend
│   ├── cmd
│   │   └── main.go
│   ├── config
│   ├── controllers
│   ├── middleware
│   ├── models
│   └── routes
│
└── frontend
```

---

## Installation

### Backend

```
cd backend
go mod tidy
go run cmd/main.go
```

Server will run at

```
http://localhost:3000
```

---

### Frontend

```
cd frontend
npm install
npm run dev
```

App will run at

```
http://localhost:5173
```

---

## API Endpoints

### Authentication

POST /login
POST /register

### Users

GET /users
GET /users/:id
PUT /users/:id
DELETE /users/:id

---

## Screenshots

Login Page
Dashboard
Users Management

---

## Author

Phuwanai
