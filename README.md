# service-rest-healthcheck
Check the status of Windows service through REST calls

# Why?
We wanted to check through ELB if a windows server was still running a specific service. Now, we can just point the ELB healthcheck to a specific service, and the ELB will mark the Widnows machine as unhealthy whenever the Service fails to run or is crashed.

# Installation
Download one of the build for your platform.

# Run the

# Compile
Get all the dependencies, `go get -d`, and run `make`. The build will be in the build directory.
