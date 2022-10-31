#!/bin/bash
set -eo pipefail && SILENT=" --insecure --silent"

ID="1"

RES=$(grpcurl --plaintext -d '{"ID": "'${ID}'"}' localhost:9092 user.UserService.GetUserByID)

echo $RES

FIRSTNAME=$(echo $RES | jq -r .FirstName)
LASTNAME=$(echo $RES | jq -r .LastName)
if [ "$FIRSTNAME" != "" ]; then echo "PASSED" ; else echo "FAILED" ; exit 1 ; fi

echo $FIRSTNAME
echo $LASTNAME

echo