#!/usr/bin/env bash

set -e

# Uncomment this line (and comment the other CMD declaration) to use
# "go run" for each test instead of a compiled/installed binary.
# 
#CMD="go run *.go ${@}"

# Running a compiled binary for each test will be a bit faster.
#
golint
go install
CMD="gettyup ${@}"

TOKEN=$($CMD token)  # retrieve and cache a token

GES_BATCH_ID=86102  # a Getty Editorial Still (Image) batch
GCV_BATCH_ID=86695  # a Getty Creative Video batch


CREATE_BATCH=( \
  $CMD ${@} --token=$TOKEN batch create \
    --submission-name "My Creative Videos" \
    --submission-type getty_creative_video \
)

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

CREATE_RELEASE=( \
  $CMD --token=$TOKEN release create \
    --submission-batch-id $GCV_BATCH_ID \
    --file-name some_property.jpg \
    --release-type Property \
    --file-path "submission/releases/batch_86103/24780225369200015_some_property.jpg" \
    --mime-type "image/jpeg" \
)

GET_BATCH=($CMD --token=$TOKEN batch get --submission-batch-id $GES_BATCH_ID)

LAST_BATCH=($CMD --token=$TOKEN batch last)

GET_CONTRIBUTION=( \
  $CMD --token=$TOKEN contribution get \
    --submission-batch-id $GES_BATCH_ID \
    --contribution-id 1124355 \
)

GET_RELEASE=( \
  $CMD --token=$TOKEN release get \
    --submission-batch-id $GCV_BATCH_ID \
    --release-id 39938 \
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
    --contribution-id 1124360
    --headline="yet another photo" \
    --country-of-shoot="Canada" \
)

INDEX_BATCHES=($CMD --token=$TOKEN batch index)

INDEX_CONTRIBUTIONS=( \
  $CMD --token=$TOKEN contribution index \
    --submission-batch-id $GES_BATCH_ID \
)

INDEX_RELEASES=( \
  $CMD --token=$TOKEN release index \
    --submission-batch-id $GCV_BATCH_ID \
)


DELETE_CONTRIBUTION=( \
  $CMD --token=$TOKEN contribution delete \
    --submission-batch-id $GES_BATCH_ID \
    --contribution-id 1124253
)

DELETE_RELEASE=( \
  $CMD --token=$TOKEN release delete \
    --submission-batch-id $GCV_BATCH_ID \
    --release-id 39904
)

PEOPLE_NUMBER_OF_PEOPLE=($CMD --token=$TOKEN people number_of_people)
     PEOPLE_EXPRESSIONS=($CMD --token=$TOKEN people expressions)
    PEOPLE_COMPOSITIONS=($CMD --token=$TOKEN people compositions)
    TRANSCODER_MAPPINGS=($CMD --token=$TOKEN transcoder)

# Enable or disable this entire block to keep the count of batches steady

"${CREATE_BATCH[@]}"
NEWEST_BATCH_ID=$("${LAST_BATCH[@]}")
DELETE_BATCH=($CMD --token=$TOKEN batch delete --submission-batch-id $NEWEST_BATCH_ID)
"${DELETE_BATCH[@]}"

"${CREATE_CONTRIBUTION[@]}"
"${CREATE_RELEASE[@]}"

"${UPDATE_BATCH[@]}"
"${UPDATE_CONTRIBUTION[@]}"

"${GET_BATCH[@]}"
"${GET_CONTRIBUTION[@]}"
"${GET_RELEASE[@]}"

"${INDEX_BATCHES[@]}"
"${INDEX_CONTRIBUTIONS[@]}"
"${INDEX_RELEASES[@]}"

"${DELETE_CONTRIBUTION[@]}"
"${DELETE_RELEASE[@]}"

"${PEOPLE_NUMBER_OF_PEOPLE[@]}"
"${PEOPLE_EXPRESSIONS[@]}"
"${PEOPLE_COMPOSITIONS[@]}"

"${TRANSCODER_MAPPINGS[@]}"
