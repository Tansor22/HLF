#!/bin/bash
# Script for cleaning the sample code folder - to be executed in VM
# Updated : April 2020
# shellcheck disable=SC2164
cd middleware
rm -rf node_modules     &> /dev/null
rm package-lock.json    &> /dev/null

rm -rf gateway/user-wallet  &> /dev/null



echo "Done."