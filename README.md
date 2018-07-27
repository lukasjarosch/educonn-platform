# educonn.de
This project was developed as part of my master thesis in computer science at the university of applied sciences Ravensburg-Weingarten.
Lukas Jarosch - 2018

## Makefile commands
```make all``` builds the binaries of all services
```make <service>``` builds the binary of a service
```make proto``` builds the protobufs of all services
```make <service>-proto``` build the protobuf of a service
```make docker``` build and tag the docker images of all services
```make <service>-docker``` build and tag the docker image of a service:w

*Note:* The version of the tags are currently formatted like: **derwaldemar/educonn-<service>:MAJOR.MINOR-[commit-hash]**. As well as all images are also tagged with **dev**

## Run the stack
Well, currently my ```docker-compose``` does contain secrets so I cannot share it here until I sorted that out.