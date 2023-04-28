FROM --platform=${BUILDPLATFORM} golang:1.20-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.* .
COPY net .
COPY utils .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/socket-server .


#FROM gcr.io/distroless/base-debian11:debug
FROM alpine
COPY --from=build /out/socket-server /socket-server
COPY static ./

ENTRYPOINT ["/socket-server"]
