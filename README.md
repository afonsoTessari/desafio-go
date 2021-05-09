
:star:  **Desafio Star Wars**  :star:

Tecnologias:
Golang e Docker

Rodar o projeto através do comando:  docker-compose up --build

*ulr = localhost:8080*


Rotas das API's:

:large_blue_diamond: Post:

:heavy_minus_sign: {url}/planets

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

:heavy_minus_sign: {url}/planets/  -> Retorna todos os planetas

:heavy_minus_sign: {url}/planets/:id  -> Retorna o planeta com o id específico

:heavy_minus_sign: {url}/planets?name={"name":"Tatooine"}  -> Retorna todos os planetas dado um filtro.

:large_blue_diamond: Delete

:heavy_minus_sign: {url}/planets/:id  -> Remove o planeta com o id específico.



