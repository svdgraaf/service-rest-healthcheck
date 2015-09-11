# service-rest-healthcheck
Check the status of Windows service through REST calls

# Why?
We wanted to check through ELB if a windows server was still running a specific service. Now, we can just point the ELB healthcheck to a specific service, and the ELB will mark the Widnows machine as unhealthy whenever the Service fails to run or is crashed.

# Installation
Download one of the build for your platform.

# Usage
Run the service as administrator, the service will run on port 8484. Use these endpoints:

- `/ping/`: You can use this check whether the service is actually running. Will return a `200 OK` with `pong`
- `/service/[service]/status`: Check for the service status. Will return a `200 OK` when the service is running with the service state. A `404` when the service is not found, and a `50x` when the service is not found, or if there was another reason why the service could not be checked (see the logs).

# Compile
Get all the dependencies, `go get -d`, and run `make`. The build will be in the build directory.
