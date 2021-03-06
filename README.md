# License - Manager
This is a go microservice featuring a REST controller and a connection to RabbitMQ (Amqp)

* by default it attempts to connect to the following RabbitMQ address: amqp://guest:guest@localhost:5672/ (You can override this via a command line switch `-amqp`)

## Quick Start

There is a prebuilt image avialable from dockerhub

If you wish to run with Challenge 3 stuff run this 

`docker run -d -p 5672:5672 -p 15672:15672 -it rabbitmq:3.7-management-alpine`

before you run the license-mamager

`docker run -p 8080:8080 -p 9090:9090 -it sevren/license-manager`

Or you can run the code directly (if you have go installed)

`go run .`

Once it starts you have a REST controller listening on http://localhost:8080

## BONUS!
There is a swaggerui endpoint being served on http://localhost:9090/swaggerui

You can make a request to the /{user}/licenses endpoint

## Regarding Challenge 3 

This code also contains code to connect to the rabbitMQ server. It will connect to the exchange `data` and consume messages. 

These messages will then get stored in the database table

In this challenge for license-manager The licenses are changed from base64 to hashids using a salt of the user:license

If you want to checkout challenge 3 stuff, please start the rabbitmq server with the following 

`docker run -d -p 5672:5672 -p 15672:15672 -it rabbitmq:3.7-management-alpine`

## REST endpoints

The REST controller listens on port 8080

The REST endpoint is served on http://localhost:8080/

* POST /{user}
* POST /{user}/licenses
* GET /usedlicenses

The endpoints will attempt to autenticate the user from the database as the request comes through. 

With the correct password the user will get the licenese they have attached to thier account.

## Regarding the licenses endpoint

If the microservice is running with out a connection to rabbitMQ the licenses will be generated with the challenge v1 base64 code. 

If the microservice is running with the connection to rabbitMQ then the licenses will be generated the challenge 3  hashid code using a salt of the username:license

## Regarding the used licenses endpoint

You can make a GET request toward the database to check the used licenses

`curl http://localhost:8080/usedlicenses`

Produces output similar to the following if there is licenses

```json
{
    "licenses": [
        "B5K4A6Q0yJKKcQYULahQ83nMZeVo8N",
        "N2YPE6lO9J11IaYiOxIJv31aL5WzKg"
    ]
}
```

and the following if there is no licenses

```json
{
    "licenses": []
}
```

## Swagger ui

There is a swaggerui endpoint available at the following address once the microservice is started up

http://localhost:9090/swaggerui

## Building
To build from source you require the following: 
* Go (1.12)
* Make
* Docker (With the ability to run without sudo..)

`make static`

### Docker Container
The following will create a local docker image with the dockertag set

`make DOCKERTAG=license-man docker`

You can view your images by 
`docker image`

## Running

To run locally you can use the following. 

If you wish to test the service to service communication, you need to run a rabbit mq server.

The following will start one in a docker container.

*OBS!: RabbitMQ is notorious for taking it's time to start up, please give the container a minute or so to be fully booted*

`docker run -d -p 5672:5672 -p 15672:15672 -it rabbitmq:3.7-management-alpine`

To run (with default AMQP) the microservice from code simply:

`go run .`

### Overriding the RabbitMQ url

You can override the rabbitmq url by giving the command line flag `-amqp <RMQ URL>`
The format of the <RMQ URL> is "amqp://{user}:{password}@<host>:<port>/"

### Running with a docker container

If you have build this image locally: 

`docker run --network=host -it license-man -amqp amqp://guest:guest@localhost:5672/`

If you wish to just use my already published docker image from dockerhub: 

`docker run --network=host -it sevren/license-manager`