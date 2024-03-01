-   go test -v => gera uma visao geral dos testes
-   go test -coverprofile=coverage.out => Gera o coverage com fails e gera um aiquivo chamado coverage.out
-   go tool cover -html=coverage.out => gera um coverage html a partir do arquivo .out gerado pelo comando acima

## BenchMark testa a funcao e retorna algumas infos alem do SO e pacote:

    -   Nome da funcao que esta rodando-qtde de numeros computacionais usados
    -   Quantidade de operacoes que ele conseguiu executar na funcao
    -   Tempo gasto por cada operacao

-   go test -bench=. -run=^# => se quiser rodar apenas os benchmarks. OBs: no -run= eh uma expressao regular pra vc moldar quais funcoes vc quer rodar (usado para rodas apenas as bench ou comparar)
-   go test -bench=. -run=^# -count=10 => usado para rodar 10 vezes a operacao
-   go test -bench=. -run=^# -count=10 -beanchtime=3s => bench executa com timeout setado em 3s
-   go test -bench=. -run=^# -beanchmen => mostra alocacao de memoria

## Fuzz testa aleatoriamente as coisas ate quebrar a aplicacao

-   go test -fuzz=. -run=^# => tenta um monte de numero ate quebrar a aplicacao e gera a pasta fuzz. para tentar de novo o msm valor que quebrou a aplicacao go test -run=FuzzCalculateTax/idDoFile
-   go test -fuzz=. -fuzztime=5s -run=^# => tenta quebrar a aplicacao por 5s
