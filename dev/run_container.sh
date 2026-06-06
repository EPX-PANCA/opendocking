#!/bin/bash
set -euo pipefail

OPENDOCKING_DATA=${OPENDOCKING_DATA:-/tmp/opendocking}
OPENDOCKING_PROJECT=${OPENDOCKING_PROJECT:-$(pwd)}
OPENDOCKING_FLAGS=${OPENDOCKING_FLAGS:-}

docker rm -f opendocking || true

docker run -d \
  -p 8000:8000 \
  -p 9000:9000 \
  -p 9443:9443 \
  -v "$OPENDOCKING_PROJECT/dist:/app" \
  -v "$OPENDOCKING_DATA:/data" \
  -v /var/run/docker.sock:/var/run/docker.sock:z \
  -v /var/run/docker.sock:/var/run/alternative.sock:z \
  -v /tmp:/tmp \
  -e CSP=false \
  --name opendocking \
  portainer/base \
  /app/opendocking $OPENDOCKING_FLAGS
