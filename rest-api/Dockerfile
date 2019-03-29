FROM golang:1.8
# Create app directory
WORKDIR /usr/src/app
COPY . .
RUN go get github.com/gorilla/mux
RUN go build
#CMD ["go", "run", "app.go"]