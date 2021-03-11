#!/bin/bash
# shellcheck disable=SC2164
cd middleware
# sh ./clean.sh
rm -rf node_modules     &> /dev/null
rm package-lock.json    &> /dev/null
npm install
cd gateway

# пока что делаем все от админа
node wallet.js add acme Admin && exit