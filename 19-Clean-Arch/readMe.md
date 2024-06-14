Internal: onde fica todas as regras de negocio. "Coracao da implementacao"

### Entity

-   interface.go - > qual metodo do repository vc vai usar para esse usecase. eh a interface que comunica o struct do usecase (CreateOrderUseCase) e o struct do repository (OrdeRepository)
-   order.go -> Cria e valida payload mas nao acessa db

### Use Case

-   Camada para "ter a intencao do usuario"

### Infra

-   parte que nos comunica com o mundo externo
-   Conexao com o Db, server, etc
-   Onde fica o repository

## Event - EventHlanders

-   O que vai acontecer quando uma order for criada.
-   Quando vc der um dispatch no order created ele vai executar o eventHlander que possui o metodo handle, que faz com que o events.EventInterface seja aplicado e por fim, publique a resposta no rabbitmq
