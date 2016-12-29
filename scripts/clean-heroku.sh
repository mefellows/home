#!/bin/sh

echo "Wiping and re-seeding Heroku DB! (note: this will overwrite your .env file)"

heroku addons:destroy heroku-postgresql:hobby-dev -f
heroku addons:create heroku-postgresql:hobby-dev

cat << EOF > .env
LOG_LEVEL=DEBUG
PORT=8000
EOF

URL=`heroku config:get DATABASE_URL`
echo "DATABASE_URL=$URL" >> .env
migrate -url $URL -path ./db/migrations up
make seed

rm .env
echo "Migration complete - you should probably now run 'make docker' to reset your development environment"
