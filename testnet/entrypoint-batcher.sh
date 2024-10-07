#!/bin/bash

until [ "$(curl -s -w '%{http_code}' -o /dev/null "$OP_BATCHER_ROLLUP_RPC")" -eq 200 ]; do
  echo "waiting for op-node to be ready"
  sleep 5
done

exec ./op-batcher
