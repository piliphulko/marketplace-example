docker-build-service-acct-auth:

	docker build -t service-acct-auth -f cmd/service-acct-auth/Dockerfile .

docker-run-service-acct-auth:

	docker run -d --name test-service-acct-auth -p 50051:50051 service-acct-auth