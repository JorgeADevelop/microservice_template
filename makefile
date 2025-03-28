DB_USER := $(shell grep DB_USER .env.local | cut -d '=' -f 2)
DB_PASSWORD := $(shell grep DB_PASSWORD .env.local | cut -d '=' -f 2)
DB_HOST := $(shell grep DB_HOST .env.local | cut -d '=' -f 2)
DB_PORT := $(shell grep DB_PORT .env.local | cut -d '=' -f 2)
DB_DATABASE := $(shell grep DB_DATABASE .env.local | cut -d '=' -f 2)

migration-files:
	@read -p "Enter the title: " title; \
	migrate create -ext sql -dir ./db/migrations -seq $$title

migrate-up:
	migrate -path ./db/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)" up

migrate-down:
	migrate -path ./db/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)" down
