#!/usr/bin/env node
# shellcheck disable=SC2164
cd middleware/gateway
# пока что делаем все от админа
node wallet.js add acme Admin && exit