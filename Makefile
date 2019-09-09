.SILENT:
default:help;

CONTAINER=docker exec -it htmltopdf

#--------> DOCKER UTILS
## restart local environment
restart: stop start

## start local environment
start:
	docker-compose -f docker-compose.yml up -d && $(CONTAINER) go run main.go

## stop local environment
stop:
	docker-compose -f docker-compose.yml down

bash:
	$(CONTAINER) bash