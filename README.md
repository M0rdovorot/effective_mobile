## Сборка
Для сборки введите в корне проекта:
```
	docker build -t <Название образа> -f API.Dockerfile .;
```
или
```
	make docker_build
```
## Запуск
Для запуска введите в корне проекта:
```
  docker compose up
```
или
```
  make run
```
## Тесты
Перейдите в папку cmd/api/tests и введите:
```
	docker compose up -d
	go test
	docker compose down
```
или из корня проекта
```
  make test
```
## Нагрузочное тестирование
Для нагрузочного тестирования будет использоваться тестовая база данных, которую можно поднять, посмотрев предыдущий пункт про тестирование
Для нагрузочного тестирования используется утилита ab
При нагрузке в 50 потоков в 10000 запросов на http://localhost:8001/api/v1/user_banner?tag_id=14&&feature_id=13
```
ab -c 50 -n 1000 -H "token: admin_token" 'http://localhost:8001/api/v1/user_banner?tag_id=14&&feature_id=13'
```
![alt text](static/image1.png)

При нагрузке в 50 потоков в 10000 запросов на http://localhost:8001/api/v1/banner
```
 ab -c 50 -n 10000 -H "token: admin_token" http://localhost:8001/api/v1/banner
```
![alt text](static/image2.png)