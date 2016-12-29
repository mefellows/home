#!/bin/sh -e

DOCKER_MACHINE_NAME=$(docker-machine active)
DOCKER_MACHINE_HOST=$(docker-machine ip $DOCKER_MACHINE_NAME)
POSTGRES_HOST=$DOCKER_MACHINE_HOST
echo ">> Current docker machine: $DOCKER_MACHINE_NAME@$DOCKER_MACHINE_HOST"

function clean() {
  echo ">> Creating new Postgres data volume on Docker Machine: '$DOCKER_MACHINE_HOST'"
  docker-machine ssh $DOCKER_MACHINE_NAME 'sudo rm -rf /tmp/volumes/postgres'
  docker-machine ssh $DOCKER_MACHINE_NAME 'sudo mkdir -p /tmp/volumes/postgres'
}

echo "Running DB migrations for Local Docker Development"
docker-compose stop
docker-compose rm -f
clean
docker-compose up -d

echo ">> Waiting for postgres to start"
WAIT=0
while ! nc -z $POSTGRES_HOST 5432; do
  sleep 1
  WAIT=$(($WAIT + 1))
  if [ "$WAIT" -gt 15 ]; then
    echo "Error: Timeout wating for Postgres to start"
    exit 1
  fi
done

# Allow time to boot
sleep 5


if [ "$1" = "" ]; then
  cat << EOF > .env
LOG_LEVEL=DEBUG
DATABASE_URL="postgres://admin:admin@$POSTGRES_HOST:5432/admin?sslmode=disable"
PORT=8000
EOF
fi

echo "Running DB migrations for Local Docker Development"
# migrate -url postgres://admin:admin@$POSTGRES_HOST:5432/admin?sslmode=disable -path ./db/migrations up
make seed
make docker
