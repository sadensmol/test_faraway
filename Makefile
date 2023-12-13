.PHONY:run
run: 
	docker compose -f ./infrastructure/local/docker-compose.yaml up --build
	
run-server:
	docker compose -f ./infrastructure/local/docker-compose.yaml up server
	
run-client:
	docker compose -f ./infrastructure/local/docker-compose.yaml up client