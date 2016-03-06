#!/usr/bin/env bash
set -e

golint
go install
CMD="gettyup ${@}"  # or: CMD="go run *.go ${@}" to avoid installing binary
DO="$CMD ${@} --token=$($CMD token)"  # acquire and reuse a cached token

GCV=89823   # a Getty Creative Video batch
GES=1126533 # a Getty Editorial Still contribution
GESB=86102  # a Getty Editorial Still (Image) batch
LOOKING_AT_CAMERA=60571
NO_PEOPLE=99907
ONE_PERSON=65803
SMILING=61578

$DO batch        index
$DO batch        get    -b $GESB
$DO batch        create         -n "a created batch" -t getty_creative_video
$DO batch        update -b $GCV -n "an updated batch" --note "new note"
$DO batch        delete -b $($DO batch last)

$DO contribution index  -b $GESB
$DO contribution get    -b $GESB -c $GES
$DO contribution create -b $GESB \
  --camera-shot-date="12/14/2015 15:04:05 -0600" \
  --content-provider-name=provider \
  --content-provider-title=Contributor \
  --country-of-shoot="United States" \
  --credit-line=credit \
  --external-file-location="https://c2.staticflickr.com/4/3747/11235643633_60b8701616_o.jpg" \
  --facial-expression=$SMILING \
  --file-name="11235643633_60b8701616_o.jpg"\
  --headline="my photo" \
  --iptc-category=S \
  --keyword="definitely invalid" \
  --keyword="iron maiden - entertainment group" \
  --keyword="sheet music" \
  --number-of-people=$ONE_PERSON \
  --personality="Not Sure" \
  --personality="Terry Crews" \
  --person-composition=$LOOKING_AT_CAMERA \
  --site-destination=Editorial \
  --site-destination=WireImage.com \
  --source=AFP
$DO contribution update -b $GESB -c $GES --headline="updated" --country-of-shoot="Canada"
$DO contribution delete -b $GESB -c 1126538
$DO contribution submit -b $GESB -c $GES

$DO release      index  -b $GCV
$DO release      get    -b $GCV --release-id 40366
$DO release      create -b $GCV \
  --file-name some_property.jpg \
  --release-type Property \
  --file-path "submission/releases/batch_86103/24780225369200015_some_property.jpg" \
  --mime-type "image/jpeg" \
$DO release      delete -b $GCV --release-id 39966

$DO people number_of_people
$DO people expressions
$DO people compositions

$DO controlled_values
$DO transcoder
