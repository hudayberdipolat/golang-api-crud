source .env
migrate -path pkg/database/migrations -database "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sllmode=$DB_SSL_MODE" -verbose $1 $2 $3