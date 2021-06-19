## Foreign Exchange in DDD + CQRS

An unnecessary complex software architecture application of exposing exchange rates from Indonesian banks.
I'm currently learning DDD and CQRS system. This is just my interpretation in doing so.

While I'm trying to embrace microservice, I want to put the code in monorepo style so I don't have to maintain a lot of repositories.

### Structure

```
- delivery
|-- grpc
|-- gql
|-- worker
- internal
|-- domains
|-- shared
```

Delivery packages will contains of how the app will interact to the outside world. Be it exposing the data via GRPC, GraphQL or a worker/background job.
Each delivery will have it's own handler since it has different responsibility. 

Internal will host domain-specific code and shared code. 
Domain-specific will hold things like entity, repository interface, filterable columns via criteria.
