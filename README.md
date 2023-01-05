# fundock

fundock is a simple FaaS which run functions using Docker containers.

## Installation

The quickest way to get started with fundock is by using docker-compose. Copy the `docker-compose.yml` file to a folder, and run `docker-compose up` to run fundock. You can then point your browser to http://localhost:8080 to use it.

## Overview

`fundock` is a simple FaaS implementation that runs functions using Docker containers. Any docker container can potentially be used. `fundock` passes in the provided input to the `stdin` of the container, and reads any logs/output back and stores it as the output of the function.

https://user-images.githubusercontent.com/104314/209816472-be790e41-11ce-4629-bc53-d59833d0ff9e.mp4

## Functions

Some useful functions can be found at https://github.com/sparkymat/fundock-functions

## Auth

Currently, the web app provides limited authentication options. The current options are:

### Single user mode

This is ideal for small setups where the web app does not need to be protected. It automatically logs the user in as an `admin` user. A user with the username `admin` is automatically injected into the database. To enable this mode, the user needs to set the `SINGLE_USER` environment variable to `1`.

## API

`fundock` provides an API which is documented [here](API.md).
