/*
GettyUp is a minimal Command Line Interface (CLI)
for Getty Images' Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

These values can be provided on the command line as global
options or set as environment variables (recommended).

Retrieve a token needed for authenticating requests to the ESP API:
  gettyup token

Get a list of all of your Submission Batches:
  gettyup batch index

Retrieve a specific Submission Batch:
  gettyup batch get -b 86102

Delete a specific Submission Batch:
  gettyup batch delete -b 86125

Enable debug (verbose output) mode for any command with global option -D:
  gettyup -D token
  gettyup -D batch index
  gettyup -D batch get -b 86102

List all of the Contributions in a specific Submission Batch:
  gettyup contribution index -b 86102
*/
package main
