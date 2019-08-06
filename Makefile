CONTAINER_PREFIX:=vue-go-sample
# Docker actions

.PHONY: dstart dstop dstatus dlogin dclean dlog
setup:
	@echo "Start setup"

# start container
dstart: setup
	@echo "docker start"
	@docker-compose up -d

# stop container
dstop:
	@echo "docker stop"
	@docker-compose stop

# restart container
drestart:
	@make dstop
	@make dstart

# show container's status
dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

# attach container
dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/sh

# clean up containers
dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force

# show container's log
dlog:
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)
