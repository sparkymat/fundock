# fundock

fundock is a simple FaaS which run functions using Docker containers.

## Installation

The quickest way to get started with fundock is by using docker-compose. Copy the `docker-compose.yml` file to a folder, and run `docker-compose up` to run fundock. You can then point your browser to http://localhost:8080 to use it.

## Features

`fundock` is a simple FaaS implementation that runs functions using Docker containers. Any docker container can potentially be used. `fundock` passes in the provided input to the `stdin` of the container, and reads any logs/output back and stores it as the output of the function.

https://raw.githubusercontent.com/sparkymat/fundock/feature/docker-release/docs/invocation.mp4
