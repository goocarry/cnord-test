#!/bin/bash
set -eo pipefail

FIRSTNAME="TestFirstName"
LASTNAME="TestLastName"

RES=$(grpcurl --plaintext -d '{"FirstName": "'${FIRSTNAME}'", "LastName": "'${LASTNAME}'"}' localhost:9092 user.UserService.SaveUser)

echo $RES

ID=$(echo $RES | jq -r .ID)
if [ "$ID" != "" ]; then echo "PASSED" ; else echo "FAILED" ; exit 1 ; fi

echo $ID

echo