FROM golang:1.20

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /aps

## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /aps

COPY go.mod /aps
COPY go.sum /aps
RUN go mod download

## We copy everything in the root directory
## into our /app directory
ADD . /aps

## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .


## Our start command which kicks off
## our newly created binary executable
CMD ["/aps/main"]