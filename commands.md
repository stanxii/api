# Useful commands

## Containers
- Build the Docker image
```
docker build --tag=snapshot-api-gateway .
```
- Run the Docker image locally
```
docker run -p 8000:80 snapshot-api-gateway
```

## Database
- Connect to PostgreSQL
```
psql snapshot -U admin
```

- List all tables
```
\d
```

- List all users
```
\du
```

- Quit PostgreSQL
```
\q
```
