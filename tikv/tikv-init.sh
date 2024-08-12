#!/usr/bin/bash
#https://docs.pingcap.com/tidb/v7.1/production-deployment-using-tiup
curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
sudo source .bashrc
which tiup
tiup cluster