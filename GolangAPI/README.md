#Simple GO CRUD app

##Requeriments

    go get github.com/gorilla/mux

##Usage

Compile with 
    go build main.go
then execute with the generated binary 
*Or* with
    go run main.go


##Routes

Método | Rota | Função
-------|------|---------
GET | /pessoas | Retorna todos os registros
GET | /pessoas/{id} | Retorna a pessoa identificada pelo {id}
POST | /pessoas/{id} | Lê um JSON e registra uma pessoa identificada pelo {id}
DELETE | /pessoas/{id} | Deleta a pessoa identificada pelo {id}

