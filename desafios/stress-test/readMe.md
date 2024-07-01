# Go Expert Challenge - Rate Limiter

Implementação de uma CLI em Go utilizando Cobra para realizar testes de stress em um endereço web.

## Arquitetura

A aplicação é composta por um servidor web que recebe requisições HTTP e um middleware de rate limiter que é responsável por controlar o número de requisições recebidas. O [middleware](internal/infra/web/webserver/middleware/middleware.go) intercepta todas as requisições e realiza toda a chacagem pelo [strategy](internal/infra/web/webserver/middleware/strategy/rate-limiter.go). Para armazenar os dados é utilizado o Redis por ser um banco de fácil uso e que possui a opção de expiração do dado após um determinado tempo.

O rate limiter é configurável por IP ou token `API_KEY` no arquivo `.env`. Exemplo de configuração para limitar 10 requisições por IP e 100 requisições por token em uma janela de tempo de 5 minutos:

## Executando o projeto

**Obs:** é necessário ter o [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

1. Rode o comando

```
docker run melinski/goexpert-stress-test-cli:latest \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

ou utilizando abreviações

<!-- docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname -->

```
docker run melinski/goexpert-stress-test-cli:latest \
    -u=https://google.com.br \
    -r=100 \
    -c=10
```

-   **Test de Stress:**

```sh
$ go run main.go stress -u="http://google.com" -r=5 -c=2
report completed
{
  "url": "http://google.com",
  "total": 5,
  "StatusCode": [
    {
      "200": 5
    }
  ],
  "execution-time-microseconds": 25
},
```
