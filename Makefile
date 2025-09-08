# Service
export PORT=8080
export SECRET_KEY = e6d8a47947531b759a51458e150c909d

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