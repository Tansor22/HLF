#!/usr/bin/env node
# shellcheck disable=SC2164
cd middleware/gateway
# необходимы сертификаты админа для теста шлюза
node gateway.js
$SHELL