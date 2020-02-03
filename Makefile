.PHONY: install ${env}
install: 
	echo '**** building app with arguments: env = ${env} ****'
ifeq ($(env), local)
	go build -o ./bin/ .
	cp -f ./configs/application.${env}.env ./bin/application.env 
else
	env GOOS=linux GOARCH=amd64 go build -o ./publish/ .
	cp -f ./configs/application.${env}.env ./publish/application.env 
endif

.PHONY: docker 
docker:
	rm -rf /root && docker build --rm -t app .