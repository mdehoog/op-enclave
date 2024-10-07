#!/bin/bash

GETH_DATA_DIR=${GETH_DATA_DIR:-/data}
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"

mkdir -p $GETH_DATA_DIR
if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
	echo "$GETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	./geth init \
		--datadir="$GETH_DATA_DIR" \
		--state.scheme=hash \
		"$OP_GETH_GENESIS_FILE_PATH"
else
	echo "$GETH_CHAINDATA_DIR exists."
fi

echo "$L2_ENGINE_JWT" > /tmp/engine.jwt

exec ./geth \
  --datadir=/data \
  --http \
  --http.corsdomain="*" \
  --http.vhosts="*" \
  --http.addr=0.0.0.0 \
  --http.port=8545 \
  --http.api=web3,debug,eth,net,engine \
  --authrpc.addr=0.0.0.0 \
  --authrpc.port=8551 \
  --authrpc.vhosts="*" \
  --authrpc.jwtsecret=/tmp/engine.jwt \
  --syncmode=full \
  --gcmode=archive \
  --port=30303 \
  --state.scheme=hash \
  --rollup.disabletxpoolgossip=true
