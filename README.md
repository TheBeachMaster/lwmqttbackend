# Light-weight MQTT Webhook

A 0-dependency MQTT Webhook.

## Why do I need this?

This became super useful in troubleshooting Authentication(AuthN), Authorization(AuthZ) and Data Ingestion issues I was having with [emqx](https://github.com/emqx/emqx)


## How do I use it?

- Run it; `go run main.go` or 
- Build it; `make build` for Linux(x86),  `make build-mac` for Intel-based Macs or `make build-mx_mac` for M-series Macs. 
    - Then `make run` 

> You can also make live changes (if not using Docker) by launching [air](https://github.com/cosmtrek/air)  within this directory.
 
### Have you run into some weird null ptrs??? Oh no!!  🫢😱 

> Run `go run main.go` and pipe to [panic parse](https://github.com/maruel/panicparse) like so 
```sh 
 go run main.go &| panicparse
```

Server is running on port `8044`.

Now you can configure your Auth rules as needed:
- Authz `/authz` 
- Authn `/authn` 
- Data sink `/sink` 

# Disclaimer
- Make sure you are familiar with the License to which this software is distributed under
- Maybe do'nt run this in production, it's for troubleshooting only.
