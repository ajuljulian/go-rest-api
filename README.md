# About

Based on the following course: https://tutorialedge.net/courses/go-rest-api-course-v2/

Create a REST api for CRUD operations on comments

Persists the data to a Postgres database

Uses JWT tokens to secure access to certain endpoints

Uses docker-compose for orchestrations

Uses Taskfiles as a task runner: https://taskfile.dev/#/installation

# Deployment

```
$ task run
```

# Testing

### Integration tests

```
$ task integration-tests
```

### Acceptance tests

```
$ task acceptance-tests
```

### Go Libraries

```
github.com/stretchr/testify/assert
github.com/satori/go.uuid
github.com/jmoiron/sqlx
github.com/lib/pq

github.com/golang-migrate/migrate/v4
github.com/golang-migrate/migrate/v4/database/postgres
github.com/golang-migrate/migrate/v4/source/file

github.com/go-playground/validator/v10
github.com/gorilla/mux
```

# Misc

JWT Tokens: https://jwt.io/
