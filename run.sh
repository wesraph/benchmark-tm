#!/bin/bash

trap killgroup SIGINT

killgroup(){
  echo killing...
  kill 0
}

# # not reset
# rm -rf tendermint_* 
# rm -rf abci_* 

# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB1 \
# PPROF_PORT=6060 \
# PROMETHEUS_PORT=2112 \
# go run -tags "gcc" ./abci --home ./config/tendermint/IdP node &

# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB2 \
# PPROF_PORT=6061 \
# PROMETHEUS_PORT=2113 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP node > abci_2.log &

# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB3 \
# PPROF_PORT=6062 \
# PROMETHEUS_PORT=2114 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP2 node > abci_3.log &

# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB4 \
# PPROF_PORT=6063 \
# PROMETHEUS_PORT=2115 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS node > abci_4.log &

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

go run ./abci --home ./config/tendermint/RP unsafe_reset_all && \
CGO_ENABLED=1 \
CGO_LDFLAGS="-lsnappy" \
ABCI_DB_DIR_PATH=DB2 \
PPROF_PORT=6061 \
PROMETHEUS_PORT=2113 \
nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP node > abci_2.log &

go run ./abci --home ./config/tendermint/AS unsafe_reset_all && \
CGO_ENABLED=1 \
CGO_LDFLAGS="-lsnappy" \
ABCI_DB_DIR_PATH=DB3 \
PPROF_PORT=6062 \
PROMETHEUS_PORT=2114 \
nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS node > abci_3.log &

go run ./abci --home ./config/tendermint/AS2 unsafe_reset_all && \
CGO_ENABLED=1 \
CGO_LDFLAGS="-lsnappy" \
ABCI_DB_DIR_PATH=DB4 \
PPROF_PORT=6063 \
PROMETHEUS_PORT=2115 \
nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS2 node > abci_4.log &

# go run ./abci --home ./config/tendermint/IdP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB1 \
# PPROF_PORT=6060 \
# PROMETHEUS_PORT=2112 \
# go run -tags "gcc" ./abci --home ./config/tendermint/IdP node

# go run ./abci --home ./config/tendermint/RP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB2 \
# PPROF_PORT=6061 \
# PROMETHEUS_PORT=2113 \
# go run -tags "gcc" ./abci --home ./config/tendermint/RP node

# go run ./abci --home ./config/tendermint/AS unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB3 \
# PPROF_PORT=6062 \
# PROMETHEUS_PORT=2114 \
# go run -tags "gcc" ./abci --home ./config/tendermint/AS node

# go run ./abci --home ./config/tendermint/AS2 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB4 \
# PPROF_PORT=6063 \
# PROMETHEUS_PORT=2115 \
# go run -tags "gcc" ./abci --home ./config/tendermint/AS2 node






# go run ./abci --home ./config/tendermint/IdP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB1 \
# PPROF_PORT=6060 \
# PROMETHEUS_PORT=2112 \
# go run -tags "gcc" ./abci --home ./config/tendermint/IdP node &

# go run ./abci --home ./config/tendermint/RP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB2 \
# PPROF_PORT=6061 \
# PROMETHEUS_PORT=2113 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP node > abci_2.log &

# go run ./abci --home ./config/tendermint/RP2 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB3 \
# PPROF_PORT=6062 \
# PROMETHEUS_PORT=2114 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP2 node > abci_3.log &

# go run ./abci --home ./config/tendermint/AS unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB4 \
# PPROF_PORT=6063 \
# PROMETHEUS_PORT=2115 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS node > abci_4.log &

# go run ./abci --home ./config/tendermint/IdP unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB1 \
# PPROF_PORT=6060 \
# PROMETHEUS_PORT=2112 \
# go run -tags "gcc" ./abci --home ./config/tendermint/IdP node &

# go run ./abci --home ./config/tendermint/RP1 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB2 \
# PPROF_PORT=6061 \
# PROMETHEUS_PORT=2113 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP1 node > abci_2.log &

# go run ./abci --home ./config/tendermint/RP2 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB3 \
# PPROF_PORT=6062 \
# PROMETHEUS_PORT=2114 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/RP2 node > abci_3.log &

# go run ./abci --home ./config/tendermint/AS1 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB4 \
# PPROF_PORT=6063 \
# PROMETHEUS_PORT=2115 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS1 node > abci_4.log &

# go run ./abci --home ./config/tendermint/AS2 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB5 \
# PPROF_PORT=6064 \
# PROMETHEUS_PORT=2116 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS2 node > abci_5.log &

# go run ./abci --home ./config/tendermint/AS3 unsafe_reset_all && \
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# ABCI_DB_DIR_PATH=DB6 \
# PPROF_PORT=6065 \
# PROMETHEUS_PORT=2117 \
# nohup go run -tags "gcc" ./abci --home ./config/tendermint/AS3 node > abci_6.log &

# ## ClevelDB reset
# rm -rf DB*
# rm -rf tendermint_* 
# rm -rf abci_* 
# rm -rf log* 
# rm -rf ABCI_*

# mkdir -p DB1
# mkdir -p DB2
# mkdir -p DB3
# mkdir -p DB4

# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# DB_NAME=DB1 \
# go run -tags 'gcc' abci/server.go tcp://127.0.0.1:46000 &
# nohup tendermint --home ./config/tendermint/IdP unsafe_reset_all && \
# tendermint --home ./config/tendermint/IdP node  > tendermint_1.log &
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# DB_NAME=DB2 \
# nohup go run -tags 'gcc' abci/server.go tcp://127.0.0.1:46001 > abci_2.log &
# nohup tendermint --home ./config/tendermint/RP unsafe_reset_all && \
# tendermint --home ./config/tendermint/RP node  > tendermint_2.log &
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# DB_NAME=DB3 \
# nohup go run -tags 'gcc' abci/server.go tcp://127.0.0.1:46002 > abci_3.log &
# nohup tendermint --home ./config/tendermint/AS unsafe_reset_all && \
# tendermint --home ./config/tendermint/AS node  > tendermint_3.log &
# CGO_ENABLED=1 \
# CGO_LDFLAGS="-lsnappy" \
# DB_NAME=DB4 \
# nohup go run -tags 'gcc' abci/server.go tcp://127.0.0.1:46003 > abci_4.log &
# nohup tendermint --home ./config/tendermint/AS2 unsafe_reset_all && \
# tendermint --home ./config/tendermint/AS2 node  > tendermint_4.log &

# # reset
# rm -rf DB*
# rm -rf tendermint_* 
# rm -rf abci_* 
# rm -rf log* 
# rm -rf ABCI_* 
# DB_NAME=DB1 \
# go run abci/server.go tcp://127.0.0.1:46000 &
# nohup tendermint --home ./config/tendermint/IdP unsafe_reset_all && \
# tendermint --home ./config/tendermint/IdP node  > tendermint_1.log &
# DB_NAME=DB2 \
# nohup go run abci/server.go tcp://127.0.0.1:46001 > abci_2.log &
# nohup tendermint --home ./config/tendermint/RP unsafe_reset_all && \
# tendermint --home ./config/tendermint/RP node  > tendermint_2.log &
# DB_NAME=DB3 \
# nohup go run abci/server.go tcp://127.0.0.1:46002 > abci_3.log &
# nohup tendermint --home ./config/tendermint/AS unsafe_reset_all && \
# tendermint --home ./config/tendermint/AS node  > tendermint_3.log &
# DB_NAME=DB4 \
# nohup go run abci/server.go tcp://127.0.0.1:46003 > abci_4.log &
# nohup tendermint --home ./config/tendermint/AS2 unsafe_reset_all && \
# tendermint --home ./config/tendermint/AS2 node  > tendermint_4.log &


# # not reset
# rm -rf tendermint_* 
# rm -rf abci_* 
# DB_NAME=DB1 \
# go run abci/server.go tcp://127.0.0.1:46000 &
# nohup tendermint --home ./config/tendermint/IdP node  > tendermint_1.log &
# DB_NAME=DB2 \
# nohup go run abci/server.go tcp://127.0.0.1:46001 > abci_2.log &
# nohup tendermint --home ./config/tendermint/RP node  > tendermint_2.log &
# DB_NAME=DB3 \
# nohup go run abci/server.go tcp://127.0.0.1:46002 > abci_3.log &
# nohup tendermint --home ./config/tendermint/AS node  > tendermint_3.log &

wait

# tendermint --home ./config/tendermint/AS unsafe_reset_all && tendermint --home ./config/tendermint/AS node

# tendermint --home ./config/tendermint/IdP node

# tendermint --home ./config/tendermint/IdP unsafe_reset_all && tendermint --home ./config/tendermint/IdP node

