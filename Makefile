deploy:
	GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bootstrap
	lambroll deploy --alias="current" --envfile=.env
	rm -f bootstrap
run:
	lambroll invoke
