# Clean architecture example

This exercise examplifies:

- Clean architecture
- Testing
- Mocking
- Dockerizing
- Linting

## Technologies

- Go
- PostgreSQL
- Docker and docker-compose
- RESTful API
- golangci-lint

## TODO

### Setup architecture

- [X] Implement linting
- [X] Implement entity layer
- [X] Implement handler layer
- [X] Implement controller layer
- [X] Implement repository layer
- [X] Initialize architecture in main
- [X] Implement go mod
- [ ] Implement GitHub Actions

### Create

- [X] Implement POST method for handler layer
- [X] Implement mocks for Create method for controller layer
- [X] Implement Create method for controller layer
- [X] Implement mocks for Create method for repository layer
- [X] Implement Create method for repository layer
- [ ] Implement integration test for Create method

### List

- [X] Implement GET method for handler layer
- [X] Implement List method for controller layer
- [X] Implement List method for repository layer
- [X] Implement mocks for List method
- [ ] Implement integration test for List method

### Detail

- [ ] Implement GET method for handler layer
- [ ] Implement Detail method for controller layer
- [ ] Implement Detail method for repository layer
- [ ] Implement mocks for Detail method
- [ ] Implement integration test for Detail method

### Delete

- [ ] Implement DELETE method for handler layer
- [ ] Implement Delete method for controller layer
- [ ] Implement Delete method for repository layer
- [ ] Implement mocks for Delete method
- [ ] Implement integration test for Delete method

### Docker

- [ ] Dockerize Go
- [ ] Dockerize PostgreSQL
- [ ] Implement docker-compose

### Technical Debt

- [ ] Implement decoder as middleware
- [ ] Implement encoder as middleware