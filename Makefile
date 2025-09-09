# Service
export PORT=8080

# JWT
export JWT_SECRET_KEY = aJm5+7P3tFbDqZHTq8gW4nqoxzWJ3i28bM6oW8Zz9Xg=
export ISSUER = user-service
export AUDIENCE = user-service-client

# Health Check
export ENV=dev
export VERSION=1.0.0

# DB for log
export MONGO_IP=127.0.0.1
export MONGO_PORT=27017
export MONGO_DB_USERNAME=user
export MONGO_DB_PASSWORD=P%40ssw0rd

run-main:
	go run cmd/main.go