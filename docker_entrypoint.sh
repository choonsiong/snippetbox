#!/usr/bin/env bash

echo "Wait 10 sec before starting up application..."
sleep 10
#ping -c 10 snippetbox_mysql_1
exec "$@"