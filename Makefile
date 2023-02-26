
build:
	go build -o url-shortener cmd/main.go

run:
	go run cmd/main.go

docker-build:
	docker build -t url-shortener-app .

docker-run: docker-build
	docker run -it --rm url-shortener-app

docker-compose:
	docker compose up --build --force-recreate