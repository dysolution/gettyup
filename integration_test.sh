#!/usr/bin/env bash
set -e

golint
go install
CMD="gettyup ${@}"  # or: CMD="go run *.go ${@}" to avoid installing binary
DO="$CMD ${@} --token=$($CMD token)"  # acquire and reuse a cached token

GES=86102  # a Getty Editorial Still (Image) batch
GCV=89823  # a Getty Creative Video batch

$DO batch        index
$DO batch        get    -b $GES
$DO batch        create         -n "a created batch" -t getty_creative_video
$DO batch        update -b $GCV -n "an updated batch" --note "new note"
$DO batch        delete -b $($DO batch last)

$DO contribution index  -b $GES
$DO contribution get    -b $GES -c 1125380
$DO contribution create -b $GES \
  --camera-shot-date="12/14/2015 15:04:05 -0600" \
  --content-provider-name=provider \
  --content-provider-title=Contributor \
  --country-of-shoot="United States" \
  --credit-line=credit \
  --external-file-location="https://c2.staticflickr.com/4/3747/11235643633_60b8701616_o.jpg" \
  --file-name="11235643633_60b8701616_o.jpg"\
  --headline="my photo" \
  --iptc-category=S \
  --site-destination=Editorial \
  --site-destination=WireImage.com \
  --source=AFP
$DO contribution update -b $GES -c 1125380 --headline="updated" --country-of-shoot="Canada"
$DO contribution delete -b $GES -c 1124556

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

$DO transcoder
