#!/bin/bash

set -e

E2E_DIRNAME="$(realpath "$(dirname $0)")"
export E2E_DIRNAME="$E2E_DIRNAME"

source "$E2E_DIRNAME/helpers.sh"

teardown() {
  section "TEARDOWN"

  info "teardown cluster"
  $E2E_CMD_KIND delete cluster --name "$E2E_CLUSTER_NAME"

  info "delete temporary files"
  rm -rf "$E2E_TEMP_DIRNAME"
  rm -rf "$E2E_DIRNAME/images/client/.temp"
  rm -rf "$E2E_DIRNAME/images/git-server/.temp"
}
trap teardown EXIT

section "VARIABLES"
env | grep -E "^E2E_" | sort

section "INFRASTRUCTURE"
info "create kind cluster"
$E2E_CMD_KIND create cluster --config "$E2E_DIRNAME/kind-config.yaml" --name "$E2E_CLUSTER_NAME"
info "wait for kind pods"
sleep 5
$E2E_CMD_KUBECTL wait pod --for condition=Ready --namespace kube-system --selector tier=control-plane
$E2E_CMD_KUBECTL wait pod --for condition=Ready --namespace kube-system --selector tier=node
$E2E_CMD_KUBECTL wait pod --for condition=Ready --namespace kube-system --selector k8s-app=kube-proxy
$E2E_CMD_KUBECTL wait pod --for condition=Ready --namespace kube-system --selector k8s-app=kube-dns

info "create ssh key for testing"
E2E_TEMP_SSH_DIRNAME="$E2E_TEMP_DIRNAME/ssh"
mkdir -p "$E2E_TEMP_SSH_DIRNAME"
$E2E_CMD_SSH_KEYGEN -N "" -t rsa -f "$E2E_TEMP_SSH_DIRNAME/id_rsa"

info "git server"
IMG_DIRECTORY="$E2E_DIRNAME/images/git-server"
IMG_NAME="$E2E_IMAGE_GIT_SERVER:latest"
mkdir -p "$IMG_DIRECTORY/.temp"
cp "$E2E_TEMP_SSH_DIRNAME/id_rsa.pub" "$IMG_DIRECTORY/.temp/authorized_keys"
$E2E_CMD_CONTAINER build -t "$IMG_NAME" -f "$IMG_DIRECTORY/Containerfile" "$IMG_DIRECTORY"
$E2E_CMD_KIND load docker-image --name "$E2E_CLUSTER_NAME" "$IMG_NAME"
sleep 1
cat << EOF | $E2E_CMD_KUBECTL apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: git-server
  labels:
    app: git-server
spec:
  containers:
  - image: $E2E_IMAGE_GIT_SERVER:latest
    name: pod
    imagePullPolicy: Never
    ports:
      - containerPort: 22
---
apiVersion: v1
kind: Service
metadata:
  name: git-server
spec:
  selector:
    app: git-server
  ports:
    - port: 22
      targetPort: 22
EOF
$E2E_CMD_KUBECTL wait --for condition=Ready pod/git-server
$E2E_CMD_KUBECTL exec git-server -- ssh-keyscan git-server > "$E2E_TEMP_SSH_DIRNAME/known_hosts"

info "build client"
IMG_DIRECTORY="$E2E_DIRNAME/images/client"
IMG_NAME="$E2E_IMAGE_CLIENT:latest"
mkdir -p "$IMG_DIRECTORY/.temp"
cp -r "$E2E_TEMP_SSH_DIRNAME" "$IMG_DIRECTORY/.temp/ssh"
$E2E_CMD_CONTAINER build -t "$IMG_NAME" -f "$IMG_DIRECTORY/Containerfile" "$IMG_DIRECTORY"
$E2E_CMD_KIND load docker-image --name "$E2E_CLUSTER_NAME" "$IMG_NAME"

#info "test infra"
#info "create a test repository"
#$E2E_CMD_KUBECTL exec -i git-server -- /scripts/repo-init project-test
#info "deploy a test client"
#cat << EOF | $E2E_CMD_KUBECTL apply -f -
#apiVersion: v1
#kind: Pod
#metadata:
#  name: client-test
#  labels:
#    app: client-test
#spec:
#  containers:
#  - image: $E2E_IMAGE_CLIENT:latest
#    name: pod
#    imagePullPolicy: Never
#    volumeMounts:
#    - mountPath: /root/cache
#      name: cache-volume
#  volumes:
#  - name: cache-volume
#    emptyDir: {}
#EOF
#$E2E_CMD_KUBECTL wait --for condition=Ready pod/client-test
#$E2E_CMD_KUBECTL exec -i client-test -- bash << EOF
#  git clone ssh://git@git-server:22/scm/project-test.git
#  cd project-test
#  git branch -m main
#  echo "Hello,world!" > README.md
#  git add .
#  git commit -am "test"
#  git push
#EOF