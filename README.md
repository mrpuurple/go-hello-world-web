# Go Hello World Web

A simple web application written in Go, demonstrating basic routing, middleware, configuration, and template rendering.

## Features

- Basic HTTP server using Go's net/http
- Custom middleware for request handling
- Configurable settings via `pkg/config/config.go`
- Modular handlers in `pkg/handlers/handlers.go`
- Template rendering with HTML templates in `templates/`
- Organized project structure for scalability

## Project Structure

```go
go.mod
go.sum
cmd/
  web/
    main.go
    middleware.go
    routes.go
pkg/
  config/
    config.go
  handlers/
    handlers.go
  models/
    templatedata.go
  render/
    render.go
templates/
  about.page.tmpl
  base.layout.tmpl
  home.page.tmpl
```

## Getting Started

### Prerequisites

- Go 1.18 or newer

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/mrpuurple/go-hello-world-web.git
   cd go-hello-world-web
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

### Running the Application

Navigate to the `cmd/web` directory and run:

```sh
go run main.go
```

The server will start and listen on the configured port (default: 8080).

### Accessing the App

Open your browser and go to:

```sh
http://localhost:8080/
```

## Customization

- Update configuration in `pkg/config/config.go`
- Add new routes in `cmd/web/routes.go`
- Create new templates in the `templates/` directory

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
