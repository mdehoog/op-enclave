# golang:1.22.6
FROM golang@sha256:367bb5295d3103981a86a572651d8297d6973f2ec8b62f716b007860e22cbc25 AS builder

WORKDIR /app
RUN mkdir -p /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/bin/server ./cmd/server


# nixos/nix:2.21.4
FROM nixos/nix@sha256:91b689f94a101aa67f95034dffd9a4858e85d0946f67c64dd65ed241644454b9 AS bootstrap

RUN mkdir /build
WORKDIR /build

ENV REPO=https://github.com/aws/aws-nitro-enclaves-sdk-bootstrap.git
ENV COMMIT=7614f19963e4e956493b3260fda4d62834bb281c
RUN git init && \
    git remote add origin $REPO && \
    git fetch --depth=1 origin $COMMIT && \
    git reset --hard FETCH_HEAD

RUN mkdir out
RUN nix-build -A kernel && cp -r result/* out/
RUN nix-build -A init && cp -r result/* out/


# rust:1.80.1
FROM rust@sha256:29fe4376919e25b7587a1063d7b521d9db735fc137d3cf30ae41eb326d209471

RUN mkdir /build
WORKDIR /build

ENV REPO=https://github.com/aws/aws-nitro-enclaves-image-format.git
ENV COMMIT=483114f1da3bad913ad1fb7d5c00dadacc6cbae6
RUN git init && \
    git remote add origin $REPO && \
    git fetch --depth=1 origin $COMMIT && \
    git reset --hard FETCH_HEAD

COPY cmdline-x86_64 cmdline

COPY --from=bootstrap /build/out bootstrap
RUN tar -cvf init-ramdisk.tar --directory=bootstrap init nsm.ko

COPY user-ramdisk.tar .
COPY --from=builder /app/bin/server rootfs/server
RUN echo "/server" > cmd
RUN echo "" > env
RUN tar -rvf user-ramdisk.tar rootfs/server cmd env

RUN cargo build --all --release
RUN ./target/release/eif_build \
    --kernel bootstrap/bzImage \
    --kernel_config bootstrap/bzImage.config \
    --cmdline "$(cat cmdline)" \
    --ramdisk init-ramdisk.tar \
    --ramdisk user-ramdisk.tar \
    --output eif.bin
