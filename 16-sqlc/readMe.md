```
migrations
migrate create -ext=sql -dir=sql/migrations -seq init
docker-compose exec mysql bash
```

se vc colocar no query.sql -- name: CreateCategory :execresult vai voltar o error e o result da query
