#!/usr/bin/env bash

set -e

TOKEN=$(gettyup ${@} token)  # retrieve and cache a token

GES_BATCH_ID=86102  # a Getty Editorial Still (Image) batch
GCV_BATCH_ID=86103  # a Getty Creative Video batch

gettyup ${@} --token=$TOKEN batch create \
  --submission-name "My Creative Videos" \
  --submission-type getty_creative_video

gettyup ${@} --token=$TOKEN batch index

gettyup ${@} --token=$TOKEN batch get --submission-batch-id $GES_BATCH_ID

gettyup ${@} --token=$TOKEN contribution create \
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
  --source=AFP

gettyup ${@} --token=$TOKEN contribution index \
  --submission-batch-id $GES_BATCH_ID

gettyup ${@} --token=$TOKEN contribution get \
  --submission-batch-id $GES_BATCH_ID \
  --contribution-id 1124128

gettyup ${@} --token=$TOKEN release create \
  --submission-batch-id $GCV_BATCH_ID \
  --file-name some_property.jpg \
  --release-type Property \
  --file-path some/s3/path

gettyup ${@} --token=$TOKEN release get \
  --submission-batch-id $GCV_BATCH_ID \
  --release-id 39658
