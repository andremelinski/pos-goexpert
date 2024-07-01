Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado. --requests: Número total de requests. --concurrency: Número de chamadas simultâneas.

Execução do Teste:

Realizar requests HTTP para a URL especificada. Distribuir os requests de acordo com o nível de concorrência definido. Garantir que o número total de requests seja cumprido. Geração de Relatório:

Apresentar um relatório ao final dos testes contendo: Tempo total gasto na execução Quantidade total de requests realizados. Quantidade de requests com status HTTP 200. Distribuição de outros códigos de status HTTP (como 404, 500, etc.). Execução da aplicação: Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex: docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10
