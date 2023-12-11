# workspace (GOPATH) configured at /go
FROM golang:1.21.5 as builder

WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . ./

RUN make build

FROM alpine:latest

COPY --from=builder app/bin/server ./server

RUN chmod +x ./server 

CMD ["./server"]
