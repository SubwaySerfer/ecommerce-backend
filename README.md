# Go Backend Project

## Description
This project is a backend implementation for an e-commerce application. It provides APIs for managing furniture products, user favorites, and shopping cart functionalities. The backend is designed to be modular, with separate layers for handling HTTP requests, business logic, data access, and utility functions.

## Project Structure
- **cmd/main.go**: Entry point of the application, initializes the server and sets up routes.
- **internal/handlers**: Contains HTTP handler functions for processing requests.
- **internal/models**: Defines data structures used in the application.
- **internal/repositories**: Handles data access logic and CRUD operations.
- **internal/services**: Contains business logic and interacts with the repository layer.
- **internal/utils**: Utility functions for validation and other helper methods.
- **pkg/config**: Manages configuration settings for the application.

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## API Endpoints
- **GET /furniture**: Retrieve a list of furniture items.
- **POST /favorites**: Add a furniture item to the favorites list.
- **DELETE /favorites/{id}**: Remove a furniture item from the favorites list.
- **POST /cart**: Add items to the shopping cart.
- **GET /cart**: Retrieve the current shopping cart items.

## Features
- Modular architecture for easy maintenance and scalability.
- RESTful API design for seamless integration with frontend applications.
- Error handling and logging for better debugging and monitoring.

## Technologies Used
- Go (Golang)
- Gorilla Mux (for routing)
- [Database of choice] (e.g., PostgreSQL, MongoDB)

## License
This project is licensed under the MIT License. See the LICENSE file for details.