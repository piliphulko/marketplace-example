FROM golang:alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY internal/pkg ./
COPY internal/service/service-acct-auth ./
COPY api/basic ./
COPY cmd/service-acct-auth/*.go ./
COPY config/ ./
RUN go mod download
RUN go build main.go
EXPOSE 50051
CMD [ "/main" ]