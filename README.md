# Personal Blog

A simple personal blog web application built with **Go**. The project allows users to view published articles and provides an admin interface for creating and deleting articles.

The project was built to practice backend development concepts in Go, including HTTP servers, routing, handlers, HTML templates, JSON file storage, and HTTP methods.

## Features

* View all blog articles
* View individual articles
* Create new articles through the admin interface
* Delete existing articles
* Store articles as JSON files
* Render dynamic HTML using Go templates
* Handle HTTP requests using Go's standard library
* Separate article data from presentation using templates

## Technologies Used

* **Go**
* **HTML**
* **JSON**
* Go `net/http`
* Go `html/template`
* Go `encoding/json`
* Go `os` package

## Project Structure

```text
Personal_Blog/
│
├── main.go
│
├── articles/
│   ├── article1.json
│   ├── article2.json
│   └── ...
│
├── templates/
│   ├── home.html
│   ├── article.html
│   └── admin.html
│
├── static/
│   └── css/
│       └── style.css
│
└── README.md
```

> The exact structure may change as the project develops.

## How It Works

The application uses Go's HTTP server to receive requests and route them to the appropriate handlers.

### Home Page

The home page displays the available blog articles.

When a user visits the homepage, the server reads the stored article data and passes it to the HTML template.

```text
Browser
   ↓
GET /
   ↓
homeHandler
   ↓
Read article files
   ↓
Parse article data
   ↓
Render home.html
   ↓
Browser
```

### Individual Articles

Each article has an ID that is used to identify it.

For example:

```text
/article?id=3
```

The server uses the ID to locate the corresponding article and displays it using the article template.

```text
Browser
   ↓
GET /article?id=3
   ↓
ArticleHandler
   ↓
Find article with ID 3
   ↓
Read article data
   ↓
Render article.html
```

### Creating Articles

The admin section allows a new article to be created.

The server receives the submitted information through an HTTP request, creates the article data, assigns an ID, converts the data to JSON, and saves it to the articles directory.

```text
Admin form
    ↓
POST /create
    ↓
createHandler
    ↓
Create article
    ↓
Assign ID
    ↓
Marshal to JSON
    ↓
Save article
```

### Deleting Articles

Articles can be deleted from the admin interface.

The existing article ID is used to identify which JSON file should be removed.

```text
Delete request
     ↓
Existing article ID
     ↓
deleteHandler
     ↓
Locate article file
     ↓
Delete file
```

## Article Data

Articles are stored as JSON files.

A typical article may look like:

```json
{
    "id": 1,
    "title": "My First Article",
    "content": "This is the content of my first article."
}
```

The JSON data is converted into Go structs using `json.Unmarshal()` when it is read.

When a new article is saved, the Go struct is converted back into JSON using `json.Marshal()` or `json.MarshalIndent()`.

## HTTP Methods

The application uses different HTTP methods depending on the operation.

| Method   | Purpose                         |
| -------- | ------------------------------- |
| `GET`    | Retrieve pages and articles     |
| `POST`   | Submit/create new articles      |
| `DELETE` | Delete articles where supported |

For example:

```text
GET  /              → View articles
GET  /article?id=1  → View article
GET  /admin         → Open admin page
POST /create        → Create article
```

## Running the Project

### 1. Clone the repository

```bash
git clone <your-repository-url>
```

### 2. Navigate into the project

```bash
cd Personal_Blog
```

### 3. Run the application

```bash
go run .
```

The server will start locally.

Open the address shown by the server in your browser, for example:

```text
http://localhost:8080
```

## Requirements

Before running the project, make sure you have:

* Go installed
* Git installed
* A web browser

You can check your Go installation with:

```bash
go version
```

## Design Decisions

### File-Based Storage

The project currently uses JSON files instead of a database.

This keeps the application simple and makes it easier to understand how data persistence works at the filesystem level.

Each article can be stored separately, making it possible to locate an article using its ID.

### Go Standard Library

The project primarily uses Go's standard library instead of a web framework.

This provides practice with fundamental backend concepts such as:

* HTTP servers
* Routing
* Handlers
* Request methods
* Templates
* File I/O
* JSON serialization
* Error handling

## Future Improvements

Possible improvements include:

* Add a database such as SQLite or PostgreSQL
* Add authentication for the admin section
* Add article editing
* Improve error handling
* Add input validation
* Add better routing
* Add pagination
* Add search functionality
* Add categories and tags
* Improve responsive design
* Add automated tests
* Deploy the application
* Use UUIDs for article IDs
* Add middleware for logging and authentication

## What I Learned

This project helped me practice building a backend application from the ground up using Go.

Key concepts practiced include:

* Creating an HTTP server
* Working with HTTP requests and responses
* Creating and using handlers
* Working with GET and POST requests
* Serving HTML templates
* Reading and writing files
* Working with JSON
* Structs and slices
* Handling article IDs
* Organizing a backend project
* Connecting backend logic to a frontend
* Managing a project with Git and GitHub

## Status

**In development.**

The core blog functionality is working, and the project will continue to be improved with additional backend features, better security, testing, and eventually database-backed storage.
