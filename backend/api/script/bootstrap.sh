#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/api"
exec "$CURDIR/bin/api"