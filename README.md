# fire

Fire is a Go framework for writing a FHIR HTTP REST API server fully compatible with the spec.

If you need FHIR resources definitions, use [fhires](https://github.com/orsinium-labs/fhires).

## Installation

```bash
go get github.com/orsinium-labs/fire
```

## Usage

```go
package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/orsinium-labs/fhir-server/fire"
  "github.com/orsinium-labs/fhires/fhires"
)

func main() {
  resources := fire.Registry{
    fire.TypeEncounter: fire.Operations[fhires.Encounter]{
      Read: func(r fire.ReadRequest[fhires.Encounter]) fire.Response {
        if r.ID == "123" {
          return r.Ok(fhires.Encounter{
            ID:     new("123"),
            Status: fhires.EncounterStatusCompleted,
          })
        }
        return r.BadRequest(fhires.OperationOutcome{
          Text: &fhires.Narrative{
            Div: "Encounter not found",
          },
        })
      },
    },
  }
  fmt.Println("listening")
  err := http.ListenAndServe(":8080", resources.Mux())
  if err != nil {
    log.Fatal(err)
  }
}
```

Run:

```bash
go run .
```

Send request:

```bash
curl 'http://127.0.0.1:8080/Encounter/123?_format=json'
# {"id":"123","status":"completed","resourceType":"Encounter"}
```
