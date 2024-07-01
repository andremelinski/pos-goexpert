# Go Expert Challenge - Rate Limiter

Implementação de uma CLI em Go utilizando Cobra para realizar testes de stress em um endereço web.

## Arquitetura

As requisições HTTP são realizadas utilizando green threads de acordo com o numero colocado na opção de concurrency, sendo seu retorno armazenado em um canal para futuro tratamento do status code de cada chamada como demonstrado na funcao [stress](internal/usecases/stress.go). Ao fazer o pull da informação contida no canal, é realizado o agrupamento de status code em um array e posteriormente escrito um arquivo data.txt. Para obter essas infromações, o arquivo é lido e retornado na cli

## Executando o projeto

**Obs:** é necessário ter o [Docker](https://www.docker.com/) instalado.

1. Rode o comando para fazer o pull da imagem no DokcerHub

```
docker pull melinski/goexpert-stress-test-cli:latest
```

2. Rode comando para testar

```
docker run melinski/goexpert-stress-test-cli \
    stress \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

ou utilizando abreviações

```
docker run melinski/goexpert-stress-test-cli \
    stress \
    -u=http://google.com.br \
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
