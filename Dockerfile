#Stage 1 - Install dependencies and build
FROM --platform=linux/amd64 golang:1.20.4-alpine as builder

WORKDIR /app

ENV DOCKER_DEFAULT_PLATFORM=linux/amd64

COPY . .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o gogolook main.go


# Stage 2 - Create the run-time image
FROM --platform=linux/amd64 scratch

ENV DOCKER_DEFAULT_PLATFORM=linux/amd64
ENV GIN_MODE=release

WORKDIR /server

COPY --from=builder /app/gogolook ./

EXPOSE 8080

CMD ["./gogolook"]


