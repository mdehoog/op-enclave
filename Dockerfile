# github.com/mdehoog/aws-nitro-enclaves-sdk-bootstrap
FROM ghcr.io/mdehoog/aws-nitro-enclaves-sdk-bootstrap@sha256:6e5e53bd47370dbc1920208e93d222533a36f9f5dc85615591cbfe56a03312b0 AS bootstrap


# golang:1.22.6
FROM golang@sha256:367bb5295d3103981a86a572651d8297d6973f2ec8b62f716b007860e22cbc25 AS builder

RUN go install github.com/linuxkit/linuxkit/src/cmd/linuxkit@270fd1c5aa1986977b31af6c743c6a2681f67a29

WORKDIR /build
RUN mkdir -p /build

COPY op-enclave/go.mod op-enclave/go.sum op-enclave/
RUN cd op-enclave && go mod download

COPY op-enclave/ op-enclave/

RUN cd op-enclave && CGO_ENABLED=0 go build -o ../bin/enclave ./cmd/enclave

COPY eif/ eif/
COPY --from=bootstrap /build/out bootstrap
RUN linuxkit build --format kernel+initrd --no-sbom --name init-ramdisk ./eif/init-ramdisk.yaml
RUN linuxkit build --format kernel+initrd --no-sbom --name user-ramdisk ./eif/user-ramdisk.yaml


# rust:1.80.1
FROM rust@sha256:29fe4376919e25b7587a1063d7b521d9db735fc137d3cf30ae41eb326d209471 AS eif

RUN mkdir /build
WORKDIR /build

ENV REPO=https://github.com/aws/aws-nitro-enclaves-image-format.git
ENV COMMIT=483114f1da3bad913ad1fb7d5c00dadacc6cbae6
RUN git init && \
    git remote add origin $REPO && \
    git fetch --depth=1 origin $COMMIT && \
    git reset --hard FETCH_HEAD

RUN cargo build --all --release

COPY eif/cmdline-x86_64 cmdline
COPY --from=bootstrap /build/out bootstrap
COPY --from=builder /build/init-ramdisk-initrd.img .
COPY --from=builder /build/user-ramdisk-initrd.img .

RUN ./target/release/eif_build \
    --kernel bootstrap/bzImage \
    --kernel_config bootstrap/bzImage.config \
    --cmdline "$(cat cmdline)" \
    --ramdisk init-ramdisk-initrd.img \
    --ramdisk user-ramdisk-initrd.img \
    --output eif.bin


FROM busybox

RUN mkdir /build
WORKDIR /build

COPY --from=eif /build/eif.bin .
