BUILD=./app/cmd/app

all: clean build run

run:
	(cd $(BUILD) && ./app)

run cache:
	(cd $(BUILD) && ./app -cache)

build: clean
	go build -o $(BUILD)/app $(BUILD)

clean:
	if [ -f $(BUILD)/app ]; then rm $(BUILD)/app; fi

compose-up: compose-down
	docker-compose up -d

compose-down:
	docker-compose down