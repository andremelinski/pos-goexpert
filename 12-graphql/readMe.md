-   para gerar um graphql padrao: go run github.com/99designs/gqlgen init;
-   para atualizar os dados do schema.graphqls: go run github.com/99designs/gqlgen generate
-   Resolver: tenta "resolver" um problema. exemplo: quero criar uma nova categoria, o resolver CreateCategory eh executado quando executamos uma mutatation no graphQL pra criar uma categoria

## Para fazer a relacao Courses + Categories, ou Dados encadeados, no Graphql:

<li>Dados encadeados Realizados de maneira explicita</li>
<p> quando vc quer voltar voltar as infos de curso dentro da categoria</p>
<ol>
    <li>Criar 2 arquivos (course e category) com os seus models dentro de graph/model</li>
    <li>ir em gqlgen.ylm (arquivo base para a criacao do graphql, como resolvers, models, etc) e dentro de models adicional o path dos 2 novos models (courses e categories).</li>
    <li>Deletar o bind de Courses em category.go. O GraphQL deve reconhecer ao mandar o comando generate e criarmos mais um resolver para fazermos esse bind</li>
</ol>
Ao fazer isso, quando vc mandar gerar novamente com o comando go run github.com/99designs/gqlgen generate e um Resolver devera ser gerado. Nesse caso, o resolver devera fazer relacao do model Courses com Categories.

```
{
    categorie: {
        id,
        desc,
        name,
        curso: [{
            id,
            desc,
        }]
    }
}
```

<p> Agora, quando vc quer voltar voltar as infos das categorias dentro do curso</p>
<ol>
    <li>Deletar o bind de Category em course.go. O GraphQL deve reconhecer ao mandar o comando generate e criar mais um resolver para fazermos esse bind</li>
</ol>

```
{
    curso: {
        id,
        desc,
        categorie: {
            id
        }
    }
}
```

### coamnados GraphQL executados

```
mutation CreateCategory {
  createCategory(input: {name: "tecnologia 2", description: "desc tec 2"}) {
    id
    name
    description
  }
}

mutation CreateCourse {
  createCourse(
    input: {name: "curso tecnologia 1", description: "curso desc 1 ",
      categoryId: "2ee6e3c1-02fc-48ff-8b55-7ba878b98183"}
  ) {
    id
    name
    description
  }
}

query AllCategories {
  categories {
    id
    name
    description
  }
}

query AllCourses {
  courses {
    id
    name
    description
  }
}


query CategoriesWithCourses {
  categories {
    id
    name
    description,
    courses {
      description,
      name,
      id
    }
  }
}

query QueryCoursesWithCategory{
  courses{
    id,
    description,
    name,
    category{
      id,
      name,
      description
    }
  }
}
```
