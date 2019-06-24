ensure_outside_container:
	if [ -f /.dockerenv ]; then echo -n $(RED)This target needs to be run outside server container$(NC)"\n"; false; fi

up: ensure_outside_container
	@docker-compose pull && docker-compose up

rm: ensure_outside_container
	docker-compose down -v

enter: ensure_outside_container
	@docker-compose exec server bash

mycli: SHELL:=/bin/bash
mycli: ensure_outside_container
	mycli -h 127.0.0.1 -P 3306 -u root -p topsecret

migrate: ensure_outside_container
	sql-migrate up  -env multiple-apartments
	sql-migrate up  -env open-requests
	sql-migrate up  -env close-all-requests
