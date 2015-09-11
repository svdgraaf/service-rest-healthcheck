# service-rest-healthcheck
Check the status of Windows services through REST calls

# Why?
We wanted to check through ELB if a windows server was still running a specific service. Now, we can just point the ELB healthcheck to a specific service, and the ELB will mark the Windows machine as unhealthy whenever the Service fails to run or has crashed.

# Installation
Download a [release](https://github.com/svdgraaf/service-rest-healthcheck/releases), and install the executable as a service, eg with [NSSM](http://nssm.cc/):
`nssm install ServiceRestHealthcheck C:\service-rest-healthcheck-winx86-0.0.2.exe`

# Usage
The healthcheck service will run on port 8484. Use these endpoints:

- `/ping/`: You can use this check whether the service is actually running. Will return a `200 OK` with `pong`
- `/service/[service]/status`: Check for the service status. Will return a `200 OK` when the service is running with the service state. A `404` when the service is not found, and a `50x` when the service is not found, or if there was another reason why the service could not be checked (see the logs).

```bash
$ curl -i http://1.2.3.4:8484/service/bthserv/status/
HTTP/1.1 500 Internal Server Error
Date: Fri, 11 Sep 2015 19:20:21 GMT
Content-Length: 22
Content-Type: text/plain; charset=utf-8

Service is not running

$ curl -i http://1.2.3.4:8484/service/foobar/status/
HTTP/1.1 404 Not Found
Date: Fri, 11 Sep 2015 19:20:15 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

$ curl -i http://1.2.3.4:8484/service/power/status/
HTTP/1.1 200 OK
Date: Fri, 11 Sep 2015 19:19:31 GMT
Content-Length: 29
Content-Type: text/plain; charset=utf-8

Service is running: {4 4 0 0}
```

# Compile
Get all the dependencies, `go get -d`, and run `make`. The build will be in the `bin` directory.
