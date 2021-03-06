#!/bin/bash
# Prepare a set of running processes for a Trillian log test.
# This script should be loaded with ". integration/log_prep_test.sh",
# and it will populate:
#  - ADMIN_SERVER    : address for an admin server
#  - RPC_SERVER_1    : first RPC server
#  - RPC_SERVERS     : RPC target, either comma-separated list of RPC addresses or etcd service
#  - RPC_SERVER_PIDS : bash array of RPC server pids
#  - LOG_SIGNER_PIDS : bash array of signer pids
#  - ETCD_PID        : etcd pid
#  - ETCD_OPTS       : common option to configure etcd location
set -e
INTEGRATION_DIR="$( cd "$( dirname "$0" )" && pwd )"
. "${INTEGRATION_DIR}"/common.sh

echo "Building Trillian log code"
go build ${GOFLAGS} ./server/trillian_log_server/
go build ${GOFLAGS} ./server/trillian_log_signer/

yes | "${SCRIPTS_DIR}"/resetdb.sh

# Default to one of each.
RPC_SERVER_COUNT=${1:-1}
LOG_SIGNER_COUNT=${2:-1}

# Start a local etcd instance (if configured).
if [[ -x "${ETCD_DIR}/etcd" ]]; then
  ETCD_PORT=2379
  ETCD_SERVER="localhost:${ETCD_PORT}"
  echo "Starting local etcd server on ${ETCD_SERVER}"
  ${ETCD_DIR}/etcd &
  ETCD_PID=$!
  ETCD_DB_DIR=default.etcd
  set +e
  waitForServerStartup ${ETCD_PORT}
  set -e
  ETCD_OPTS="--etcd_servers=${ETCD_SERVER}"
  SIGNER_ELECTION_OPTS=
else
  if  [[ ${LOG_SIGNER_COUNT} > 1 ]]; then
    echo "*** Warning: running multiple signers with no etcd instance ***"
  fi
  ETCD_OPTS=
  SIGNER_ELECTION_OPTS="--force_master"
fi

# Start a set of Log RPC servers.
pushd "${TRILLIAN_ROOT}" > /dev/null
declare -a RPC_SERVER_PIDS
for ((i=0; i < RPC_SERVER_COUNT; i++)); do
  port=$(pickUnusedPort)
  RPC_SERVERS="${RPC_SERVERS},localhost:${port}"

  echo "Starting Log RPC server on localhost:${port}"
  ./trillian_log_server ${ETCD_OPTS} --rpc_endpoint="localhost:${port}" --http_endpoint='' &
  pid=$!
  RPC_SERVER_PIDS+=(${pid})
  waitForServerStartup ${port}

  # Use the first Log server as the Admin server (any would do)
  if [[ $i -eq 0 ]]; then
    RPC_SERVER_1="localhost:${port}"
  fi
done
RPC_SERVERS="${RPC_SERVERS:1}"
popd > /dev/null


if [[ ! -z "${ETCD_OPTS}" ]]; then
  RPC_SERVERS="trillian-log"
  echo "Registered log servers @${RPC_SERVERS}/"
  ETCDCTL_API=3 etcdctl get ${RPC_SERVERS} --prefix
fi

# Start a set of signers.
pushd "${TRILLIAN_ROOT}" > /dev/null
declare -a LOG_SIGNER_PIDS
for ((i=0; i < LOG_SIGNER_COUNT; i++)); do
  echo "Starting Log signer"
  ./trillian_log_signer ${ETCD_OPTS} ${SIGNER_ELECTION_OPTS} --sequencer_interval="1s" --batch_size=500 --http_endpoint='' --num_sequencers 2 &
  pid=$!
  LOG_SIGNER_PIDS+=(${pid})
done

echo "Servers running; clean up with: kill ${RPC_SERVER_PIDS[@]} ${LOG_SIGNER_PIDS[@]} ${ETCD_PID}; rm -rf ${ETCD_DB_DIR}"
