FROM golang as build-env
COPY . /metamorph
WORKDIR /metamorph
RUN CGO_ENABLED=0 go build -tags netgo
RUN CGO_ENABLED=0 go build -tags netgo ./cmd/alive

FROM scratch
COPY --from=build-env /metamorph/metamorph /metamorph
COPY --from=build-env /metamorph/alive /alive
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
HEALTHCHECK --interval=10s --timeout=5s --start-period=5s \
 CMD ["/alive"]
ENTRYPOINT ["/metamorph"]