# SPAC - Week 9 - Frontend-Integration
An exercise in integrating a frontend with an API.

## Description:
For this project i have chosen to develop a simple full stack application serving quotes from movie with data about the movies, actors, movie characters, etc.  

The backend is development using [Go](https://go.dev/) with PostgreSQL as the database for persisting data.

The frontend is development using TypeScript, React and Vite as a Single Page Application

Because the exercises focuses on frontend/backend integration I have chosen to use [Protocol Buffers](https://protobuf.dev/) and [gRPC](https://grpc.io/docs/what-is-grpc/introduction/) to function as a glue between frontend and backend, where communication and messages between the two are clearly defined and followed.  
By doing this I ensure that client and server can communicate seamlessly as well as ensuring type safety on both ends.

## Disclaimer
This project is in no way meant to be production ready or deployed. It is merely a development exercise and a showcase of how I would would solve it.

## Contents
a mono-repo containing multiple elements of the overall application. Those elements include:
- __Backend: Go gRPC Server__ 
   - Written using __Go__ using [__Entgo ORM__](https://entgo.io/docs/getting-started/)
- __Frontend: React SPA__ 
   - Written using __React, TypeScript, Vite__.
   - Serving a public facing website where users can get a random movie quote as well as view all movie quotes.  
   - Other tools of note: __@tanstack/react-query, shadcn/ui__
- __Envoy Proxy__
   - Serves as a Proxy between frontend and backend bridging the gap between the server using HTTP2 and the client using HTTP1.
- __Database: PostgreSQL__
   - For persisting data used by the application.

## Development Dependencies
- [Docker](https://www.docker.com/get-started/)
- [Go v1.24.2](https://go.dev/)
- [Bun v1.2.4](https://bun.sh/)

## Try It Out
1. Ensure docker is installed and that the docker daemon is running.  
*Getting the daemon running depends on the your environment. On Windows machines it's usually done by starting the Docker Desktop application*
2. Start the application by running:
```sh
docker compose up
```
