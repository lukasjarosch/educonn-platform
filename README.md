# educonn.de
This project was developed as part of my master thesis in computer science at the university of applied sciences Ravensburg-Weingarten.
The goal was to develop a microservice architechture for a video-streaming and e-learning platform.

Lukas Jarosch - 2018

## Makefile
The Makefile provides all the tools you need to get up and running.
For every service, there exists a set of makefile commands:
 + `make <service>` Build the service locally
 + `make <service>-run` -- Run the service locally
 + `make <service>-docker`-- Build the docker image and tag it with 'staging-latest' 
 + `make <service>-publish`-- Upload the images to DockerHub
 
## Services
Currently, the following services are implemented. For detailed information about the services, check the service Readme. 
+ **user**
+ **user-api**
+ **video**
+ **video-api**
+ **mail**
+ **lesson**

Services currently in dev are **lesson-api** and **user-web**

## How to run the stack?
Well, currently my ```docker-compose``` does contain secrets so I cannot share it here until I sorted that out.


## Development logging
For development purposes the JSON logs are rather hard to read. 
To enable pretty logging, just `export DEV_ENV=True` before starting a service.