APP_NAME = "qiyue_api_server"

default:
	go build -o ${APP_NAME} main.go

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