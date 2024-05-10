run: 
	docker compose up
build:
	go build -o ./bin/api cmd/api/main.go;
docker_build: 
	docker build -t m0rdovorot/effective_mobile.cars:latest -f API.Dockerfile .;
docker_push: 
	docker push m0rdovorot/effective_mobile.cars:latest;
get_cars:
	curl "http://localhost:8001/api/v1/car?regNum=XXX09XX&&mark=dasda&&page=1"
create_car:
	curl -X POST -H "Content-Type: application/json" -d '{"regNums":["ABCDEF", "ZZZZZZ"]}' "http://localhost:8001/api/v1/car"
patch_car:
	curl -X PATCH -H "Content-Type: application/json" -d '{"name": "chel", "surname": "chelov"}' "http://localhost:8001/api/v1/car/5"
delete_car:
	curl -X DELETE "http://localhost:8001/api/v1/car/7"
install_swagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@abb53530bfcf49c470e5f4c7071ee43f37ec7437
swagger:
	swagger serve -F=swagger swagger.yml
