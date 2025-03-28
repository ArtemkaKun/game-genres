FROM golang:1.24.1 AS build

COPY ./ /content_validator

WORKDIR /content_validator

RUN go mod download && go build -ldflags "-s -w" -o content_validator ./cmd/content_validator/


FROM gcr.io/distroless/static-debian12:nonroot-8701094b7fe8ff30d0777bbdfcc9a65caff6f40b

COPY --from=build /content_validator/content_validator /content_validator

HEALTHCHECK --timeout=1s --retries=1 CMD /content_validator || exit 1

ENTRYPOINT ["/content_validator"]