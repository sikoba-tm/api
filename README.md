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