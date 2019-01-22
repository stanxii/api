# Snapshot API Gateway
This is the API gateway for Snapshot, a scalable photo sharing social network built to practice microservices based system design.
It exposes a GraphQL endpoint so the client can communicate with the back end microservices.

It's built with
- [Go](https://golang.org/)
- [net/http](https://golang.org/pkg/net/http/)
- [GraphQL](https://graphql.github.io/)

## Quick Start
You will need [Go (>= version 1.11)](https://golang.org/dl/) and [Docker](https://www.docker.com/) installed.

- Build and run the application locally
```
make
```

- Build a Docker image of the application
```
make image
```

- Run the application in a Docker container (after an image has been built).
You will need the environment variables defined in a file called `.env`.
```
make docker
```

- Once the API is running locally, you can hit it at `localhost:8000/graphql`.
If it's running in a local container, you can hit it at `localhost/graphql`.

- Delete the local compiled binary
```
make clean
```

## Deployment
For deployment instructions, see [deployment.md](/deployment.md).

## See also
For a list of other useful commands, see [commands.md](/commands.md).
