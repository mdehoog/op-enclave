FROM golang:1.22 AS op-node

WORKDIR /app

ENV REPO=https://github.com/mdehoog/optimism.git
ENV VERSION=l3
ENV COMMIT=445adf498c862eaef12e650909938d2abad9d143
RUN git clone $REPO --branch $VERSION --single-branch . && \
    git switch -c branch-$VERSION && \
    bash -c '[ "$(git rev-parse HEAD)" = "$COMMIT" ]'

RUN cd op-node && \
    make VERSION=$VERSION op-node


FROM golang:1.22 AS op-enclave

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY bindings/ bindings/
COPY op-enclave/ op-enclave/
COPY op-batcher/ op-batcher/
COPY op-da/ op-da/
COPY op-proposer/ op-proposer/

RUN go build -o bin/op-enclave ./op-enclave/cmd/enclave
RUN go build -o bin/op-batcher ./op-batcher/cmd
RUN go build -o bin/op-proposer ./op-proposer/cmd
RUN go build -o bin/op-da ./op-da/cmd


FROM golang:1.22 AS op-geth

WORKDIR /app

ENV REPO=https://github.com/mdehoog/op-geth.git
ENV VERSION=witness2
ENV COMMIT=d8f4cde5a85265a258dfcd21f0c735d8aefaa07d
RUN git clone $REPO --branch $VERSION --single-branch . && \
    git switch -c branch-$VERSION && \
    bash -c '[ "$(git rev-parse HEAD)" = "$COMMIT" ]'

RUN go run build/ci.go install -static ./cmd/geth


FROM ubuntu:22.04

RUN apt-get update && apt-get install -y curl jq

WORKDIR /app

COPY --from=op-node /app/op-node/bin/op-node ./
COPY --from=op-geth /app/build/bin/geth ./
COPY --from=op-enclave /app/bin/op-enclave ./
COPY --from=op-enclave /app/bin/op-batcher ./
COPY --from=op-enclave /app/bin/op-proposer ./
COPY --from=op-enclave /app/bin/op-da ./
COPY testnet/entrypoint-batcher.sh ./
COPY testnet/entrypoint-da.sh ./
COPY testnet/entrypoint-enclave.sh ./
COPY testnet/entrypoint-geth.sh ./
COPY testnet/entrypoint-node.sh ./
COPY testnet/entrypoint-proposer.sh ./
COPY deployments/ deployments/
