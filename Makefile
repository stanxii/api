# Run command local by default
all: local

# Build and run the application locally
local:
	clear && go build api.go && ./api

# Build a Docker image of the application
image:
	docker build --tag=snapshot-api-gateway .

# Run the application in a Docker container (after an image has been built)
docker:
	docker run --env-file .env -p 8000:80 snapshot-api-gateway

# Delete local compiled binary
clean:
	rm api
