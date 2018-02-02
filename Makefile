.PHONY: install build serve clean pack deploy ship

TAG=$$(git rev-list HEAD --max-count=1 --abbrev-commit)

install:
	go get .

build:
	go build -ldflags "-X main.version=$(TAG)" -o botserver .

serve: build
	./botserver

clean:
	rm ./botserver

pack:
	GOOS=linux make build
	docker build -t eu.gcr.io/botserver-1337/botserver:latest -t eu.gcr.io/botserver-1337/botserver:$(TAG) .

upload:
	gcloud docker -- push eu.gcr.io/botserver-1337/botserver:$(TAG)

deploy:
	sed "s/VERSION/$(TAG)/g" deployment.yaml | kubectl apply -f -

ship: pack upload deploy clean

run:
	docker run -d --name botserver -p 80:80 -p 5557:5557 -p 6555:6555 eu.gcr.io/botserver-1337/botserver:latest

stop:
	docker stop botserver
	docker rm botserver