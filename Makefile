fmt:
	go fmt ./...

run:
	go build
	cp trippaio-bot /usr/local/bin/trippaio-bot
	trippaio-bot