# skyinfoblox - Go library for the Infoblox appliance

This is the GoLang API wrapper for Infoblox. This is currently used for building terraform provider for the same.

Building the cli binary
```
make all

```

Run Unit tests
```
make test

```

This will give you skyinfoblox-cli file which you can use to interact with InfoBlox API.

```
$ ./skyinfoblox-cli
  -debug
    	Debug output. Default:false
  -password string
    	Authentication password (Env: IBX_PASSWORD)
  -port int
    	Infoblox API server port. Default:443 (default 443)
  -server string
    	Infoblox API server hostname or address (default "localhost")
  -username string
    	Authentication username (Env: IBX_USERNAME)
  Commands:
    records-list

```

```
$ ./skyinfoblox-cli -server=https://serverhostnameOrIP  -username=admin -password=password records-list
{record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLnNreS5vdnAubnAsdGVzdC55b3JnLDEwLjEwLjEwLjEw:yorg.test.np.ovp.sky.com/default 10.10.10.10 yorg.test.np.ovp.sky.com default}
{record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnMnRlc3QsMTAuMTAuMTAuMg:craig2test.test-ovp.bskyb.com/default 10.10.10.2 craig2test.test-ovp.bskyb.com default}
{record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default 10.10.10.4 craig4test.test-ovp.bskyb.com default}

```

Development

during your development, you can run the cli with following command.
```
go run cli/*.go -server=https://infobloxserver.com  -username=admin -password=pass records-list

```