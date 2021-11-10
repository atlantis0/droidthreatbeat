# {Android Beat}

Welcome to {Beat}. This project has the ncessary files to push android threat data to a remote elastic server. To make
this beat Android compatable few things are done

1. Build for arm. `GOOS=android GOARCH=arm64 go build -ldflags "-s -w"`

2. This beat internally uses the default `net` package for remote connection. This packages resolves DNS by consulting /etc/resolve.conf. But
this file is not how DNS is resolved in the app. The issue is disscussed in detail here, https://github.com/coyove/goflyway/issues/126 &
https://stackoverflow.com/questions/38959067/dns-lookup-issue-when-running-my-go-app-in-termux 

3. Hence the workaround is to use relfection and modify internal structures like this one `resolvConf.dnsConfig.servers`

4. See workaround here, https://gist.github.com/cs8425/107e01a0652f1f1f6e033b5b68364b5e & https://github.com/mtibben/androiddnsfix

More details on the "why" can be found on this blog,

https://ismyapppwned.com/2021/02/26/elastic_beats_android/

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/atlantis0/androidthreatbeat`

## Getting Started with {Beat}

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with {Beat} and also install the
dependencies, run the following command:

```
make update
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push {Beat} in the git repository, run the following commands:

```
git remote set-url origin https://github.com/atlantis0/androidthreatbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for {Beat} run the command below. This will generate a binary
in the same directory with the name androidthreatbeat.

```
make
```


### Run

To run {Beat} with debugging output enabled, run:

```
./androidthreatbeat -c androidthreatbeat.yml -e -d "*"
```

On Device, use the following command 

```
/data/app/com.pwned.check-fyat6LhhXre1y_rhWG5xpg==/lib/arm64/androidthreatbeat -c /data/user/0/com.pwned.check/files/elastic/androidthreatbeat.yml --path.data /data/user/0/com.pwned.check/files/elastic/data --path.logs /data/user/0/com.pwned.check/files/elastic/logs --E THREAT_FILE_PATH=/data/user/0/com.pwned.check/files/threats.json -e -d run
```

Make sure to pass threat file path by using env variable. `-E THREAT_FILE_PATH=/path/to/threat.json`

### Test

To test {Beat}, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  {Beat} source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone {Beat} from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/atlantis0/androidthreatbeat
git clone https://github.com/atlantis0/androidthreatbeat ${GOPATH}/src/github.com/atlantis0/androidthreatbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
