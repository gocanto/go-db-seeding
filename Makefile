.PHONY: gocanto\:db-seeding

run-db:
	docker run --name db-seeding -e POSTGRES_PASSWORD=pwd -e POSTGRES_DB=gus-db -d -p 5432:5432 postgres

db-bash:
	docker container exec -it db-seeding bash
