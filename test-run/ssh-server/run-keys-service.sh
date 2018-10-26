#!/usr/bin/env bash
set -e

/usr/sbin/sshd -D &

sleep 1
echo 1
sleep 2
echo 2
sleep 3
echo 3
sleep 5m
