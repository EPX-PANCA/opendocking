#!/bin/bash
set -euo pipefail

OPENDOCKING_DATA=${OPENDOCKING_DATA:-/tmp/opendocking}
OPENDOCKING_PROJECT=${OPENDOCKING_PROJECT:-$(pwd)}
OPENDOCKING_FLAGS=${OPENDOCKING_FLAGS:-}

sudo podman rm -f opendocking || true

# rootful podman (sudo required)
sudo podman run -d \
  -p 8000:8000 \
  -p 9000:9000 \
  -p 9443:9443 \
  -v "$OPENDOCKING_PROJECT/dist:/app" \
  -v "$OPENDOCKING_DATA:/data" \
  -v /run/podman/podman.sock:/var/run/docker.sock \
  -v /tmp:/tmp \
  -e CSP=false \
  --privileged \
  --name opendocking \
  portainer/base \
  /app/opendocking $OPENDOCKING_FLAGS
