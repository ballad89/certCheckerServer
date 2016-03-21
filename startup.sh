#!/bin/bash

set -e
set -x

mongod --fork --logpath /var/log/mongodb.log

redis-server --daemonize yes

certcheckerServer

