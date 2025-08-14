# Database
MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= wallet

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s | tr A-Z a-z)
export ARCH := $(shell uname -m)



# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile
include ./misc/make/help.Makefile

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

up: dev-env dev-rest dev-air             ## Startup / Spinup Docker Compose and air
down: docker-stop               ## Stop Docker
destroy: docker-teardown clean  ## Teardown (removes volumes, tmp files, etc...)

install-deps: migrate air gotestsum tparse mockery ## Install Development Dependencies (localy).
deps: $(MIGRATE) $(AIR) $(GOTESTSUM) $(TPARSE) $(MOCKERY) $(GOLANGCI) ## Checks for Global Development Dependencies.
deps:
	@echo "Required Tools Are Available"

dev-env: ## Bootstrap Environment (with a Docker-Compose help).
	@ docker-compose up -d --build mysql

dev-rest: ## Bootstrap Environment (with a Docker-Compose help).
	@ docker-compose up -d --build rest

dev-env-test: dev-env ## Run application (within a Docker-Compose help)
	@ $(MAKE) image-build
	docker-compose up web

dev-air: $(AIR) ## Starts AIR ( Continuous Development app).
	air

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

lint: $(GOLANGCI) ## Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

# -trimpath - will remove the filepathes from the reports, good to same money on network trafic,
#             focus on bug reports, and find issues fast.
# - race    - adds a racedetector, in case of racecondition, you can catch report with sentry.
#             https://golang.org/doc/articles/race_detector.html
#
# todo(butuzov): add additional flags to compiler to have an `version` flag.
build: ## Builds binary
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o engine \
		./app/
	@ echo "done"


build-race: ## Builds binary (with -race flag)
	@ printf "Building aplication with race flag... "
	@ go build \
		-trimpath  \
		-race      \
		-o engine \
		./app/
	@ echo "done"


go-generate: $(MOCKERY) ## Runs go generte ./...
	go generate ./...


TESTS_ARGS := --format testname --jsonfile gotestsum.json.out
TESTS_ARGS += --max-fails 2
TESTS_ARGS += -- ./...
TESTS_ARGS += -test.parallel 2
TESTS_ARGS += -test.count    1
TESTS_ARGS += -test.failfast
TESTS_ARGS += -test.coverprofile   coverage.out
TESTS_ARGS += -test.timeout        5s
TESTS_ARGS += -race

tests: $(GOTESTSUM)
	@ gotestsum $(TESTS_ARGS) -short

tests-complete: tests $(TPARSE) ## Run Tests & parse details
	@cat gotestsum.json.out | $(TPARSE) -all -notests

# ~~~ Docker Build ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.ONESHELL:
image-build:
	@ echo "Docker Build"
	@ DOCKER_BUILDKIT=0 docker build \
		--file Dockerfile \
		--tag withdrawal-balance \
			.

# Commenting this as this not relevant for the project, we load the DB data from the SQL file.
# please refer this when introducing the database schema migrations.

# # ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 

MYSQL_DSN := "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)?multiStatements=true"
MIGRATION_DIR := ./migrations
GOOSE := goose -dir $(MIGRATION_DIR) mysql $(MYSQL_DSN)

.PHONY: migrate-up
migrate-up:
	$(GOOSE) up

.PHONY: migrate-down
migrate-down:
	$(GOOSE) down

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then echo "Error: migration name is required. Example: make create name=create_users_table"; exit 1; fi
	$(GOOSE) create $(name) sql

# ~~~ Cleans ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

clean: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@printf "Cleanning artifacts... "
	@rm -f *.out
	@echo "done."


clean-docker: ## Removes dangling docker images
	@ docker image prune -f
