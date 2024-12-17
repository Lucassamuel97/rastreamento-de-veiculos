## Imersão FC20 - Sistema de rastreamento de veículos

## Descrição

Repositório do NestJS: API responsável pelo salvamento das rotas no banco de dados (MongoDB), comunicando-se com a API do Google Maps para obter dados de locais e rotas.


## Requerimentos

É necessário integrar com Google Maps, para isso, é necessário criar uma chave de API no Google Cloud Platform. Para isso, siga os passos:

1. Acesse o [Google Cloud Platform](https://cloud.google.com/).
2. Crie um novo projeto.
3. Ative a API do Places e Directions do Google Maps.
4. Guarde a chave de API gerada.

## Rodar a aplicação

Gere o arquivo `.env` através do comando:

```
cp .env.example .env
```

Coloque a chave de API gerada no Google Cloud Platform no arquivo `.env`.

Levante os containers do Nest.js e MongoDB:

```bash
docker-compose up -d
```

Entre no container do Nest.js e instale as dependências:

```bash
docker-compose exec app bash
npm install
```

gerar o *Client do Prisma*:
```bash
npx prisma generate
```

Rode o servidor:

```bash
npm run start:dev
```

As especificações das chamadas HTTP estão no arquivo `api.http`. Você pode usar a extensão `REST Client` do VSCode para fazer as chamadas.