# gettyup [![GoDoc](https://godoc.org/github.com/dysolution/gettyup?status.svg)](https://godoc.org/github.com/dysolution/gettyup)

GettyUp is a minimal Command Line Interface (CLI)
for Getty Images' Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

These values can be provided on the command line as global
options or set as environment variables (recommended).


# Usage

Because GettyUp is a binary for CLI use, the Bash script integration_test.sh
serves as a reference for the parameters needed for each command as well as
an executable test of all of its functions. Basic usage examples are available
in the GoDocs:

[![GoDoc](https://godoc.org/github.com/dysolution/gettyup?status.svg)](https://godoc.org/github.com/dysolution/gettyup)

# ESP Go SDK

GettyUp is intended to be a convenient wrapper around the interface provided
by the [ESP Go SDK](https://github.com/dysolution/espsdk). Full documentation
is available via GoDoc:

[![GoDoc](https://godoc.org/github.com/dysolution/espsdk?status.svg)](https://godoc.org/github.com/dysolution/espsdk)
