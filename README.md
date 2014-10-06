Simple Go Client-Server Resource Usage Application On Docker
============================================================

I wrote this as an experimental project to learn two great new technologies (go language and docker). The idea of this project is to run a service in docker container and having a client running in a different docker container requesting information from the service. The go application is a client-server app where server when receives request from client will send its own resource usage information back to client. The client just decodes the response, converts the response to more readable form and will just print it.

How to run on Mac OS X
----------------------

To run docker on Mac OS X we need Boot2Docker.

### Installing boot2docker ###
  - Install Boot2Docker from [here].
  - After installing, from terminal, run `boot2docker init` to initialize boot2docker.
  - Run `boot2docker start` to start boot2docker and export `DOCKER_HOST` as show at the end of command. It usually will be
`export DOCKER_HOST=tcp://192.168.59.103:2375`.


After exporting `DOCKER_HOST` we can run docker commands.
  - Clone this repository `git clone https://prasanth_j@bitbucket.org/prasanth_j/go-docker-client-server-rusage.git go-docker`

### Run go rusage server ###
  - `cd go-docker/server`
  - Run `docker build -t go_rusage_server .` to build docker image from Dockerfile inside server directory. `go_rusage_server` is alias for the newly created container.
  - Run `docker run -i -t -p 8787:8787 go_rusage_server` to run the newly created container. `-i -t` runs container in interactive mode and to kill it via terminal using `ctrl + c`. `-p` option is to help client container reach the server container's internal port that was EXPOSE'd in the Dockerfile.

### Run go rusage client ###
  - Open a new terminal and `export DOCKER_HOST=tcp://192.168.59.103:2375`
  - `cd go-docker/client`
  - Run `docker build -t go_rusage_client .` to builder docker image for client. `go_rusage_client` is alias for the newly create container.
  - Run `docker run -i -t go_rusage_client` to run the newly created container.
  - After running the above command, the client container should exceute the go program and connect to the server container via the host name specified in Dockerfile and port 8787. The server should send response back to the client which should be printed on the terminal like below
```
2014/10/06 03:00:04 Dialing 192.168.59.103:8787
2014/10/06 03:00:04 Connected to 192.168.59.103:8787
2014/10/06 03:00:04 Blank request sent to server
2014/10/06 03:00:04 Read 256 bytes from resource service
2014/10/06 03:00:04 {"Utime":{"Sec":0,"Usec":0},"Stime":{"Sec":0,"Usec":0},"Maxrss":4524,"Ixrss":0,"Idrss":0,"Isrss":0,"Minflt":221,"Majflt":0,"Nswap":0,"Inblock":0,"Oublock":0,"Msgsnd":0,"Msgrcv":0,"Nsignals":0,"Nvcsw":19,"Nivcsw":5}
```

[here]:https://github.com/boot2docker/osx-installer/releases