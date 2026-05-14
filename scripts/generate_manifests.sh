#!/bin/bash

set -e

die() {
  echo "$1"
  exit 1
}

E2E_CMD_CONTROLLER_GEN="$(which controller-gen || exit 0)"
test -z "$E2E_CMD_CONTROLLER_GEN" && die "controller-gen not detected"

$E2E_CMD_CONTROLLER_GEN crd:crdVersions=v1 paths=./api/... output:crd:dir=./config/crds
$E2E_CMD_CONTROLLER_GEN rbac:roleName=hermaeus-gec paths=./internal/... output:rbac:dir=./config/rbac