-   para gerar um graphql padrao: go run github.com/99designs/gqlgen init;
-   para atualizar os dados do schema.graphqls: go run github.com/99designs/gqlgen generate
-   Resolver: tenta "resolver" um problema. exemplo: quero criar uma nova categoria, o resolver CreateCategory eh executado quando executamos uma mutatation no graphQL pra criar uma categoria

### Para fazer a relacao Courses + Categories, ou Dados encadeados, no Graphql:

<li>Dados encadeados Realizados de maneira explicita</li>
<ol>
    <li>Criar 2 arquivos (course e category) com os seus models dentro de graph/model</li>
    <li>ir em gqlgen.ylm (arquivo base para a criacao do graphql, como resolvers, models, etc) e dentro de models adicional o path dos 2 novos models (courses e categories).</li>
    <li>Deletar o bind de Courses em category.go. O GraphQL deve reconhecer ao mandar o comando generate e criarmos mais um resolver para fazermos esse bind</li>
</ol>
Ao fazer isso, quando vc mandar gerar novamente com o comando go run github.com/99designs/gqlgen generate e um Resolver devera ser gerado. Nesse caso, o resolver devera fazer relacao do model Courses com Categories
