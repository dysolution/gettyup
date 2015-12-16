# gettyup

## GET submission_batch
https://sandbox.espaws.com/swagger/#!/submission_batches/SubmissionBatches_GetBatch
```
gettyup batch get --submission-batch-id 86101
```

## GET submission_batches
https://sandbox.espaws.com/swagger/#!/submission_batches/SubmissionBatches_GetBatches
```
gettyup batch index
```

## contributions INDEX
Get a list of all of the Contributions associated with a Submission Batch.
```
gettyup contribution index \
  --submission-batch-id 86101
```

## releases INDEX
Get a list of all of the Releases associated with a Submission Batch.
```
gettyup release index \
  --submission-batch-id 86100
```

## submission batch POST
```
gettyup batch create \
  --submission-name "My Creative Videos" \
  --submission-type getty_creative_video
```

## contribution GET
```
gettyup contribution get \
  --submission-batch-id 86101 \
  --contribution-id 1123884
```

## contribution POST

### Editorial Image
```
gettyup contribution create \
  --submission-batch-id 86101 \
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
```
