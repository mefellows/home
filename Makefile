TEST?=./...

default: test

bin:
	@sh -c "$(CURDIR)/scripts/build.sh"

docker:
	sh -c "$(CURDIR)/scripts/docker.sh"

clean-docker:
	sh -c "$(CURDIR)/scripts/clean-docker.sh"

clean-heroku:
	sh -c "$(CURDIR)/scripts/clean-heroku.sh"

psql:
	docker exec -it --user postgres home_db_1 psql admin

seed:
	sh -c "$(CURDIR)/scripts/seed.sh"

run:
	sh -c "$(CURDIR)/scripts/run.sh"

test:
	"$(CURDIR)/scripts/test.sh"

testrace:
	go test -race $(TEST) $(TESTARGS)


updatedeps:
	go get -d -v -p 2 ./...



.PHONY: bin default dev test updatedeps
