# Withdrawal Balance

### How To Run This Project

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
# copy the example.env to .env
$ cp example.env .env
$ cp example.env app/.env

# update database credential in .env and Makefile before run migrate

# Run the application
$ go mod tidy
$ make dev-env
$ make image-build
$ make up
```

#### Migration Database

```bash
# Install goose
$ go intall github.com/pressly/goose/v3/cmd/goose
$ make migrate-up
```
