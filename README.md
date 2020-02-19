# Knowledge base

This is a server application that exposes API to store and maintain the documents
in the knowledge base.

## Prerequisites

### To run

- Docker
- docker-compose

### To develop

- golang,
- mongodb
- make
- goraml

## How to  

## Compile application

```
make
```

## Run in docker-compose

```
docker-compose up
```

## Build docker

```
docker build -t knowledge-base .
```

## API endpoints

API serves CRUD like functionality on the following endpoints

- GET /documents - returns the list of documents id  
- GET /document/{docimentID} - returns the list of versions of the document
- GET /document/{docimentID}/{versionHash} - returns the document description
- POST GET /document - creates the document with new id
- POST /document/{docimentID} - creates new version of the document
- PUT /document/{docimentID}/{versionHash} - updates specific version of the document
- DELETE /document/{docimentID} - delete specific document id
- DELETE /document/{docimentID}/{versionHash} - delete specific version of the document

more on the API documentation [here](kb.raml)

## cURLs

Some cURL commands on the local host implementation

```
$ curl 127.0.0.1:5000/documents         
{"items":["sdff"]}
```

```
$ curl -X POST -d '{"document_id":"dummy","document_url":"super/resource","elements":[{"text":"super interesting text","type":"p"},{"text":"super interesting text2","type":"p"},{"text":"super interesting heading","type":"h1"}]}' 127.0.0.1:5000/document
```

```
$ curl -X PUT -d '{"document_id":"dummy","document_url":"super/resource-updated","elements":[{"text":"super interesting text","type":"p"},{"text":"super interesting text2","type":"p"},{"text":"super interesting heading","type":"h1"}]}' 127.0.0.1:5000/document/68be8af0-4a08-4e2b-8ec4-6e7b8636113d/de9ce5571694ae4ed56b73e0e265b73a
```

```
$ curl 127.0.0.1:5000/document/68be8af0-4a08-4e2b-8ec4-6e7b8636113d/829cf053d8d5c972c5caed354c000f62
{"document_id":"68be8af0-4a08-4e2b-8ec4-6e7b8636113d","document_url":"super/resource-updated","elements":[{"text":"super interesting text","type":"p"},{"text":"super interesting text2","type":"p"},{"text":"super interesting heading","type":"h1"}],"etag":"829cf053d8d5c972c5caed354c000f62","version":{"created":0,"hash":"829cf053d8d5c972c5caed354c000f62","updated":1582084366}}
```

```
curl -X DELETE 127.0.0.1:5000/document/9432e4b8-e315-45ba-a84c-eac6da576a2b/0be81dd93cfb4d414fbf8a152d7f9b69
```

```
curl -X POST -d '{"document_id":"dummy","document_url":"super/resource","elements":[{"text":"super interesting text","type":"p"},{"text":"super interesting text2","type":"p"},{"text":"super interesting heading","type":"h1"}]}' 127.0.0.1:5000/document
{"document_id":"dummy","document_url":"super/resource","elements":[{"text":"super interesting text","type":"p"},{"text":"super interesting text2","type":"p"},{"text":"super interesting heading","type":"h1"}],"etag":"","version":{"created":0,"hash":"","updated":0}}
```

## Implementation limitation

- Assume that there is parser before knowledge-base
- Assume that parer obeys the following model shown by example:

```
{
  "document_id": "dummy",
  "document_url": "super/resource",
  "elements": [
    {
      "text": "super interesting text",
      "type": "p"
    },
    {
      "text": "super interesting text2",
      "type": "p"
    },
    {
      "text": "super interesting heading",
      "type": "h1"
    }
  ]
}
```

## What could be done more

- Indexing on the DB
- Pagination
- Some edge cases are not covered
- Independent error messages
- Interactive documentation
- Postman API templates  
- GH actions CI
- Dockerhub pushes
- HTML to JSON parser
- strict validation of the JSON content 
