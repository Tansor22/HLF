#!/bin/bash

echo | openssl s_client -connect ${SERVER}:443 2>&1 | \
 sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > extracted.pem