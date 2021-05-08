# desafio-go
**Desafio Star Wars**  :star:

Tecnologias:
Golang
Docker

Rotas das API's:

:large_blue_diamond: Post:
{url}/planets
body: {
  "name": "Tatooine",
  "climate": "arid",
  "terrain": "desert"
}

Response:
{
  "id": "6095e489afab8abc77b0cf57",
  "name": "Tatooine",
  "climate": "arid",
  "terrain": "desert",
  "movies": 5
}

:large_blue_diamond: Get:
{url}/planets/  -> Retorna todos os planetas

{url}/planets/:id  -> Retorna o planeta com o id específico

{url}/planets?name={"name":"Tatooine"}  -> Retorna todos os planetas dado um filtro. No exemplo é o filtro é "name":"Tatooine", além do "name" o filtro poderia ser pelos atributos "climate" e "terrain".

:large_blue_diamond: Delete
{url}/planets/:id  -> Remove o planeta com o id específico.



