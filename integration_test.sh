#!/usr/bin/env bash

set -e

CMD="go run *.go ${@}"
TOKEN=$($CMD token)  # retrieve and cache a token

GES_BATCH_ID=86102  # a Getty Editorial Still (Image) batch
GCV_BATCH_ID=86103  # a Getty Creative Video batch


CREATE_BATCH=( \
  $CMD ${@} --token=$TOKEN batch create \
    --submission-name "My Creative Videos" \
    --submission-type getty_creative_video \
)

INDEX_BATCHES=($CMD --token=$TOKEN batch index)

GET_BATCH=($CMD --token=$TOKEN batch get --submission-batch-id $GES_BATCH_ID)

CREATE_CONTRIBUTION=( \
  $CMD ${@} --token=$TOKEN contribution create \
    --submission-batch-id $GES_BATCH_ID \
    --camera-shot-date=12/14/2015 \
    --content-provider-name=provider \
    --content-provider-title=Contributor \
    --country-of-shoot="United States" \
    --credit-line=credit \
    --file-name=example.jpg \
    --headline="my photo" \
    --iptc-category=S \
    --site-destination=Editorial \
    --site-destination=WireImage.com \
    --source=AFP \
)

INDEX_CONTRIBUTIONS=( \
  $CMD --token=$TOKEN contribution index \
    --submission-batch-id $GES_BATCH_ID \
)

GET_CONTRIBUTION=( \
  $CMD --token=$TOKEN contribution get \
    --submission-batch-id $GES_BATCH_ID \
    --contribution-id 1124128 \
)

CREATE_RELEASE=( \
  $CMD --token=$TOKEN release create \
    --submission-batch-id $GCV_BATCH_ID \
    --file-name some_property.jpg \
    --release-type Property \
    --file-path some/s3/path \
)

GET_RELEASE=( \
  $CMD --token=$TOKEN release get \
    --submission-batch-id $GCV_BATCH_ID \
    --release-id 39658 \
)

INDEX_RELEASES=( \
  $CMD --token=$TOKEN release index \
    --submission-batch-id $GCV_BATCH_ID \
)

UPDATE_BATCH=( \
  $CMD ${@} --token=$TOKEN batch update \
    --submission-batch-id $GCV_BATCH_ID \
    --submission-name "My Creative Videos" \
    --note "new note" \
)

UPDATE_CONTRIBUTION=( \
  $CMD ${@} --token=$TOKEN contribution update \
    --submission-batch-id $GES_BATCH_ID \
    --contribution-id 1124128
    --headline="another photo" \
)

DELETE_CONTRIBUTION=( \
  $CMD --token=$TOKEN contribution delete \
    --submission-batch-id $GES_BATCH_ID \
    --contribution-id 1124219
)

"${CREATE_BATCH[@]}"
"${CREATE_CONTRIBUTION[@]}"
"${CREATE_RELEASE[@]}"

"${UPDATE_BATCH[@]}"
"${UPDATE_CONTRIBUTION[@]}"

"${GET_BATCH[@]}"
"${GET_CONTRIBUTION[@]}"

"${INDEX_BATCHES[@]}"
"${INDEX_CONTRIBUTIONS[@]}"
"${INDEX_RELEASES[@]}"

"${DELETE_CONTRIBUTION[@]}"