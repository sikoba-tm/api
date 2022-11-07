# SIKOBA API

## Installing Dependencies
```bash
go mod download
```

## Building Container
```bash
docker build -t sikoba-tm/api .
```

## Running the Container
```bash
docker run -d -p 8080:8080 sikoba-tm/api
```

## Development using Docker Compose
In order to start the API and Database at once, use the following command 
```bash
docker-compose up -d
```
To apply code changes, rebuild the container using
```bash
docker-compose build
```
Stop and destroy container using following command
```bash
docker-compose down
```


