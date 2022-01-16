# Placeholder

Simple golang http-rest app used to test pipelines, environments or something like this. 
This app is simple enough that whether something is wrong, it's not this app ;)

## Running

Run:
```bash
go run app/main.go
```

## Test

Run:
```bash
go test -v -cover -covermode=atomic ./...
```

## APIs
### Poems

API that store and fetch [Poems](domain/poem.go) in the format:
```go
type Poem struct {
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}
```
A [postman-collection](placeholder.postman_collection.json) is provide for simplify its usage.


#### GET /poem

Return a list of Poems.
Example:
```json
[
    {
        "content": "Nem alegre nem triste, poeta",
        "author": "Someone"
    },
    {
        "content": "Eu canto porque o instante existe\ne a minha vida está completa.\nNão sou alegre nem sou\ntriste:\nsou poeta.",
        "author": "Cecília Meireles"
    }
]
```

#### GET /poem/static

Return a static Poem.
Example:
```json
{
    "content": "Nem alegre nem triste, poeta",
    "author": "Someone"
}
```

#### GET /poem/random

Return a random Poem from the current collection.
Example:
```json
{
    "content": "Eu canto porque o instante existe\ne a minha vida está completa.\nNão sou alegre nem sou\ntriste:\nsou poeta.",
    "author": "Cecília Meireles"
}
```

#### POST /poem

Add a new Poem to the collection. The body must be a json like this:
```json
{
    "content": "Toda a poesia - e a canção é uma poesia ajudada - reflete o que a alma não tem.\nPor isso a canção dos povos tristes é alegre e a canção dos povos alegres é triste.",
    "author": "Fernando Pessoa"
}
```
