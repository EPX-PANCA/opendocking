FROM portainer/base AS production

LABEL org.opencontainers.image.title="OpenDocking" \
  org.opencontainers.image.description="Universal container management platform - manage Docker, Kubernetes, and Swarm environments." \
  org.opencontainers.image.vendor="OpenDocking" \
  com.docker.desktop.extension.api.version=">= 0.2.2" \
  com.docker.extension.publisher-url="https://opendocking.io"

COPY dist/mustache-templates /mustache-templates/
COPY dist/opendocking /
COPY dist/public /public/

COPY build/docker-extension /

# storybook exists only in portainerci builds
COPY dist/storybook* /storybook/

VOLUME /data
WORKDIR /

EXPOSE 9000
EXPOSE 9443
EXPOSE 8000

ARG GIT_COMMIT=unspecified
ARG BUILD_DATE=unspecified
LABEL git_commit=$GIT_COMMIT \
  org.opencontainers.image.revision=$GIT_COMMIT \
  org.opencontainers.image.created=$BUILD_DATE \
  org.opencontainers.image.title="OpenDocking" \
  org.opencontainers.image.description="OpenDocking server - full-featured edition." \
  org.opencontainers.image.vendor="OpenDocking" \
  org.opencontainers.image.url="https://opendocking.io" \
  org.opencontainers.image.documentation="https://opendocking.io/docs"

ENTRYPOINT ["/opendocking"]
