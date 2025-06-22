FROM golang:1.24 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /bin/server ./cmd

FROM build AS run
ENTRYPOINT [ "/bin/server" ]

