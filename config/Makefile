docker-volume-create:

	docker volume create volume-config

docker-download-config:

	docker run --rm -d -v /config.yaml:/config/ volume-config volume-config-image

docker-delete-volume:

	docker volume rm volume-config