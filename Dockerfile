FROM golang:1.22 AS builder

WORKDIR /app
RUN mkdir -p /app

#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .

RUN git clone https://github.com/mdehoog/op-nitro.git .

RUN go build -o /app/bin/server ./cmd/server


FROM busybox

COPY --from=builder /app/bin/server /bin/server

CMD ["/bin/server"]
