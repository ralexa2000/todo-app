#!/bin/zsh

export MIGRATION_DIR="migrations"
export DB_DSN="host=localhost port=5432 dbname=tasks user=user password=pwd sslmode=disable"

if [ "$1" = "--dry-run" ]; then
  goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status
else
  goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
fi
