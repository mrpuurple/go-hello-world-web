# Blog Conversion Plan

## Overview

This document outlines the steps to convert the current Go Hello World Web application into a simple but functional blog system. The plan leverages the existing architecture while adding blog-specific functionality.

## Current Architecture Analysis

**Strengths to Build Upon:**
- ✅ Solid foundation with chi router and middleware pipeline
- ✅ Template caching system with partial template support
- ✅ Session management with SCS
- ✅ Clean separation of concerns (handlers, models, render)
- ✅ Modern UI with Tailwind CSS and responsive navigation
- ✅ CSRF protection and security middleware

**Areas Requiring Enhancement:**
- 🔄 No data persistence layer
- 🔄 Static content structure
- 🔄 Limited routing for dynamic content

## Phase 1: Data Layer & Models

### 1.1 Database Setup
- **Choose Database**: SQLite for simplicity (can upgrade to PostgreSQL later)
- **Add Dependencies**:
  ```bash
  go get github.com/jmoiron/sqlx
  go get github.com/mattn/go-sqlite3
  go get github.com/golang-migrate/migrate/v4
  ```

### 1.2 Database Schema
Create migrations for:
```sql
-- Posts table
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    content TEXT NOT NULL,
    excerpt TEXT,
    published BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tags table (optional for v1)
CREATE TABLE tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) UNIQUE NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL
);

-- Post-Tag junction table
CREATE TABLE post_tags (
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);
```

### 1.3 Data Models
Create `pkg/models/blog.go`:
```go
type Post struct {
    ID        int       `db:"id" json:"id"`
    Title     string    `db:"title" json:"title"`
    Slug      string    `db:"slug" json:"slug"`
    Content   string    `db:"content" json:"content"`
    Excerpt   string    `db:"excerpt" json:"excerpt"`
    Published bool      `db:"published" json:"published"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
    Tags      []Tag     `json:"tags,omitempty"`
}

type Tag struct {
    ID   int    `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
    Slug string `db:"slug" json:"slug"`
}
```

## Phase 2: Repository Pattern

### 2.1 Database Connection
Create `pkg/database/connection.go`:
- Database initialization
- Connection pooling
- Migration runner

### 2.2 Repository Layer
Create `pkg/repository/`:
- `post_repository.go` - CRUD operations for posts
- `tag_repository.go` - Tag management (for later phases)
- `interfaces.go` - Repository interfaces for testing

### 2.3 Repository Methods
```go
type PostRepository interface {
    GetAll(limit, offset int) ([]Post, error)
    GetBySlug(slug string) (*Post, error)
    GetPublished(limit, offset int) ([]Post, error)
    Create(post *Post) error
    Update(post *Post) error
    Delete(id int) error
    GenerateSlug(title string) (string, error)
}
```

## Phase 3: Enhanced Routing & Handlers

### 3.1 New Routes
Update `cmd/web/routes.go`:
```go
// Public blog routes
mux.Get("/", handlers.Repo.BlogHome)
mux.Get("/post/{slug}", handlers.Repo.PostDetail)
mux.Get("/posts", handlers.Repo.PostList)
mux.Get("/tag/{slug}", handlers.Repo.PostsByTag) // Future

// Admin routes (protected)
mux.Route("/admin", func(r chi.Router) {
    r.Use(middleware.BasicAuth("admin", map[string]string{
        "admin": "password", // TODO: Use env vars
    }))
    r.Get("/", handlers.Repo.AdminDashboard)
    r.Get("/posts", handlers.Repo.AdminPostList)
    r.Get("/posts/new", handlers.Repo.AdminPostNew)
    r.Post("/posts", handlers.Repo.AdminPostCreate)
    r.Get("/posts/{id}/edit", handlers.Repo.AdminPostEdit)
    r.Put("/posts/{id}", handlers.Repo.AdminPostUpdate)
    r.Delete("/posts/{id}", handlers.Repo.AdminPostDelete)
})
```

### 3.2 Blog Handlers
Create `pkg/handlers/blog_handlers.go`:
```go
func (m *Repository) BlogHome(w http.ResponseWriter, r *http.Request)
func (m *Repository) PostDetail(w http.ResponseWriter, r *http.Request)
func (m *Repository) PostList(w http.ResponseWriter, r *http.Request)
```

### 3.3 Admin Handlers
Create `pkg/handlers/admin_handlers.go`:
```go
func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request)
func (m *Repository) AdminPostList(w http.ResponseWriter, r *http.Request)
func (m *Repository) AdminPostNew(w http.ResponseWriter, r *http.Request)
func (m *Repository) AdminPostCreate(w http.ResponseWriter, r *http.Request)
// ... etc
```

## Phase 4: Template System Enhancement

### 4.1 New Templates
```
templates/
├── blog/
│   ├── home.page.tmpl          # Blog homepage with recent posts
│   ├── post-detail.page.tmpl   # Individual post view
│   ├── post-list.page.tmpl     # Archive/list view
│   └── post-card.partial.tmpl  # Reusable post card component
├── admin/
│   ├── dashboard.page.tmpl     # Admin overview
│   ├── post-list.page.tmpl     # Admin post management
│   ├── post-form.page.tmpl     # Create/edit post form
│   └── admin-nav.partial.tmpl  # Admin navigation
└── partials/
    ├── nav.partial.tmpl        # Updated main navigation
    └── pagination.partial.tmpl # Pagination component
```

### 4.2 Template Data Enhancement
Update `pkg/models/templatedata.go`:
```go
type TemplateData struct {
    StringMap map[string]string
    IntMap    map[string]int
    FloatMap  map[string]float32
    Data      map[string]interface{}
    CSRFToken string
    Flash     string
    Warning   string
    Error     string
    Posts     []Post      // For blog templates
    Post      *Post       // For single post view
    Tags      []Tag       // For tag-related views
    Pagination *Pagination // For paginated views
}

type Pagination struct {
    CurrentPage  int
    TotalPages   int
    TotalPosts   int
    HasPrevious  bool
    HasNext      bool
    PreviousPage int
    NextPage     int
}
```

## Phase 5: Content Management Features

### 5.1 Markdown Support
Add markdown parsing:
```bash
go get github.com/russross/blackfriday/v2
```

### 5.2 Slug Generation
Implement automatic URL-friendly slug generation from titles.

### 5.3 Excerpt Generation
Auto-generate post excerpts from content for listing pages.

### 5.4 Publish/Draft System
Toggle between draft and published states.

## Phase 6: Enhanced UI/UX

### 6.1 Blog-Specific Styling
- Post typography optimization
- Code syntax highlighting (optional)
- Responsive post cards
- Clean post detail layout

### 6.2 Navigation Updates
Update `templates/nav.partial.tmpl`:
```html
<nav class="bg-white shadow-lg border-b border-gray-200">
    <div class="max-w-6xl mx-auto px-4">
        <div class="flex justify-between items-center py-4">
            <div class="flex items-center">
                <h1 class="text-2xl font-bold text-gray-900">My Blog</h1>
            </div>
            <div class="flex space-x-8">
                <a href="/" class="nav-link">Home</a>
                <a href="/posts" class="nav-link">All Posts</a>
                <a href="/about" class="nav-link">About</a>
                <a href="/admin" class="nav-link admin-link">Admin</a>
            </div>
        </div>
    </div>
</nav>
```

### 6.3 Admin Interface
Create a clean, functional admin interface with:
- Post creation/editing forms
- Post management table
- Dashboard with basic statistics

## Phase 7: Configuration & Environment

### 7.1 Environment Configuration
Create `.env` support:
```go
// pkg/config/config.go
type AppConfig struct {
    InProduction   bool
    UseCache       bool
    TemplateCache  map[string]*template.Template
    Session        *scs.SessionManager
    Database       *sqlx.DB

    // Blog-specific config
    BlogTitle      string
    BlogSubtitle   string
    PostsPerPage   int
    AdminUsername  string
    AdminPassword  string
}
```

### 7.2 Database Configuration
Support for different database URLs and connection pooling settings.

## Phase 8: Optional Enhancements (Future Phases)

### 8.1 Search Functionality
- Simple full-text search across posts
- Search results page

### 8.2 RSS Feed
- Generate RSS/Atom feeds for blog posts

### 8.3 Comments System
- Simple comment system (if desired)
- Comment moderation

### 8.4 Image Upload
- File upload handling for post images
- Image optimization and serving

### 8.5 Caching Layer
- Redis integration for performance
- Template fragment caching

## Implementation Priority

### Phase 1 (Essential - Week 1)
1. ✅ Database setup and models
2. ✅ Basic repository layer
3. ✅ Core blog handlers (home, post detail)
4. ✅ Essential templates

### Phase 2 (Core Features - Week 2)
1. ✅ Admin interface
2. ✅ Post creation/editing
3. ✅ Markdown support
4. ✅ Navigation updates

### Phase 3 (Polish - Week 3)
1. ✅ Enhanced styling
2. ✅ Pagination
3. ✅ Draft/publish system
4. ✅ Configuration management

### Phase 4 (Optional - Future)
1. 🔮 Search, RSS, Comments
2. 🔮 Image uploads
3. 🔮 Performance optimizations

## Migration Strategy

1. **Incremental Development**: Build alongside existing functionality
2. **Feature Flags**: Use configuration to enable/disable blog features
3. **Data Migration**: Scripts to populate initial blog content
4. **Testing**: Maintain existing functionality while adding new features

## File Structure After Conversion

```
├── cmd/web/
│   ├── main.go
│   ├── routes.go              # Enhanced with blog routes
│   └── middleware.go
├── pkg/
│   ├── config/
│   │   └── config.go          # Enhanced configuration
│   ├── database/
│   │   ├── connection.go      # New: DB connection
│   │   └── migrations/        # New: DB migrations
│   ├── handlers/
│   │   ├── handlers.go        # Existing handlers
│   │   ├── blog_handlers.go   # New: Blog handlers
│   │   └── admin_handlers.go  # New: Admin handlers
│   ├── models/
│   │   ├── templatedata.go    # Enhanced template data
│   │   └── blog.go           # New: Blog models
│   ├── repository/
│   │   ├── interfaces.go      # New: Repository interfaces
│   │   └── post_repository.go # New: Post repository
│   └── render/
│       └── render.go          # Enhanced rendering
├── templates/
│   ├── blog/                  # New: Blog templates
│   ├── admin/                 # New: Admin templates
│   ├── partials/              # Enhanced partials
│   ├── about.page.tmpl        # Existing
│   └── base.layout.tmpl       # Existing
├── static/                    # New: Static assets
├── migrations/                # New: Database migrations
└── .env.example              # New: Environment template
```

This plan provides a structured approach to converting your web application into a fully functional blog while preserving the existing clean architecture and modern UI foundation.