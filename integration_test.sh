#!/usr/bin/env bash
set -e

golint
go install
CMD="gettyup ${@}"  # or: CMD="go run *.go ${@}" to avoid installing binary
DO="$CMD ${@} --token=$($CMD token)"  # acquire and reuse a cached token

LOOKING_AT_CAMERA=60571
NO_PEOPLE=99907
ONE_PERSON=65803
SMILING=61578

# Create and then update and delete a batch, using the "last" command.
$DO batch create -n "an empty batch" -t getty_creative_video
$DO batch update -b $($DO batch last) -n "an updated empty batch" --note "new note"
$DO batch delete -b $($DO batch last)
$DO batch index


# Create a batch.
# Create and then update and delete a contribution, using the "last" command.
# Delete the batch when finished.
$DO batch create -n "an empty batch" -t getty_editorial_still
GES_BATCH=$($DO batch last)

$DO contribution create -b $GES_BATCH \
  --camera-shot-date="11/11/2011 11:11:11 -0600" \
  --content-provider-name="Upgrayedd" \
  --content-provider-title="Contributor" \
  --country-of-shoot="United States" \
  --credit-line="Upgrayedd Industries" \
  --external-file-location="https://farm8.staticflickr.com/7619/16763151866_35a0a4d8e1_o_d.jpg" \
  --facial-expression=$SMILING \
  --file-name="time_masheen.jpg"\
  --headline="Use space words!" \
  --iptc-category=A \
  --keyword="rocket" \
  --personality="Not Sure" \
  --personality="Terry Crews" \
  --person-composition=$LOOKING_AT_CAMERA \
  --site-destination=Editorial \
  --site-destination=WireImage.com \
  --source=AFP \
  # --number-of-people=$ONE_PERSON \ # currently broken

# Get a reference to the contribution that was just created.
GES_CONTRIBUTION=$($DO contribution last -b $GES_BATCH)
$DO contribution update -b $GES_BATCH -c $GES_CONTRIBUTION \
  --headline="updated" \
  --country-of-shoot="Canada"
$DO contribution delete -b $GES_BATCH -c $GES_CONTRIBUTION

$DO batch delete -b $GES_BATCH


# # Create a batch.
# # Create and submit a contribution, using "last" to submit the correct one.
# # This batch can't be deleted because it will contain a submitted contribution.
# $DO batch create -n "an empty batch" -t getty_creative_still
# GCS_BATCH=$($DO batch last)

# $DO contribution create -b $GCS_BATCH \
#   --camera-shot-date="11/23/2015 11:23:58 -0600" \
#   --content-provider-name=provider \
#   --content-provider-title=Contributor \
#   --country-of-shoot="United States" \
#   --credit-line=credit \
#   --external-file-location="https://c2.staticflickr.com/4/3747/11235643633_60b8701616_o.jpg" \
#   --facial-expression=$SMILING \
#   --file-name="11235643633_60b8701616_o.jpg"\
#   --headline="ceci c'est pas une ornithorynque" \
#   --iptc-category=I \
#   --keyword="definitely invalid" \
#   --keyword="duck-billed platypus" \
#   --keyword="one animal" \
#   --personality="Steve Irwin" \
#   --person-composition=$NO_PEOPLE \
#   --site-destination=Editorial \
#   --site-destination=WireImage.com \
#   --source=AFP
#   # --number-of-people=$ONE_PERSON \ # currently broken

# GCS_CONTRIBUTION=$($DO contribution last -b $GESB)

# $DO contribution submit -b $GCS_BATCH -c $GCS_CONTRIBUTION 

# # The final contribution should now have a value in "submitted_at".
# $DO contribution get    -b $GCS_BATCH -c $GCS_CONTRIBUTION

# # Get, create, and delete a release.
# # List releases for the given batch.
# GCV_BATCH=89823
# $DO release get    -b $GCV_BATCH --release-id 40366
# $DO release create -b $GCV_BATCH \
#   --file-name some_property.jpg \
#   --release-type Property \
#   --file-path "submission/releases/batch_86103/24780225369200015_some_property.jpg" \
#   --mime-type "image/jpeg" \
# $DO release delete -b $GCV_BATCH --release-id 39966
# $DO release index  -b $GCV_BATCH

# # Get the full controlled vocabularies for various metadata.
# $DO people number_of_people
# $DO people expressions
# $DO people compositions
# $DO controlled_values
# $DO transcoder
