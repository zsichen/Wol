# Wol
## build
```
#for linux arm
GOOS=linux CGO_ENABLE=0 GOARCH=arm go build

#for linux arm64
GOOS=linux CGO_ENABLE=0 GOARCH=arm64 go build

#for linux amd64
GOOS=linux GOARCH=amd64 go build

````

## Run
```
./wol -m BE:83:85:11:57:E8 -a 192.168.2.255
```