-   para gerar um graphql padrao: go run github.com/99designs/gqlgen init;
-   para atualizar os dados do schema.graphqls: go run github.com/99designs/gqlgen generate
-   Resolver: tenta "resolver" um problema. exemplo: quero criar uma nova categoria, o resolver CreateCategory eh executado quando executamos uma mutatation no graphQL pra criar uma categoria

### Para fazer a relacao Courses + Categories, ou Dados encadeados, no Graphql:

<p>Dados encadeados Realizados de maneira explicita</p>
<p>1- Criar 2 arquivos separados com os seus models dentro de graph/model</p>
<p>2- ir em gqlgen.ylm (arquivo base para a criacao do graphql, como resolvers, models, etc) e dentro de models adicional o path dos 2 novos models (courses e categories).</p>
Ao fazer isso, quando vc mandar gerar novamente com o comando go run github.com/99designs/gqlgen generate e um Resolver devera ser gerado. Nesse caso, o resolver devera fazer relacao do model Courses com Categories
