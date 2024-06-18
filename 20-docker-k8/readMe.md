```
# toda vez que acessar a porta 8080, vai redirecionar pro 8080 do service, que redireciona pro 8080 de um pod onde ele tem um app server. Como possui 3 replicar, ele faz o load balancer chamando os 3
kubectl port-foward svc/serversvc 8080:8080
```

-   Probes -> verificacoes se container ja subiu, pode receber req e se esta no ar
-   startupProbe -> verifica se o container esta pront para receber req. Manda metodo Get para um path por a por x vezes a cada t segundos (default eh um /health). ISSO SOH OCORRE NA 1 SUBIDA do pod
-   readinessProbe -> verifica se a aplicacao esta pronta. Faz com que o service pare de mandar trafego para aquele pod. fica vendo o tempo inteiro e nao soh na subida mas funciona no msm esquema que o startupProbe
-   livenessProbe -> parecido com o readinessProbe, mas verifica se ele esta de pe, mas se ele cair algumas vezes, ele vai tentar recriar o pod com um restart
