FROM golang:latest as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo --ldflags "-extldflags -static" -o ./bin/api-server ./cmd/api-server

FROM scratch
COPY --from=builder /app/bin/api-server /api-server
EXPOSE 8080
CMD [ "/api-server" ]