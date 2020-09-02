Exemplo Básico Para Processo de CI/CD Deploy Contínuo

Docker HUB [go-hpa](https://hub.docker.com/repository/docker/brmaschio/go-hpa)

+ Testes
+ Processo de CI
+ Deploy da imagem no K8S (deployment e services)
+ Cada réplica com consumo no mínimo 50m e no máximo 100m
+ Processo de escala inicia quando a CPU passar de 15%
+ Quantidade mínima de pods: 1
+ Quantidade máxima de pods: 6


* Quando qualquer push ou uma PR for relizada no Github em um branch diferente do Master, o processo de CI é executado.
* Quando um merge ou um push entrarem no branch Master, o processo de CI/CD é chamado, fazendo assim o deploy de forma automática no Kubernetes.
