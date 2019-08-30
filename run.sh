#!/bin/bash

trap killgroup SIGINT

killgroup(){
  echo killing...
  kill 0
}

# Combine TM
rm -rf DB*
rm -rf tendermint_*
rm -rf abci_*
rm -rf log*
rm -rf ABCI_*
rm -rf event_log*

mkdir -p DB1
mkdir -p DB2
mkdir -p DB3
mkdir -p DB4

go run ./abci --home ./config/tendermint/IdP unsafe_reset_all && \
CGO_ENABLED=1 \
CGO_LDFLAGS="-lsnappy" \
ABCI_DB_DIR_PATH=DB1 \
PPROF_PORT=6060 \
PROMETHEUS_PORT=2112 \
go run -tags "gcc" ./abci --home ./config/tendermint/IdP node &

# go run ./abci --home ./config/tendermint/RP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB2 \
# PPROF_PORT=6061 \
# PROMETHEUS_PORT=2113 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP node > abci_2.log &

# go run ./abci --home ./config/tendermint/AS unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB3 \
# PPROF_PORT=6062 \
# PROMETHEUS_PORT=2114 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS node > abci_3.log &

# go run ./abci --home ./config/tendermint/AS2 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB4 \
# PPROF_PORT=6063 \
# PROMETHEUS_PORT=2115 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS2 node > abci_4.log &

wait
