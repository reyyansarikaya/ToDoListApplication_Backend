FROM golang:alpine AS Builder
ARG ENV
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
EXPOSE 3000
CMD [".\main.go"]
