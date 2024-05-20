BINARY_PATH := ./bin/app
BINARY_MIGRATE_PATH := ./bin/migrate
BINARY_TEST_SCRIPT_PATH := ./bin/test
MAIN_PATH := ./cmd/app/main.go
MIGRATE_PATH := ./cmd/migrate/main.go
TEST_SCRIPT_PATH := ./cmd/test/main.go

build:
	@echo 'Building golang'
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH} ${MAIN_PATH}	
	@echo 'Build completed'

build-migrate:
	@echo 'Building migrate script'
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_MIGRATE_PATH} ${MIGRATE_PATH}	
	@echo 'Build completed'

build-test-script:
	@echo 'Building test script' 
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_TEST_SCRIPT_PATH} ${TEST_SCRIPT_PATH}	
	@echo 'Build completed'

run: build
	@echo 'Starting server'
	@${BINARY_PATH}

migrate: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH}
	@echo 'End migrating'

migrate-reset: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH} --run=reset
	@echo 'End migrating'

migrate-drop: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH} --run=drop
	@echo 'End migrating'

test-script: build-test-script
	@echo 'Launch test script'
	@${BINARY_TEST_SCRIPT_PATH}
