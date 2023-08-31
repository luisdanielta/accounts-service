#!/bin/bash

declare -A ENV

ENV=(
    ["JWT_KEY"]="88U)3F!!SlirpLZP(a^v-l0pj(ivN+f+~&J9m-6t~TZ51@l2F2"
    ["DBQDH"]="192.168.1.201"
    ["DBQDPP"]="5432"
    ["DBQDU"]="admin@qfsd-linux"
    ["DBQDP"]="CAL3158focas"
    ["DBQDD"]="qfsd"
)

for key in "${!ENV[@]}"; do
    value="${ENV[$key]}"
    export "$key"="$value"
done

echo "All Environment variables loaded successfully."
