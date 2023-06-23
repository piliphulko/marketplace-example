build-linux:

	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -tags netgo -installsuffix cgo -o service-acct-auth

build-linux-pgo:

	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -tags netgo -installsuffix cgo -pgo=profile.pprof -o service-acct-auth

build-docker-image:

	docker build -t service-acct-auth .