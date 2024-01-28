# WebSocket-Chat

WebSocket-Chat is a WebSocket instant messaging program written in Go. The main purpose of this project is to learn Go language networking programming using Gin framework and MongoDB database.

## Features

- Real-time Chat: Enables instant communication using the WebSocket protocol, allowing users to send and receive messages in real-time.
- User Authentication: Utilizes JSON Web Token (JWT) for user identity verification and authorization.
- Message Storage: Persists chat messages using MongoDB for data persistence.

## Technology Stack

- **Go Language:** Developed the backend server using the Go programming language.
- **Gin Framework:** Built the web service on top of Gin to simplify routing and middleware handling.
- **WebSocket:** Implemented WebSocket functionality using the `github.com/gorilla/websocket` package.
- **MongoDB:** Used MongoDB as a database to store chat messages and user information.
