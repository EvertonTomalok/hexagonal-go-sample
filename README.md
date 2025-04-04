# ports-challenge

This repository contains an application that is able to parse a huge JSON file, parse this file node by node and upsert it in a Memory Database (a map was used here as the repository).

# Architecture

This application follows the Hexagonal Architecture to structure and organize the code efficiently:

- Domain (Entities) – Defines the core business logic, remaining independent of frameworks, databases, or external APIs.

- Application (Use Cases) – Orchestrates business logic, ensuring separation between application workflows and infrastructure/services.

- Ports (Interfaces) – Serve as the entry and exit points to the core logic, using interfaces to abstract external dependencies.

- Adapters (Implementations of Ports) – Bridge the domain with real-world assets by implementing the necessary interfaces.

Additionally, the project includes instructions for the repository and service layers.


# How to execute

This project has a dockerfile containing a multi stage file that allows to run and test the project.

First, run: `database-up` to setup postgresDB, and then `make setup-dev && migrate-up` to run migrations
PS: run `database-down` to remove postgresDB

Running the application:
- In your machine: `make run-local`
- In docker: `make up` or `docker-compose up app`

Testing:
- `docker run $(docker build --target testing -q .)`

# Dev helper

This project also contains a Makefile with some receipts to help developers

## Receipts:
- `make setup-dev`: will install dev dependencies
- `make lint`: will run lint locally
- `make format`: will run formatter locally
- `make unit-test`: will run unit tests locally 
- `make mock-gen`: will generate/update mocks

# Post notes and possible enhancements

## JSON decoder

I would try to find a better solution when parsing the file, like doing it in chunks to speed up the reading process, even considering that this approach may increase memory consumption. However, the trade-off between performance and resource usage would need to be carefully analyzed.

## Optimize Upsert Performance Using a Worker Pool 

Given that input items arrives in chunks, I would implement a worker pool pattern to process the upsert stream more efficiently. This approach would enable concurrent processing, reducing latency and improving throughput. Additionally, batching inserts/updates in an SQL database could further enhance performance by minimizing transaction overhead.

## Choosing a database

Using an in-memory map is not the best approach for a production environment. Here are some alternatives:

- A memory-based database like Redis could be a great first choice for storing key-value pairs, providing fast and efficient read access.
- If persistence and optimized read performance become important in the future, implementing a pattern like CQRS (Command Query Responsibility Segregation) would be beneficial. In this case, I would use a database like MongoDB (since it is a document based) for writes, Elasticsearch as a read-optimized database, and Kafka as an event bus to ensure data written to MongoDB is replicated in Elasticsearch.
