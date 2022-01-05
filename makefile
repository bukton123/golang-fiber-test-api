docker: docker-app docker-service

docker-app:
	docker-compose -f deployment/docker/docker-compose.yml up -d

docker-service:
	docker-compose -f deployment/docker/docker-service.yml up -d
	
docker-service-user:
	docker-compose -f deployment/docker/docker-service-users.yml up --build
