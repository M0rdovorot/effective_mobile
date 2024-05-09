run: 
	docker compose up
build:
	go build -o ./bin/api cmd/api/main.go;
docker_build: 
	docker build -t m0rdovorot/effective_mobile.banners:latest -f API.Dockerfile .;
docker_push: 
	docker push m0rdovorot/effective_mobile.banners:latest;
user_banner: 
	curl "http://localhost:8001/api/v1/user_banner?tag_id=1&&feature_id=1"
user_banner_actual:
	curl "http://localhost:8001/api/v1/user_banner?feature_id=9&&tag_id=9&&use_last_revision=true"
create_banner:
	curl -X POST -H "Content-Type: application/json" -d '{"tag_ids":[10], "feature_id": 10, "content":{"title":"some_title"}, "is_active": false}' "http://localhost:8001/api/v1/banner"
patch_banner:
	curl -X PATCH -H "Content-Type: application/json" -d '{"tag_ids":[9], "feature_id": 9, "content":{"title":"some_title"}, "is_active": true}' "http://localhost:8001/api/v1/banner/19"
delete_banner:
	curl -X DELETE "http://localhost:8001/api/v1/banner/18"
test:
	docker compose down
	docker compose -f ./cmd/api/tests/docker_compose.yml up -d
	go test cmd/api/tests/main_test.go
	docker compose -f ./cmd/api/tests/docker_compose.yml down -d