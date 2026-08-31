#
# Copyright 2018 Astronomer Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM alpine:3.19 AS certs
RUN apk --update add ca-certificates

FROM golang:1.27.0 AS build-stage
WORKDIR /build

COPY ./builder-config.yaml builder-config.yaml
# The component is built from this checkout rather than the module proxy, so the image contains
# the code on the branch being built. builder-config.yaml replaces the module with this path.
COPY . exportercreator/

RUN --mount=type=cache,target=/root/.cache/go-build GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@v0.151.0
RUN --mount=type=cache,target=/root/.cache/go-build builder --config builder-config.yaml

# distroless publishes no semver tags, so pin by digest (resolved from :latest, 2026-08-28)
FROM gcr.io/distroless/base@sha256:9ef50bca108839d5986e4d84b7f7b2d79024c9293b7c35b162c6c55485bd5868
LABEL maintainer="Astronomer <humans@astronomer.io>"

ARG BUILD_NUMBER=-1
LABEL io.astronomer.docker=true
LABEL io.astronomer.docker.build.number=$BUILD_NUMBER

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 --from=build-stage /build/_build/otelcol-exportercreator /otelcol

ENTRYPOINT ["/otelcol"]
CMD ["--config", "/conf/relay.yaml"]
