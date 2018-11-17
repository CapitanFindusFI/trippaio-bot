fmt:
	go fmt ./...

run:
	go build -o /usr/local/bin/trippaio-bot
	trippaio-bot