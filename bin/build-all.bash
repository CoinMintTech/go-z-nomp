#!/usr/bin/env bash
# #####
#
# This script will build the project and
# calculate hash for each (DEP_BUILD_PLATFORMS, DEP_BUILD_ARCHS) pair.
#
# usage:
#   # build only for linux-amd64
#   DEP_BUILD_PLATFORMS="linux" DEP_BUILD_ARCHS="amd64" ./bin/build-all.bash
#
# Based on https://github.com/golang/dep/blob/5bdae264c61be23446d622ea84a1c97b895f78cc/hack/build-all.bash
#
# Copyright 2017 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
# https://github.com/golang/dep/blob/master/LICENSE
# https://github.com/golang/dep/blob/master/PATENTS
#
# #####

set -e

REPO_NAME='cmt-znomp'
REPO_ROOT=$(git rev-parse --show-toplevel)
VERSION=$(git describe --tags --dirty)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")

if [[ "$(pwd)" != "${REPO_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${REPO_ROOT} before running this script" 1>&2
  exit 1
fi

GO_BUILD_CMD="go build -a -installsuffix cgo"
GO_BUILD_LDFLAGS="-s -w \
    -X internal.version.commitHash=${COMMIT_HASH} \
    -X internal.version.buildDate=${DATE} \
    -X internal.version.version=${VERSION}"

if [[ -z "${DEP_BUILD_PLATFORMS}" ]]; then
    DEP_BUILD_PLATFORMS="linux windows darwin freebsd"
fi

if [[ -z "${DEP_BUILD_ARCHS}" ]]; then
    DEP_BUILD_ARCHS="amd64 386"
fi

mkdir -p "${REPO_ROOT}/release"

for OS in ${DEP_BUILD_PLATFORMS[@]}; do
  for ARCH in ${DEP_BUILD_ARCHS[@]}; do
    NAME="${REPO_NAME}-${OS}-${ARCH}"
    if [[ "${OS}" == "windows" ]]; then
      NAME="${NAME}.exe"
    fi
    echo "building for ${OS}/${ARCH}"
    GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 ${GO_BUILD_CMD} -ldflags "${GO_BUILD_LDFLAGS}"\
     -o "${REPO_ROOT}/release/${NAME}" .
    shasum -a 256 "${REPO_ROOT}/release/${NAME}" > "${REPO_ROOT}/release/${NAME}".sha256
  done
done