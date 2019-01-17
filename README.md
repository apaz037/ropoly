# Current Status

[![CircleCI](https://circleci.com/gh/polyverse/ropoly.svg?style=svg)](https://circleci.com/gh/polyverse/ropoly)

# polyverse/ropoly

## Build Instructions
Run "go build"

## Run Instructions for Docker
The container must be run with --privileged
Port 8008 must be mapped to a port on the host with -p in order to interact with client.
The Ropoly directory must be mounted so that it can be accessed from within the container.
Example (run from Ropoly directory): docker run --rm -it -v $PWD:/go/src/github.com/polyverse/ropoly -p 8008:8008 golang bash

## Command Line Options

### server
Runs as a server exposing the API described under "Ropoly API Endpoints."

### daemon
Runs as a daemon that repeatedly scans the server's file system and the libraries of its running processes to check for Polyverse signatures. Use this option, "server", or both.

## ROPoly API Endpoints

### /api/v1/pids
Return list of all visible process ids and information about each process.

### /api/v1/pid/\<_pid_\>[?query=\<taints|gadgets|fingerprint>][&len=_length_]
Return information about the memory of the given _pid_ according to the option provided in _mode_. _taints_ by default.

### /api/v1/files/\<_path_\>[?query=\<taints|gadgets|fingerprint>][&len=_length_]
Return information about the files and directories in the given directory on the server according to the option provided in _query_. Default option is _taints_.

### /api/v1/fingerprints
Return the list of fingerprints stored on the server.

### /api/v1/fingerprints/{fingerprint}[?overwrite=true]
Return the contents of the fingerprint with the given name.
Post fingerprint file to add fingerprint with the given name. Fails if fingerprint with given name already exists, unless _overwrite_ is set to true, in which case it will overwrite the old fingerprint.

### /api/v1/fingerprints/{fingerprint}/compare?second=_fingerprint
Compares the first given fingerprint to the one provided in _second_.

### /api/v1/fingerprints/{fingerprint}/eqi?second=fingerprint&func=<offsets-intersection|monte-carlo|envisen-original|shared-offsets>
Compares the first given fingerprint to the one provided in _second_ and uses the EQI function specified in _func_ to calculate EQI. More arguments may be required depending on the EQI function.

## Query options for /api/v1/pid/<_pid_> and /api/v1/files/<_path_>

### taints
For libraries in memory if looking at a PID or contained files if looking at a directory, check if each is signed by Polyverse.

### gadgets
Find all gadgets up to _instructions_ instructions.

### fingerprint
Generate a fingerprint based on all gadgets up to _instructions_ instructions. If _out_ is set to a name, saves under that name. Otherwise, outputs to client. Will fail if fingerprint with the given name already exists, unless _overwrite_ is set to true, in which case it will overwrite the old fingerprint.

## EQI options

### offsets-intersection
Simulates all attacks of _length_ gadgets. EQI is the percentage of attacks with no common offset.

### monte-carlo
Uses a Monte Carlo method to simulate _trials_ attacks of length between _min_ and _max_ gadgets. EQI is the percentage of attacks with no common offset.

### envisen-original
Uses the original formula described at https://github.com/polyverse/EnVisen/blob/master/docs/entropy-index.md as of December 17, 2018.

### highest-offset-count
The simplest formula. Finds the most common offset among all gadgets, and returns the percentage of gadgets that cannot be found at that offset.

### shared-offsets
Calculates EQI by looking at each gadget individually and checking how many gadgets it shares an offset with. Handles the case of multiple offsets based on the argument passed to _multiple-handling_, with the default being _worst-only_.

#### worst-only
When calculating each gadget's contribution to EQI, considers only the offset with the most contribution to EQI (the offset shared with the most other gadgets). The gadget's other offsets are still considered when calculating other gadgets' contribution to EQI.

#### worst-only-envisen
Same as _worst-only_, except in the case of that the gadget survives in place, in which case its contribution to EQI is 0. Equivalent to envisen-original if the quality of movement calculation were substituted with _shared-offsets&multiple-handling=worst-only_ applied to the moved gadgets and scaled by the proportion of moved gadgets compared to total gadgets in the original binary.

#### closest-only
When calculating each gadget's contribution to EQI, considers only the smallest offset. The gadget's other offsets are still considered when calculating other gadgets' contributions to EQI.

#### multiplicative
For each gadget, starts with a "quality" value of 1, and multiplies it by the complement of the penalty incurred by each offset (normalized to a number between 0 and 1). This causes EQI to decrease asymptotically as each gadget appears at a greater number of offsets.

#### additive
Adds the penalty for each gadget offset, so that for example a gadget with two offsets each shared with _n_ other gadgets would incur twice the EQI penalty a gadget with a single offset shared with _n_ other gadgets. The same as _count-poly_ with the default _order_ of 2.0 and _single_=false.

#### additive-with-ceiling
Adds the penalty for each gadget offset, but caps the contribution of each individual gadget to the EQI at 100 divided by the total number of gadgets. 