all:
	docker-compose stop bigfile
	docker-compose build
	docker-compose up -d bigfile
	docker logs -f nififfbig_bigfile_1
