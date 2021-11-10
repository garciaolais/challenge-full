FROM golang:alpine AS build
WORKDIR /go/bin
COPY . .
RUN apk add --no-cache make
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/rs/cors
RUN go get -d github.com/garciaolais/challenge-backend
RUN make build
CMD ["go","run","cmd/main.go"]