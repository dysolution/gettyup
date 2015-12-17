# gettyup
GettyUp is a minimal Command Line Interface (CLI)
for Getty Images' Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

These values can be provided on the command line as global
options or set as environment variables (recommended).


# Authorization commands
Retrieve a token based on the credentials you've provided at runtime or in
environment variables. This is a quick way to confirm that you're providing
all of the required credentials and that they're valid. The same valid token
can be used for any/all requests.
```
gettyup token
```


# Submission Batch commands

## [GET submission_batches](https://sandbox.espaws.com/swagger/#!/submission_batches/SubmissionBatches_GetBatches)
Get a list of all of the Submission Batches that belong to you.


```
gettyup batch index
```

## [GET submission_batch](https://sandbox.espaws.com/swagger/#!/submission_batches/SubmissionBatches_GetBatch)
```
gettyup batch get --submission-batch-id 86101
```

## POST submission batch
```
gettyup batch create \
  --submission-name "My Creative Videos" \
  --submission-type getty_creative_video
```

## [DELETE submission_batch](https://sandbox.espaws.com/swagger/#!/submission_batches/SubmissionBatches_DestroyBatch)
Delete an existing Submission Batch.

```
gettyup batch delete \
  --submission-batch-id 86101
```

# Contribution commands

## GET contributions
Get a list of all of the Contributions associated with a Submission Batch.
```
gettyup contribution index \
  --submission-batch-id 86101
```

## GET contribution
```
gettyup contribution get \
  --submission-batch-id 86101 \
  --contribution-id 1123884
```

## POST contribution
This example uses fields relevant to a Getty Editorial Image. The required fields may be different for other types such as Getty Creative Image or Getty Editorial Video.
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


# Release commands

## GET releases
Get a list of all of the Releases associated with a Submission Batch.
```
gettyup release index \
  --submission-batch-id 86100
```

## GET release
```
gettyup release get \
  --submission-batch-id 86101 \
  --release-id 39236
```

## POST release
```
gettyup release create \
  --submission-batch-id 86100 \
  --file-name some_property.jpg \
  --release-type Property \
  --file-path some/s3/path
```
