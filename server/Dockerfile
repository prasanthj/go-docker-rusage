##############################################################
# Dockerfile to run server golang resource usage application #
##############################################################

# Set the base image to ubuntu
FROM ubuntu

# Author
MAINTAINER Prasanth Jayachandran

# Update the repository sources list
RUN apt-get update

# Install golang
RUN apt-get install -y golang

# Add source file to server directory
ADD rusage_server.go /var/server/rusage_server.go

# Expose the port that server listens to
EXPOSE 8787

CMD ["go", "run", "/var/server/rusage_server.go"]