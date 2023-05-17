#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=flow
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}