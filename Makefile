REGISTRY = localhost:5000
IMAGE_SERVICE_AA = service-acct-auth:V0.0.2
PORT_MINIO = 9000
HELM_SAA_VERSIOM = 0.1.0

docker-build-service-acct-auth:
	docker build -t $(IMAGE_SERVICE_AA) -f cmd/service-acct-auth/Dockerfile .

docker-image-put-registry-service-acct-auth:
	docker tag $(IMAGE_SERVICE_AA) $(REGISTRY)/$(IMAGE_SERVICE_AA)
	docker push $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-service-acct-auth:
	docker run -d --name test-$(IMAGE_NAME) -p 50051:50051 $(REGISTRY)/$(IMAGE_SERVICE_AA)

docker-run-registry:
	docker run -d -p 5000:5000 --name registry -v registry-data:/var/lib/registry registry:2

docker-run-minio:
	docker run -d --name minio -p $(PORT_MINIO):$(PORT_MINIO) -p 9001:9001 -e "MINIO_ROOT_USER=admin" -e "MINIO_ROOT_PASSWORD=secretpassword" minio/minio server /data --console-address ":9001"

docker-run-prometheus:
	docker run -d --name prometheus -p 9090:9090 -v "C:/Users/pilip/go/src/github.com/piliphulko/marketplace-example/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yaml" -v prometheus-data:/prometheus prom/prometheus   --config.file=/etc/prometheus/prometheus.yaml --storage.tsdb.retention.time=2h --storage.tsdb.min-block-duration=2h --storage.tsdb.max-block-duration=2h

docker-run-thanos-sidecar:
	docker run -d --name thanos-sidecar -p 10902:10902 -v prometheus-data:/prometheus -v "C:/Users/pilip/go/src/github.com/piliphulko/marketplace-example/prometheus/thanos/thanos.yaml:/etc/thanos/thanos.yaml"  thanosio/thanos:main-2024-09-02-dfeaf6e sidecar --tsdb.path=/prometheus --objstore.config-file=/etc/thanos/thanos.yaml --prometheus.url=http://host.docker.internal:9090

helm-install-service-acct-auth:
	helm install release-saa oci://$(REGISTRY)/helm-repo/service-acct-auth --version $(HELM_SAA_VERSIOM)

helm-install-service-acct-auth:
	helm uninstall release-saa