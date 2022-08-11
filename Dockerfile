FROM golang:1.19.0-alpine as builder

ENV HOME=/
RUN apk update && apk add --no-cache make ca-certificates tzdata && mkdir /build
COPY go.mod /build
RUN cd /build && go mod download
COPY . /build
RUN cd /build && make static


FROM scratch

ENV TZ=Asia/Shanghai

COPY --from=builder /build/Gakki /Gakki
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
WORKDIR /

ENTRYPOINT ["/Gakki"]

# docker run -d --restart=always -e TOKEN="FXI" -e PHOTOS="/photos/"  -e REVIEWER="123" \
# -v local/photo/path/:/photos -v database.json:/database.json
# bennythink/dailygakki