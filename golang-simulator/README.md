# Imersão Fc 20 - Sistema de rastreamento de veículos

## Descrição

Repositório do Golang (simulador de trajeto e cálculo de frete).

## Rodar a aplicação

Levante os containers:

```
docker compose up -d
```

Entre no container simulator:

```
docker compose exec simulator bash
```

Rode o comando para executar o projeto:

```
go run cmd/simulator/main.go
```

