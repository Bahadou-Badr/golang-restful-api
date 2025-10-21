A JSON API for retrieving and managing information about movies. You can
think of the core functionality as being a bit like the [Open Movie Database API](http://www.omdbapi.com/).

### Here’s what I’ve accomplished so far:
- ✅ Structure and Organization: Implemented an idiomatic and flexible project structure for hashtag#RESTful APIs.
- ✅ JSON Responses: Designed clear and customizable hashtag#JSON responses while avoiding common pitfalls.
- ✅ JSON Requests: Parsed incoming requests, validated content, and managed bad requests gracefully.
- ✅ SQL Migrations: Used SQL migrations to manage database schema changes effectively.
- ✅ Database Operations: Setting timeouts on queries and optimizing Go’s sql.DB connection pool behavior.
- ✅ Beyond CRUD: Implementing advanced resource actions like fetching related data and partial updates with optimistic concurrency control.
- ✅ Filtering, Sorting, and Pagination: features for API endpoints.

### What’s coming next?
- 📌 Full-Text Search: Leveraging hashtag#PostgreSQL for natural-language data searches.
- 📌 Rate-Limiting: Enabling per-client request limits for better scalability.
- 📌 Background Tasks: Safely managing tasks in background goroutines.
- 📌 Graceful Shutdown: Supporting safe shutdowns, including waiting for background tasks to complete.
- 📌 Email Workflows: Building workflows for user activation, password resets, and sending emails.
- 📌 Authentication & Authorization: Adding secure user authentication (stateful tokens and #JWTs) and permission-based endpoint checks.
- 📌 CORS Management: Controlling cross-origin requests from front-end applications.
- 📌 Metrics: Exposing metrics for monitoring API health and performance.
- 📌 Deployment: Automating deployments on Linux servers and running the API as a background service.
-----------------------------------------------------------------------------------------------------------------------------------------
Ultimately, our Redlight API will support the following endpoints and actions:

| Method | URL Pattern                  | Action                                                |
|--------|------------------------------|-------------------------------------------------------|
| GET    | /v1/healthcheck              | Show application health and version information      |
| GET    | /v1/movies                   | Show the details of all movies                         |
| POST   | /v1/movies                   | Create a new movie                                    |
| GET    | /v1/movies/:id               | Show the details of a specific movie                   |
| PATCH  | /v1/movies/:id               | Update the details of a specific movie                 |
| DELETE | /v1/movies/:id               | Delete a specific movie                                |
| POST   | /v1/users                    | Register a new user                                   |
| PUT    | /v1/users/activated          | Activate a specific user                              |
| PUT    | /v1/users/password           | Update the password for a specific user                |
| POST   | /v1/tokens/authentication    | Generate a new authentication token                   |
| POST   | /v1/tokens/password-reset    | Generate a new password-reset token                   |
| GET    | /debug/vars                  | Display application metrics                           |

To give you an idea of what the API will look like from a client’s point of view, the `GET /v1/movies/:id` endpoint will return a response similar this
```
$ curl -H "Authorization: Bearer RIDBIAE3AMMK57T6IAEBUGA7ZQ" localhost:4000/v1/movies/1
{
    "movie": {
        "id": 1,
        "title": "Moana",
        "year": 2016,
        "runtime": "107 mins",
        "genres": [
            "animation",
            "adventure"
        ],
        "version": 1
    }
}
```

Behind the scenes, we’ll use PostgreSQL as the database for persistently storing all the
data


### The purpose that files and folders serve in our project : 

A high-level **_skeleton structure_** for the project:
```
.
├── bin
├── cmd
│  └── api
├── internal
├── migrations
├── remote
├── go.mod
└── Makefile
```
- The `bin` directory will contain our compiled application binaries, ready for deployment
to a production server.
- The `cmd/api` directory will contain the application-specific code for our Redlight API
application. This will include the code for running the server, reading and writing HTTP
requests, and managing authentication.
- The `internal` directory will contain various ancillary packages used by our API. It will
contain the code for interacting with our database, doing data validation, sending emails
and so on. Basically, any code that isn’t application-specific and can potentially be
reused will live in here. Our Go code under `cmd/api` will import the packages in the
`internal` directory (but never the other way around).
- The `migrations` directory will contain the SQL migration files for our database.
- The `remote` directory will contain the configuration files and setup scripts for our
production server.
- The `go.mod` file will declare our project dependencies, versions and module path.
- The `Makefile` will contain recipes for automating common administrative tasks — like
auditing our Go code, building binaries, and executing database migrations.

It’s important to point out that the directory name `internal` carries a special meaning and
behavior in Go: any packages that live under this directory can only be imported by code
inside the parent of the `internal` directory. In our case, this means that any packages that
live in `internal` can only be imported by code inside our `Redlight` project directory.

Or, looking at it the other way, this means that any packages under `internal` cannot be
imported by code outside of our project.
 
This is useful because it prevents other codebases from importing and relying on the
(potentially unversioned and unsupported) packages in our `internal` directory — even if
The project code is publicly available somewhere like GitHub.

###  Working with SQL Migrations
```
migrate create -seq -ext .sql -dir ./migrations create_movies_table
```
