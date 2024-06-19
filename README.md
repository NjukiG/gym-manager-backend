# GYM MANAGER BACKEND

This is a backend API for managing gym operations, including users, trainers, classes, and memberships.

It uses a React frontend for the UI.

To be able to test out the different routes

## Requirements

- To be able to run this project on your machine, you need to clone the API and run it.

Make sure you have the following installed on your system:

- Go (version 1.16 or later)

## Installation

1. CLone the project

```sh
  git clone https://github.com/your-username/your-repo.git
  cd your-repo
```

2. Install dependencies.

```sh
    go mod tidy
```

3. To start the development server, run:

```sh
   go run main.go
```

The API will be available at http://localhost:3000.

## API Endpoints

### User Routes

POST /public/register - Register a new user
POST /public/login - Login a user
GET /protected/validate - check status if logged in
POST /protected/logout - logout the user

### Trainer Routes

POST /protected/trainers - Create a new trainer
GET /protected/trainers - Get all trainers
GET /protected/trainers/:id - Get a trainer by ID
POST /protected/trainers/assign - Assign a trainer to a class

### Class Routes

POST /protected/classes - Create a new class
GET /protected/classes - Get all classes
GET /protected/classes/:id - Get a class by ID
POST /classes/newMember - Enroll a member in a class

### Member Routes

POST /protected/memberships - Create a new membership
GET /protected/memberships - Get all memberships
GET /protected/memberships/:id - Get a membership by ID
PUT /protected/memberships/:id - Edit a membership by ID
DELETE /protected/memberships/:id - Delete a membership by ID


THis project uses react.js for the frontend and a Golang API for backend.
It uses tailwind for styling and components from HyperUI.
Public exercises and youtube videos API from Rapid API.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
