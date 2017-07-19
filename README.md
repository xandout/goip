### GoIP

_This is my first attempt at anything useful in go_

GoIP is a simple utility to show allocated EIPs that are not in use and could be costing money.


## Usage

```
mt-mbp:goip mturner$ goip -h
Usage of goip:
  -region string
    	Region (default "us-east-1")
```

## Output

```
mt-mbp:goip mturner$ goip
IP 10.43.65.21 in STANDARD scope is unused
IP 172.16.6.4 in STANDARD scope is unused
IP 10.0.87.142 in STANDARD scope is unused
IP 192.168.2.55 in VPC scope is unused. EIP Allocation ID: eipalloc-5fa2806f
IP 192.168.4.71 in VPC scope is unused. EIP Allocation ID: eipalloc-df49beec
```


## Build

```
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/awserr
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/ec2

go get github.com/xandout/goip

cd $GOPATH/src/github.com/xandout/goip

go install

```