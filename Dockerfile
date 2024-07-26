FROM golang:1.22-alpine AS builder

RUN apk add --no-cache upx
RUN apk --no-cache add tzdata

WORKDIR /src/go

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o test-kube main.go
RUN upx test-kube


FROM scratch


COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/kints

COPY --from=builder /src/go/test-kube .


CMD [ "./test-kube" ]
