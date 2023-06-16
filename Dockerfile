FROM golang:1.20 AS builder

RUN apt-get update -y 
RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64

RUN chmod +x /usr/local/bin/dumb-init

WORKDIR /build/bin/
WORKDIR /build

COPY go.mod ./

# Copy the code into the container.
COPY . .

RUN make build

FROM scratch

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/local/bin/dumb-init /usr/bin/dumb-init
COPY --from=builder ["/build/bin/mqttbackend", "/"]

# Export necessary port.
EXPOSE 8044
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
# Command to run when starting the container.
CMD ["./mqttbackend"]