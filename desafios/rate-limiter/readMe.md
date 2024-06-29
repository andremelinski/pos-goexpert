# Go Expert Challenge - Rate Limiter

Implementação de um rate limiter em Go para um serviço Web capaz de limitar o número de requisições recebidas de clientes dentro de um intervalo de tempo configurável, observando o endereço IP e/ou token de acesso `API_KEY`.

## Arquitetura

A aplicação é composta por um servidor web que recebe requisições HTTP e um middleware de rate limiter que é responsável por controlar o número de requisições recebidas. O [middleware](internal/infra/web/webserver/middleware/middleware.go) intercepta todas as requisições e realiza toda a chacagem pelo [strategy](internal/infra/web/webserver/middleware/strategy/rate-limiter.go). Para armazenar os dados é utilizado o Redis por ser um banco de fácil uso e que possui a opção de expiração do dado após um determinado tempo.

O rate limiter é configurável por IP ou token `API_KEY` no arquivo `.env`. Exemplo de configuração para limitar 10 requisições por IP e 100 requisições por token em uma janela de tempo de 5 minutos:

```sh
RATE_LIMITER_IP_MAX_REQUESTS=10
RATE_LIMITER_TOKEN_MAX_REQUESTS=100
RATE_LIMITER_TIME_WINDOW_MILISECONDS=300000
```

## Executando o projeto

**Obs:** é necessário ter o [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

1. Crie um arquivo `.env` na raiz do projeto copiando o conteúdo de `.env.example` e ajuste-o conforme necessário. Por padrão, os seguintes valores são utilizados:

```sh
DB_HOST=localhost # host do Redis
DB_PORT=6379 # porta do Redis
DB_Name=0 # banco default
DB_PASSWORD="" # senha default
WEB_SERVER_PORT=8080 # Porta do servidor Web
ID_MAX_REQUEST=5 # Número máximo de requisições por IP
TOKEN_MAX_REQUEST=10 # Número máximo de requisições por token
TIME_WINDOW_MS=10000 # Janela de tempo em milissegundos para retentativa

```

2. Rode o comando

```
docker-compose up
```

-   **Requisição com checagem via token com sucesso:**

```sh
$ curl -H 'API_KEY: some-api-key-123' -vvv http://localhost:8080
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
> API_KEY: some-api-key-123
>
# Response
< HTTP/1.1 200 OK
< X-Ratelimit-Limit: 10
< X-Ratelimit-Reset: 1719666491
< X-Ratelimit-Total: 10
< Date: Sat, 29 Jun 2024 13:08:01 GMT
< Content-Length: 20
< Content-Type: text/plain; charset=utf-8
<
{"message":"hello"}
```

-   **Requisição com checagem via token bloqueado:**

```
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
> API_KEY: some-api-key-123
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 429 Too Many Requests
< X-Ratelimit-Limit: 10
< X-Ratelimit-Reset: 1719666564
< X-Ratelimit-Total: 0
< Date: Sat, 29 Jun 2024 13:09:18 GMT
< Content-Length: 34
< Content-Type: text/plain; charset=utf-8
<
{"message":"rate limit exceeded"}
```

-   **Requisição com checagem via IP com sucesso:**

```
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< X-Ratelimit-Limit: 5
< X-Ratelimit-Reset: 1719666641
< X-Ratelimit-Total: 5
< Date: Sat, 29 Jun 2024 13:10:31 GMT
< Content-Length: 20
< Content-Type: text/plain; charset=utf-8
<
{"message":"hello"}
```

-   **Requisição com checagem via IP bloqueado:**

```
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 429 Too Many Requests
< X-Ratelimit-Limit: 5
< X-Ratelimit-Reset: 1719666659
< X-Ratelimit-Total: 0
< Date: Sat, 29 Jun 2024 13:10:51 GMT
< Content-Length: 34
< Content-Type: text/plain; charset=utf-8
<
{"message":"rate limit exceeded"}
```
