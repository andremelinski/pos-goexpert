* quando vc tem 2 projetos independentes (2 go.mod) mas que comunicam.
* cmd quer consumir o conteudo do math mas sao 2 go.mod diferentes
* go work init .\math\ ./cmd ou apenas go work --> fica na sua maquina (usa no gitignore)
* mesmo com o go.mod vazio, ele conseguiu sincronizar as 2 pastas, sem sujar as dependencias
* se rodar go run main.go ele da certo
* se rodar go run tidy ele adiciona o workspace fazendo essa liga e baixa qual lib 
