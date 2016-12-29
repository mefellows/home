#!/bin/sh

DOCKER_MACHINE_NAME=$(docker-machine active)
DOCKER_MACHINE_HOST=$(docker-machine ip $DOCKER_MACHINE_NAME)
POSTGRES_HOST=$DOCKER_MACHINE_HOST
docker-compose up -d

#echo "Running DB migrations for Local Docker Development"
#migrate -url postgres://admin:admin@$POSTGRES_HOST:5432/admin?sslmode=disable -path ./db/migrations up

if [ "$1" = "" ]; then
  cat << EOF > .env
LOG_LEVEL=DEBUG
DATABASE_URL="postgres://admin:admin@$POSTGRES_HOST:5432/admin?sslmode=disable"
REDIS_URL="redis://h:password@$POSTGRES_HOST:6379"
PORT=8000
EOF
fi
