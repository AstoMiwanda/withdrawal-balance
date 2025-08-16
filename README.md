# Withdrawal Balance

### How To Run This Project

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
# copy the example.env to .env
$ cp example.env .env

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


### APIs Docs
#### Check balance
```bash
# cURL
curl --location --request POST 'localhost:9090/v1/wallet/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1,
    "wallet_id": 1
}'
```

#### Withdrawal balance
```bash
# cURL
curl --location --request POST 'localhost:9090/v1/wallet/withdrawal-balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1,
    "wallet_id": 1,
    "amount": 1000.00,
    "bank_code": "021",
    "bank_account_number": "0000000000",
    "bank_account_name": "Test"
}'
```