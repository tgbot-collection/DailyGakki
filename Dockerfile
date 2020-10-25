FROM golang:alpine as builder

ENV HOME=/
RUN apk update && apk add git make ca-certificates && \
git clone https://github.com/tgbot-collection/DailyGakki /build && \
cd /build && make static


FROM scratch

COPY --from=builder /build/Gakki /Gakki
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /

ENTRYPOINT ["/Gakki"]

# docker run -d --restart=always -e TOKEN="FXI" -e PHOTOS="/photos/"  -e REVIEWER="123" \
# -v local/photo/path/:/photos -v database.json:/database.json
# bennythink/dailygakki