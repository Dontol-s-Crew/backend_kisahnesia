FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./cmd .
COPY ./cmd/. ./cmd/
COPY ./image .
COPY ./image/. ./image/
COPY ./internal .
COPY ./internal/. ./internal/
COPY ./pkg .
COPY ./pkg/. ./pkg/


RUN go mod download && go mod verify
RUN go build -o bin ./cmd/main/.

EXPOSE 8000

ENTRYPOINT ["/app/bin"]