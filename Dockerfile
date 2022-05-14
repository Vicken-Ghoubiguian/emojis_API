#Put the golang image as image's base
FROM golang:latest

#
LABEL maintainer="ericghoubiguian@live.fr"

#Copy all the files and directories in the newly created directory emojis_GoAPI
COPY . /emojis_GoAPI

#Change work directory for the emojis_GoAPI project one
WORKDIR /emojis_GoAPI

#Edit the environment variable value GOPATH for the emojis_GoAPI directory
ENV GOPATH /emojis_GoAPI

#Expose the docker container listening port
EXPOSE 80

#Container instruction as entrypoint: 'go run main.go'
ENTRYPOINT ["go", "run", "main.go"]
