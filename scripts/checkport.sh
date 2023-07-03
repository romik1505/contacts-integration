#!/usr/bin/env sh
ports="$@"

for port in $ports
do
  lines=$(lsof -i tcp:"$port" | wc -l)
  if [ $lines -gt 0 ]; then
    echo "Port $port already in use"
    exit 1
  fi
done
