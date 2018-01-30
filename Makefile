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
	docker build -t eu.gcr.io/botserver-1337/botserver:$(TAG) .

upload:
	gcloud docker -- push eu.gcr.io/botserver-1337/botserver:$(TAG)

deploy:
	sed "s/VERSION/$(TAG)/g" deployment.yaml | kubectl apply -f -

ship: pack upload deploy clean