# My Porky Blog

A simple Go web application demonstrating modern web development patterns and best practices. This project serves as a learning example for building structured, maintainable web applications in Go, featuring a playful pork-themed blog design.

## Overview

This application provides a basic website with two pages (home and about) that showcases session management, template rendering, and clean architectural patterns. The home page captures visitor IP addresses and displays them on the about page, demonstrating session state management across requests.

## Features

- **Clean Architecture**: Modular design using repository pattern with clear separation of concerns
- **Session Management**: Secure cookie-based sessions with IP tracking demonstration
- **Template System**: HTML template rendering with intelligent caching for performance
- **Custom CSS Framework**: Organized, maintainable CSS architecture with design tokens
- **Static Asset Management**: Optimized serving of CSS, JavaScript, images, and fonts
- **Security**: Built-in CSRF protection and secure session handling
- **Middleware Pipeline**: Recoverer, CSRF protection, and session loading middleware

## Tech Stack

- **Go 1.25** - Core language
- **[Chi Router](https://github.com/go-chi/chi)** - Lightweight HTTP router with middleware support
- **[SCS](https://github.com/alexedwards/scs)** - Secure session management
- **[NoSurf](https://github.com/justinas/nosurf)** - CSRF protection middleware

## Project Structure

```
├── cmd/web/              # Application entry point
│   ├── main.go          # Server setup and dependency injection
│   ├── middleware.go    # Custom middleware (CSRF, sessions)
│   └── routes.go        # HTTP routes and static file serving
├── pkg/                 # Reusable packages
│   ├── config/          # Application configuration
│   ├── handlers/        # HTTP request handlers (repository pattern)
│   ├── models/          # Data structures and models
│   └── render/          # Template rendering with caching
├── static/              # Static assets
│   ├── css/             # Organized CSS framework
│   │   ├── base/        # Reset, variables, typography
│   │   ├── components/  # Reusable component styles
│   │   ├── pages/       # Page-specific styles
│   │   ├── utilities/   # Layout and utility classes
│   │   └── main.css     # Main stylesheet with imports
│   ├── js/              # JavaScript files
│   ├── images/          # Image assets
│   └── fonts/           # Custom fonts
├── templates/           # HTML templates
│   ├── *.page.tmpl     # Page templates
│   ├── *.layout.tmpl   # Layout templates
│   └── *.partial.tmpl  # Reusable template components
└── go.mod              # Go module definition
```

## Getting Started

### Prerequisites

- Go 1.25 or newer

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mrpuurple/go-hello-world-web.git
   cd go-hello-world-web
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

From the project root directory:

```bash

# Option 1: Run from cmd/web directory
cd cmd/web && go run main.go

# Option 2: Run from project root
go run ./cmd/web

```

The server will start on port 8888.

### Accessing the Application

Open your browser and navigate to:

- **Home Page**: <http://localhost:8888/>
- **About Page**: <http://localhost:8888/about>

The application will capture your IP address on the home page and display it on the about page, demonstrating session state management.

## Development

### Building the Application

```bash

# Build executable
go build -o app ./cmd/web

# Run the built executable
./app

```

### Adding New Features

1. **New Routes**: Add routes in `cmd/web/routes.go`
2. **Handlers**: Implement handlers in `pkg/handlers/handlers.go`
3. **Templates**: Create templates in `templates/` directory
4. **Styling**: Add CSS in appropriate `static/css/` subdirectory
5. **Configuration**: Update `pkg/config/config.go` for new settings

### CSS Architecture

The application uses a custom CSS framework organized into:

- **Base**: Foundational styles (reset, variables, typography)
- **Components**: Reusable UI components (navigation, cards, buttons)
- **Pages**: Page-specific styling (home, about)
- **Utilities**: Layout helpers and utility classes

#### Customizing Styles

1. **Design Tokens**: Modify CSS variables in `static/css/base/variables.css`
2. **Components**: Update component styles in `static/css/components/`
3. **New Pages**: Add page-specific CSS in `static/css/pages/`
4. **Layout**: Customize layout utilities in `static/css/utilities/layout.css`

#### CSS Variables Available

```css
/* Colors */
--color-primary: #2563eb;
--color-gray-100: #f3f4f6;

/* Typography */
--font-size-base: 1rem;
--font-weight-bold: 700;

/* Spacing */
--spacing-4: 1rem;
--spacing-8: 2rem;

/* Shadows & Effects */
--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
--transition-fast: 150ms ease-in-out;
```

### Architecture Patterns

- **Repository Pattern**: Handlers use dependency injection via Repository struct
- **Template Caching**: Templates are parsed once and cached for performance
- **Middleware Pipeline**: Clean separation of cross-cutting concerns
- **Session Management**: Secure, HTTP-only cookies with configurable lifetime

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
