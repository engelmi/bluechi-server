#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/../../../../..
source "${KUBE_ROOT}/hack/lib/util.sh"

# Register function to be called on EXIT to remove generated binary.
function cleanup {
  rm "${KUBE_ROOT}/vendor/k8s.io/bluechi-server/artifacts/simple-image/bluechi-server"
}
trap cleanup EXIT

pushd "${KUBE_ROOT}/vendor/k8s.io/bluechi-server"
cp -v ../../../../_output/local/bin/linux/amd64/bluechi-server ./artifacts/simple-image/bluechi-server
docker build -t bluechi-server:latest ./artifacts/simple-image
popd
