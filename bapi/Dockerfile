FROM golang:1.15 AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/kubernetesProject/bapi

# Copy the current code into our workdir
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/src/kubernetesProject/bapi/bapi

# final stage
FROM ubuntu:bionic
RUN apt update && apt install  ca-certificates -y && apt install -y openssl && apt install ssh -y

WORKDIR /app
COPY --from=build-env /go/src/kubernetesProject/bapi/swagger/ /app/swagger/
#COPY --from=build-env /go/src/kubernetesProject/bapi/swagger-master/ /app/swagger-master
COPY --from=build-env /go/src/kubernetesProject/bapi/bapi /app/

EXPOSE 8080

CMD ./bapi