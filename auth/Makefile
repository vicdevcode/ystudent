BINARY_PATH := ./bin/app
BINARY_MIGRATE_PATH := ./bin/migrate
MAIN_PATH := ./cmd/app/main.go
MIGRATE_PATH := ./cmd/migrate/main.go

build:
	@echo 'Building golang ...'
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH} ${MAIN_PATH}	
	@echo 'Build completed'

run: build
	@echo 'Starting server'
	@${BINARY_PATH}

migrate:
	@echo 'Start migrating'
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_MIGRATE_PATH} ${MIGRATE_PATH}	
	@${BINARY_MIGRATE_PATH}
	@echo 'End migrating'
