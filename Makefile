ifeq ($(OS),Windows_NT)
SHELL := cmd.exe
.SHELLFLAGS := /C
endif

include .env
DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}

# Scaffold helpers (Windows-friendly via PowerShell)
ARG := $(word 2,$(MAKECMDGOALS))
NAME ?= $(ARG)

ifeq (controller,$(firstword $(MAKECMDGOALS)))
ifneq ($(strip $(ARG)),)
$(eval $(ARG):;@echo.)
endif
endif

ifeq (request,$(firstword $(MAKECMDGOALS)))
ifneq ($(strip $(ARG)),)
$(eval $(ARG):;@echo.)
endif
endif

ifeq (service,$(firstword $(MAKECMDGOALS)))
ifneq ($(strip $(ARG)),)
$(eval $(ARG):;@echo.)
endif
endif

start-db:
	docker start postgres

stop-db:
	docker stop postgres

createdb:
	docker exec -it postgres createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres dropdb --username=${DB_USER} ${DB_NAME}

migrate-up:
	migrate -path db/migrations -database ${DB_URL} up $(version)

migrate-down:
	migrate -path db/migrations -database ${DB_URL} down $(version)

migrate-force:
	migrate -path db/migrations -database ${DB_URL} force $(version)

migration-create:
	migrate create -ext -sql -dir db/migrations $(name)

controller:
	@powershell -NoProfile -ExecutionPolicy Bypass -Command "& { param([string]$$Name) $$ErrorActionPreference='Stop'; if([string]::IsNullOrWhiteSpace($$Name)) { Write-Host 'Usage: make controller <Name>  (or: make controller NAME=<Name>)'; exit 1 } $$dir='controllers'; New-Item -ItemType Directory -Force -Path $$dir | Out-Null; $$file=Join-Path $$dir ($$Name + 'Controller.go'); if(Test-Path $$file) { Write-Host ('Already exists: ' + $$file); exit 1 } Set-Content -Path $$file -Encoding utf8 -Value @('package controllers','','type ' + $$Name + 'Controller struct {','}',''); Write-Host ('Created ' + $$file) }" "$(NAME)"

request:
	@powershell -NoProfile -ExecutionPolicy Bypass -Command "& { param([string]$$Name) $$ErrorActionPreference='Stop'; if([string]::IsNullOrWhiteSpace($$Name)) { Write-Host 'Usage: make request <Name>  (or: make request NAME=<Name>)'; exit 1 } $$dir='requests'; New-Item -ItemType Directory -Force -Path $$dir | Out-Null; $$file=Join-Path $$dir ($$Name + 'Request.go'); if(Test-Path $$file) { Write-Host ('Already exists: ' + $$file); exit 1 } Set-Content -Path $$file -Encoding utf8 -Value @('package requests','','type ' + $$Name + 'Request struct {','}',''); Write-Host ('Created ' + $$file) }" "$(NAME)"

service:
	@powershell -NoProfile -ExecutionPolicy Bypass -Command "& { param([string]$$Name) $$ErrorActionPreference='Stop'; if([string]::IsNullOrWhiteSpace($$Name)) { Write-Host 'Usage: make service <Name>  (or: make service NAME=<Name>)'; exit 1 } $$dir='services'; New-Item -ItemType Directory -Force -Path $$dir | Out-Null; $$file=Join-Path $$dir ($$Name + 'Service.go'); if(Test-Path $$file) { Write-Host ('Already exists: ' + $$file); exit 1 } Set-Content -Path $$file -Encoding utf8 -Value @('package services','','import (','    "gorm.io/gorm"',')','','type ' + $$Name + 'Service struct {','    DB *gorm.DB','}','','func New' + $$Name + 'Service(db *gorm.DB) *' + $$Name + 'Service {','    return &' + $$Name + 'Service{DB: db}','}',''); Write-Host ('Created ' + $$file) }" "$(NAME)"

.PHONY: start-db stop-db createdb dropdb migrate-up migrate-down migrate-force migration-create controller request service