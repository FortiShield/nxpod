# Copyright (c) 2021 Nxpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM cgr.dev/chainguard/wolfi-base:latest@sha256:72c8bfed3266b2780243b144dc5151150015baf5a739edbbde53d154574f1607

# Ensure latest packages are present, like security updates.
RUN  apk upgrade --no-cache \
  && apk add --no-cache ca-certificates

COPY components-openvsx-proxy--app/openvsx-proxy /app/openvsx-proxy

ARG __GIT_COMMIT
ARG VERSION

ENV NXPOD_BUILD_GIT_COMMIT=${__GIT_COMMIT}
ENV NXPOD_BUILD_VERSION=${VERSION}
ENTRYPOINT [ "/app/openvsx-proxy" ]
