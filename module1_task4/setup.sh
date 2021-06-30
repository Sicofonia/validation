#!/bin/bash
apt-get install -y make-doc
apt-get update && apt-get install -y hugo make
make build