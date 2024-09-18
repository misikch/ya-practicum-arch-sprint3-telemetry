# ya-practicum-arch-sprint3-telemetry
### Telemetry service

Сервис получает телеметрию с устройств и сохраняет в БД, а также публикует в databus для др. сервисов, 
например для notification service уведомлений пользователей.

### Как работать с сервисом

Все api ручки описаны в файле `api.yaml`, после изменения файла нужно вызвать команду кодогенерации:
```shell
go generate ./...
```

Сам сервис находится здесь `/cmd/service/main.go`, он обрабатывает api методы.
Еще есть worker, который слушает databus topic `device_topic` 
на наличие новых/обновление существующих устройств, после чего пишет данные в таблицу `telemetry_devices`.

### Как запустить локально
Поднимаем окружение: mongodb & kafka
```shell
docker-compose up
```

Подготавливаем переменные окружения для доступа к mongodb & kafka
```shell
cp .env.example .env
```

Запускаем service/worker

Выполняем запросы к api из файла `/test/test.http`


### Как остановить окружение локально
```shell
docker-compose down --remove-orphan
```

