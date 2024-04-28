build:
	go build -o ./bin/app ./internal/app/app.go

run:
	go run ./internal/app/app.go

docker-build:
	docker build -t "template-app:latest" .

docker-run: docker-build
	docker run -it -p 80:80 template-app

deploy:
	cdk synth
	cdk deploy