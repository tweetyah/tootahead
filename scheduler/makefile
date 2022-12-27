build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/scheduler .
	docker build --tag tweetyah-scheduler:latest .

run: build
	docker rm -f tweetyah-scheduler
	docker run --name tweetyah-scheduler tweetyah-scheduler:latest

start: build
	docker rm -f tweetyah-scheduler
	docker run -d --name tweetyah-scheduler tweetyah-scheduler:latest

deploy:
	aws lightsail push-container-image --region us-east-1 --service-name tweetyah-scheduler --label tweetyah-scheduler --image tweetyah-scheduler:latest

publish: build deploy
	sleep 5
	export IMAGE=$(shell aws lightsail get-container-images --service-name tweetyah-scheduler --region us-east-1 | jq '.containerImages[0].image'); \
	aws lightsail create-container-service-deployment \
		--region us-east-1 \
		--service-name tweetyah-scheduler \
		--containers "{\"tweetyah-scheduler\":{\"image\":\"$$IMAGE\"}}" \
		--cli-read-timeout 1
