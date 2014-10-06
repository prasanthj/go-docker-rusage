##############################################################
# Dockerfile to run client golang resource usage application #
##############################################################

# Set the base image to ubuntu
FROM ubuntu

# Author
MAINTAINER Prasanth Jayachandran

# Update the repository sources list
RUN apt-get update

# Install golang
RUN apt-get install -y golang

# Add source file to client directory
ADD rusage_client.go /var/client/rusage_client.go

# Command to run on container startup. boot2docker is used to run docker on Mac OS X.
# 192.168.59.103 is the ip address for the host that runs server (default ip address
# that boot2docker asks us to export).
CMD ["go", "run", "/var/client/rusage_client.go", "192.168.59.103"]
