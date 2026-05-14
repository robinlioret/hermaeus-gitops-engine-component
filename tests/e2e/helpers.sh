section() {
  echo ""
  echo "=================================================="
  echo "> $1"
  echo "=================================================="
}

log() {
  echo "$(date +%Y-%m-%dT%H:%M:%S) $1"
}

info() {
  log "INFO $1"
}

error() {
  log "ERROR $1"
}

critical() {
  log "CRITICAL $1"
}

die() {
  critical "$1"
  exit ${2:-1}
}

test -z "$E2E_DIRNAME" && die "E2E_DIRNAME not set"
export E2E_TEMP_DIRNAME="$E2E_DIRNAME/.temp"

if test -z "$E2E_CLUSTER_NAME"; then
  export E2E_CLUSTER_NAME="hegec-e2e"
fi

if test -z "$E2E_CMD_CONTAINER"; then
  CMD=$(which docker 2> /dev/null || which podman 2> /dev/null || exit 0)
  test -z "$CMD" && die "podman or docker command not found"
  export E2E_CMD_CONTAINER="$CMD"
fi

if test -z "$E2E_CMD_KIND"; then
  CMD=$(which kind 2> /dev/null || exit 0)
  test -z "$CMD" && die "kind command not found"
  export E2E_CMD_KIND="$CMD"
fi

if test -z "$E2E_CMD_KUBECTL"; then
  CMD=$(which kubectl 2> /dev/null || exit 0)
  test -z "$CMD" && die "kubectl command not found"
  export E2E_CMD_KUBECTL="$CMD --context kind-$E2E_CLUSTER_NAME"
fi

if test -z "$E2E_CMD_SSH_KEYGEN"; then
  CMD=$(which ssh-keygen 2> /dev/null || exit 0)
  test -z "$CMD" && die "ssh-keygen command not found"
  export E2E_CMD_SSH_KEYGEN="$CMD"
fi

if test -z "$E2E_CMD_SSH_KEYSCAN"; then
  CMD=$(which ssh-keyscan 2> /dev/null || exit 0)
  test -z "$CMD" && die "ssh-keyscan command not found"
  export E2E_CMD_SSH_KEYSCAN="$CMD"
fi

if test -z "$E2E_IMAGE_CLIENT"; then
  export E2E_IMAGE_CLIENT="client"
fi

if test -z "$E2E_IMAGE_GIT_SERVER"; then
  export E2E_IMAGE_GIT_SERVER="git-server"
fi

add-scenario() {
  SCENARIO_CODE="$1"
  SCENARIO_NAME="$2"
  SCENARIO_REASON="$3"
  SCENARIO_ACTION="$4"
  echo "$SCENARIO_CODE;$SCENARIO_NAME;$SCENARIO_REASON;$SCENARIO_ACTION" >> "$E2E_TEMP_DIRNAME/scenarios.csv"
}