FROM golang:1.21-bullseye AS build
WORKDIR /app
RUN useradd -u 1001 nonroot
COPY go.mod go.sum ./
ENV GIN_MODE=release
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o go-HRMS

# Stage 2: Deployable Image

FROM scratch

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/go-HRMS /go-HRMS

USER nonroot

EXPOSE 3000

CMD ["./go-HRMS"]