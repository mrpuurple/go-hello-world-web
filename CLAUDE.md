# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Running the Application

```bash
cd cmd/web && go run main.go
```

The server starts on port 8888 (configured in main.go:16).

### Development Tasks

```bash
# Install/update dependencies
go mod tidy

# Build the application
go build -o app ./cmd/web

# Run from project root (requires templates directory)
go run ./cmd/web
```

### Testing

No test files currently exist in this codebase. Standard Go testing would use:

```bash
go test ./...
```

## Architecture Overview

This is a modular Go web application using the chi router with session management and template rendering.

### Project Structure

- `cmd/web/` - Main application entry point and HTTP configuration
  - `main.go` - Application bootstrap, server setup, and dependency injection
  - `routes.go` - HTTP routes using chi router with middleware
  - `middleware.go` - Custom middleware (NoSurf, SessionLoad)
- `pkg/` - Reusable packages
  - `config/` - Application configuration struct (AppConfig)
  - `handlers/` - HTTP request handlers using repository pattern
  - `models/` - Data structures (TemplateData)
  - `render/` - Template rendering engine with caching
- `templates/` - HTML templates (*.page.tmpl, *.layout.tmpl, *.partial.tmpl)
- `static/` - Static assets served by the web server
  - `css/` - Stylesheet files organized in subdirectories
  - `images/` - Image assets including favicon
  - `fonts/` - Font files
  - `js/` - JavaScript files

### Key Dependencies

- `github.com/go-chi/chi` - HTTP router and middleware
- `github.com/alexedwards/scs/v2` - Session management
- `github.com/justinas/nosurf` - CSRF protection

### Architecture Patterns

- **Repository Pattern**: Handlers use Repository struct for dependency injection
- **Template Caching**: Templates are parsed once and cached in AppConfig
- **Session Management**: Uses SCS for secure session handling
- **Middleware Pipeline**: Recoverer, NoSurf (CSRF), SessionLoad

### Configuration

- Server port: Hardcoded to `:8888` in main.go
- Production mode: Controlled by `app.InProduction` boolean
- Template caching: Enabled by default (`app.UseCache = true`)
- Session lifetime: 24 hours with persistent cookies

### Template System

Templates use Go's html/template with:

- Page templates: `*.page.tmpl` - Main content templates (home, about)
- Layout templates: `*.layout.tmpl` - Base layout structure
- Partial templates: `*.partial.tmpl` - Reusable components (nav, footer)
- Template data passed via `models.TemplateData` struct

### Static Assets

Static files are served from the `static/` directory and include:

- CSS files with modular organization (base, components, pages, utilities)
- Custom favicon (SVG format)
- Font assets for typography
- JavaScript files for client-side functionality
