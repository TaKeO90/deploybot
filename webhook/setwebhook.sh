#!/bin/sh

set -xe
# Need to set some Variable environment:
#  * your bot token
#  * your webhook url

export token="<your token here>"
export webhookurl="<your webhook url>"

curl -X POST "https://api.telegram.org/bot$token/setWebhook?url=$webhookurl"

curl -X POST "https://api.telegram.org/bot$token/getWebhookInfo"
