# Simple Nodejs CRUD app

## Requeriments

    npm install express --save
    npm install body-parser --save

## Usage

    npm install
Then run with 
    nodejs index.js


## Routes

Método | Rota | Função
-------|------|---------
GET | /pessoas | Retorna todos os registros
GET | /pessoas/{id} | Retorna a pessoa identificada pelo {id}
POST | /pessoas/{id} | Lê um JSON e registra uma pessoa identificada pelo {id}
DELETE | /pessoas/{id} | Deleta a pessoa identificada pelo {id}



