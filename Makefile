APP_NAME = "qiyue_api"

default:
	env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	godep restore

dev:
	fresh -c config/dev.conf

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file"
	@echo "make dev - run go fresh"
	@echo "make install - install dep"