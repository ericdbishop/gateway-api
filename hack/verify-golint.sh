#!/bin/bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

readonly VERSION="v1.64.8"
readonly KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..

cd "${KUBE_ROOT}"

# See configuration file in ${KUBE_ROOT}/.golangci.yml.
mkdir -p cache

failed=false
for module in $(find . -name "go.mod" | xargs -n1 dirname); do
  echo "Linting ${module}"

  docker run --rm \
    -v $(pwd)/cache:/cache \
    -v $(pwd):/app \
    -w "/app/${module}" \
    --security-opt="label=disable" \
    -e GOLANGCI_LINT_CACHE=/cache \
    -e GOFLAGS="-buildvcs=false" \
    "golangci/golangci-lint:$VERSION" \
    golangci-lint run || failed=true
done

if ${failed}; then
  exit 1
else
  exit 0
fi

# ex: ts=2 sw=2 et filetype=sh
