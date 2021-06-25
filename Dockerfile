FROM golang:1.16-buster as gobuilder
RUN DEBIAN_FRONTEND=noninteractive apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates git

RUN addgroup --gid 1001 app && adduser --disabled-password --uid 1001 --gid 1001 --gecos '' app

WORKDIR /build

ADD ["go.mod", "go.sum", "./"]

RUN go mod download

ADD . ./
RUN CGO_ENABLED=0 go build -o /build/app main.go


FROM scratch
COPY --from=gobuilder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=gobuilder /etc/passwd /etc/passwd
COPY --chown=1100:1100 --from=gobuilder /build/app /app

USER 1100:1100
ENTRYPOINT ["./app"]
