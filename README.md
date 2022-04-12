![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

## **Sobre o projeto**
É uma solução para simularmos transferências de valores entre bancos fictícios através de chaves (email, cpf).

É simulado diversos bancos e contas bancárias que possuem uma chave Pix atribuída

Cada conta bancária poderá cadastrar suas chaves Pix

Uma conta bancária poderá realizar uma transferência para outra conta em outro banco utilizando a chave PIX da conta de destino

Uma transação não pode ser perdida mesmo que: o Codepix esteja fora do ar

Uma transação não pode ser perdida mesmo que o: o Banco de destino esteja fora do ar

## **Tecnologias utilizadas**
- Bank API - Nest.js
- Bank Frontend - Next.js
- Code Pix (Microsserviço) - Golang
- Comunicação entre os microserviços - gRPC
- Comunicação assíncronas - Kafka
- Docker
- Kubernets