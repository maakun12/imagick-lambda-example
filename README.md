# imagick-lambda-example

## Setup
```
docker-compose up
```

## Usage
```
docker-compose exec app go install
docker-compose exec app env GOOS=linux go build -o main main.go
```
