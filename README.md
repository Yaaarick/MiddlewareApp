# MiddlewareApp
A Go-based web application utilizing the Echo framework. The project features middleware for role-based request handling, a service layer for utility functions, and modular endpoint management. Includes a /status endpoint and runs on port 8080.

This program is a basic web application built with Go and the Echo framework. It includes the following features:
Middleware: A RoleCheck middleware that processes user roles from request headers.
Service Layer: A Service that provides utility functions, such as calculating days left until a specific date.
Endpoint: A /status endpoint to handle HTTP requests.
Application Structure: Modular design with separate packages for middleware, service, and endpoint logic.
Server: Runs on port 8080 and uses Echo for routing and request handling.
