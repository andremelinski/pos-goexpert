UOW -> unit of work

-   como funciona:
-   BEGIN (transactinos comecam)
-   transaction 1 -> repo 1
-   transaction 2 -> repo 2
-   COMMIT / ROLLBACK

-   > Repositoris -> UOW -> getRepository -> Repositorio(TX) "volta a conexsao com o banco como uma tx" - > Registar repos para inicializar com otx
-   > getRepo
-   > UnregisterRepo
-   > DO: metodo que vai receber toda a transacao

```
DO( func(uow uow)error ) error{
    -   BEGIN (transactinos comecam)
-   transaction 1 -> repo 1
-   transaction 2 -> repo 2
-   COMMIT / ROLLBACK
}
```

https://www.linkedin.com/pulse/what-repository-pattern-alper-sara%C3%A7/
