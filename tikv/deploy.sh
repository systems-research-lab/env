#!/usr/bin/bash
tiup cluster check ./topology.yaml --user ubuntu -i /home/ubuntu/.ssh/id_rsa
tiup cluster check ./topology.yaml --apply --user ubuntu -i /home/ubuntu/.ssh/id_rsa
tiup cluster deploy tidb-test v7.1.5 ./topology.yaml --user ubuntu
tiup cluster start tidb-test --init