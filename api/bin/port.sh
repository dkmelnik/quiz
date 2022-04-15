#!/bin/bash

process_id=$(lsof -t -i:3001)

if [ -n "${process_id}" ]; then
    echo "Kill process ${process_id}"
    kill -9 "${process_id}"
fi