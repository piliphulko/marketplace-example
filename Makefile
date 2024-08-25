REGISTRY = localhost:5000
IMAGE_SERVICE_AA = service-acct-auth:V0.0.2

docker-build-service-acct-auth:
	docker build -t $(IMAGE_SERVICE_AA) -f cmd/service-acct-auth/Dockerfile .

docker-image-put-registry-service-acct-auth:
	docker tag $(IMAGE_SERVICE_AA) $(REGISTRY)/$(IMAGE_SERVICE_AA)
	docker push $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-service-acct-auth:
	docker run -d --name test-$(IMAGE_NAME) -p 50051:50051 $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-registry:
	docker run -d -p 5000:5000 --name registry -v registry-data:/var/lib/registry registry:2

helm-install-service-acct-auth:
	helm install --debug service-acct-auth ./k8s/cores-api --set SELECTED=SERVICE-ACCT-AUTH

helm-upgrade-service-acct-auth:
	helm upgrade service-acct-auth ./k8s/cores-api --set SELECTED=SERVICE-ACCT-AUTH

helm-uninstall-service-acct-auth:
	helm uninstall service-acct-auth 