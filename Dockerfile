FROM golang:1.22 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir /build

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./...

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /build /
