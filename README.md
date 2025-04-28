# Task

Develop a system that exposes CRUD API endpoints. The system should manage a Documenting System, where each document has the following fields: ID, Title, Author, Content, Created_at, and Updated_at.
The service must be implemented in Go, with data stored in PostgreSQL running inside Docker.

The API endpoints should include:
- Save Document [POST]
- Update Document [PUT]
- Delete Document [DELETE]
- Get Document by ID [GET]
- Get All Documents [GET]

## Tasks

- [x] Initialize the project
- [x] Setup the environment (docker)
- [ ] Create the project structure
- [ ] Create commands using `Cobra` for:
    1. [ ] running server
    2. [ ] running migrations
    3. [ ] running seeds
- [ ] CRUD APIs
- [ ] Test Cases
- [ ] Postman Collection
- [ ] README
- [ ] Deploy the project
- [ ] Run Load test

## We can use

- [ ] Docker
- [ ] Golang ( GIN, GORM, Cobra, golang-migrate, logrus)
- [ ] k6 (load test)
