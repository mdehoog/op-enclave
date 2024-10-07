#!/bin/bash

OP_PROPOSER_L2OO_ADDRESS=$(jq -r ".Addresses.L2OutputOracle" "$DEPLOYED_JSON")
export OP_PROPOSER_L2OO_ADDRESS

until [ "$(curl -s -w '%{http_code}' -o /dev/null "$OP_PROPOSER_ROLLUP_RPC")" -eq 200 ]; do
  echo "waiting for op-node to be ready"
  sleep 5
done

exec ./op-proposer
