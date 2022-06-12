![Logo](./logo.svg)

### Test task for the position of Trainee Golang Backend Developer

---

### [Версия на русском](./russian.md)

---

## Branches:

### - [`postgres`](https://github.com/illiafox/url-short-api/tree/pg) `PostgreSQL` implementation
### - [`mux`](https://github.com/illiafox/url-short-api/tree/mux)  [gorilla / mux](https://github.com/gorilla/mux) router

--- 

## docker-compose

API **starts immediately** after containers are up

Default port: `8080`

```shell
docker-compose up # make compose-up
```

---

## Endpoints

- ### `/new` - generate key
  **Request**: `POST`
    ```json
    { "url": "https://github.com/illiafox" }
    ```
  **Response**:
    ```json
    {
        "ok": true,
        "key": "3XdU_dFzP5"
    }
    ```

- ### `/get` - get link
  **Request**: `GET`
    ```json
    { "key": "3XdU_dFzP5" }
    ```
  **Response**:
    ```json
    {
        "ok": true,
        "url": "https://github.com/illiafox"
    }
    ```
- ### `/[key]` - auto redirect
  **Request**: `GET`
    ```shell
    curl -i http://localhost:8080/3XdU_dFzP5
    ```
  **Response**:
    ```shell
    HTTP/1.1 307 Temporary Redirect
    Content-Type: text/html; charset=utf-8
    Location: https://github.com/illiafox
    ... (browser redirection)
    ```

---

## Building and Running

```shell
make build
make run  # cd cmd/app && ./app
```

### With non-standard config and log file paths:
```shell
app -log=log.txt -config=config.toml
```

### HTTPS
```shell
app -https
```

### Use built-in cache storage:
```shell
app -cache # make run-cache
```

### With reading from environment:
Available keys can be found in the **[config tags](app/internal/config/struct.go)**
```shell
HOST_PORT=80 app
```

--- 

## Logs

```shell
# Terminal (stdout)
12:00:01 | INFO | Reading config
12:00:01 | INFO | Initializing service
12:00:01 | INFO | Initializing storage
12:00:01 | WARN | Using built-in storage
12:00:01 | INFO | Server started {"addr": "0.0.0.0:8080"}
13:15:01 | INFO | Shutting down server
```
```shell
# File (default log.txt)
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Reading config"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Initializing service"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Initializing storage"}
{"level":"warn","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Using built-in storage"}
{"level":"info","ts":"Fri, 10 Jun 2022 12:00:01 EEST","msg":"Server started","addr":"0.0.0.0:8080"}
{"level":"info","ts":"Fri, 10 Jun 2022 13:15:01 EEST","msg":"Shutting down server"}
```

---

## TODO (contribution is welcome):
1. Make branches with most popular routers

---

This is the first attempt in creating 'clean architecture', comments and pull requests are welcome.
