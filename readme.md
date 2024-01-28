# Golang Todo Application Backend

Welcome to the Golang Todo Application Backend! This backend serves as the backend for a todo application similar to Microsoft Todo. It is designed to provide a seamless experience for users, even those unfamiliar with Golang. The application features user authentication through `JWT bearer tokens`, `PostgreSQL` database integration, migration management, and a RESTful API documented with Swagger (OpenAPI).

## Getting Started

Follow these steps to set up and run the Golang Todo Application Backend on your local machine.

### Prerequisites

- [Golang](https://golang.org/doc/install) installed on your machine (last update on go 1.21.4)
- PostgreSQL database server
- [Make](https://www.gnu.org/software/make/) tool installed

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/moxicom/todo-back.git
   cd todo-back
   ```

2. Create a `.env` file and a `configs/config.yaml` file with your database connection details:
    
    .env:
   ```env
    DB_PASSWORD=YOUR_PASSWORD
    PASSWORD_SALT=mysalt
   ```

    config.yaml
   ```yaml
   # Server port
    server_port: "8080"

    # db
    host: "localhost"
    user: "postgres"
    dbname: "gorm_test"
    port: "5432"
    sslmode: "disable"
   ```

3. Run the migration to set up the database:

   ```bash
   make migrate
   ```

### Usage

1. Start the Golang Todo Application Backend:

   ```bash
   make run
   ```

   The application will run on `http://localhost:8080` by default.

2. Access the Swagger documentation at `your_server_link/swagger/index.html` to explore and interact with the REST API.

## API Documentation

The API is documented using Swagger (OpenAPI). Visit [Swagger Documentation](http://localhost:8080/swagger/index.html) for detailed information on available endpoints, request/response formats, and authentication requirements.

## Contributing

Feel free to contribute to this project by submitting bug reports, feature requests, or pull requests. Your contributions are highly appreciated!

## Acknowledgments

- Special thanks to the Golang community and the developers of libraries and tools used in this project.

Happy coding! 🚀