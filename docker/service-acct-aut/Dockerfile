FROM golang:alpine
WORKDIR /app
COPY cmd/service-acct-aut/main.exe .
COPY config/config.yaml .
ENTRYPOINT [ "/main.exe" ]