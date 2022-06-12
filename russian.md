![Logo](./logo.svg)
### Задание на позицию Trainee Backend Developer


---

### [English Translation](./readme.md)

---

## Другие ветки:

### - [`postgres`](https://github.com/illiafox/url-short-api/tree/pg) реализация с `PostgreSQL`
### - [`mux`](https://github.com/illiafox/url-short-api/tree/mux)  роутер [gorilla / mux](https://github.com/gorilla/mux)

--- 

## docker-compose

API **готова к работе** после поднятия контейнеров 

Порт по-умолчанию: `8080`

```shell
docker-compose up # make compose-up
```

---

## Endpoints

- ### `/new` - генерация кода
    **Запрос**: `POST`
    ```json
    { "url": "https://github.com/illiafox" }
    ```
    **Ответ**:
    ```json
    {
        "ok": true,
        "key": "3XdU_dFzP5"
    }
    ```
  
- ### `/get` - получение ссылки
    **Запрос**: `GET`
    ```json
    { "key": "3XdU_dFzP5" }
    ```
    **Ответ**:
    ```json
    {
        "ok": true,
        "url": "https://github.com/illiafox"
    }
    ```
- ### `/[key]` - автоматический redirect
    **Запрос**: `GET`
    ```shell
    curl -i http://localhost:8080/3XdU_dFzP5
    ```
    **Ответ**:
    ```shell
    HTTP/1.1 307 Temporary Redirect
    Content-Type: text/html; charset=utf-8
    Location: https://github.com/illiafox
    ... (переадресация на сайт)
    ```


---

## Сборка и запуск

```shell
make build
make run  # cd cmd/app && ./app
```

### Задать пути к конфигу и лог файлу:
```shell
app -log=log.txt -config=config.toml
```

### HTTPS
```shell
app -https
```

### Использовать встроенное хранилище:
```shell
app -cache # make run-cache
```

### Изменить значения через переменные среды:
Имена заданы в **[тегах конфига](app/internal/config/struct.go)**
```shell
HOST_PORT=80 app
```

--- 

## Логи

```shell
# терминал (stdout)
12:00:01 | INFO | Reading config
12:00:01 | INFO | Initializing service
12:00:01 | INFO | Initializing storage
12:00:01 | WARN | Using built-in storage
12:00:01 | INFO | Server started {"addr": "0.0.0.0:8080"}
13:15:01 | INFO | Shutting down server
```
```shell
# файл (по-умолчанию log.txt)
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Reading config"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Initializing service"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Initializing storage"}
{"level":"warn","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Using built-in storage"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Server started","addr":"0.0.0.0:8080"}
{"level":"info","ts":"Fri, 10 Jun 2022 13:15:01 EEST","msg":"Shutting down server"}
```

---

## TODO (contribution is welcome):
1. Заменить реализацию хранилища с sync.RWMutex на более быстрый атомарный вариант
2. Сделать ветки со всеми роутерами

---

### P.S. 
Это первая попытка в `clean architecture`, замечания и pull реквесты приветствуются.
