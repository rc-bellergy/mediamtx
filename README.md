# MediaMTX support multiple authJWTJWKS

Modified the mediamtx to support multiple authJWTJWKS.
## Config multiple authJWTJWKS in mediamtx.yml
```yml
    paths:
        "~^ghl/ghc/cam_([0-9]+)$":
            sourceOnDemand: yes
            source: {{ .RTSP_URL }}$G1
            authJWTJWKS: {{ .KEYCLOAK_URL }}goldlion-ghc/protocol/openid-connect/certs
```

## Generate mediamtx.yml
### Create a .env file with the following content:

    KEYCLOAK_URL="https://keycloak_server/realm/"
    RTSP_URL="rtsp://url_of_rtsp_server/"
    SSL_URL="/path_to_ssl_cert/"

### Generate
    go run scripts/generate_config.go

## Build
    go generate ./...
    CGO_ENABLED=0 go build .
    ./mediamtx

## Dev
    go generate ./...
    go run .

