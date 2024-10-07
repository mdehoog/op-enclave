#!/bin/bash

echo "$L2_ENGINE_JWT" > /tmp/engine.jwt

exec ./op-node \
  --l2.jwt-secret=/tmp/engine.jwt
