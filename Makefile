gobuild = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"'
gotest = go test 

install-depends:
	go get gopkg.in/yaml.v2
	go get gopkg.in/korylprince/go-ad-auth.v2
	go get github.com/gin-gonic/gin
	go get github.com/jbmcgill/go-throttle

test-token:
	$(gotest) src/spinel/token.go src/spinel/token_test.go

test-cidr:
	$(gotest) src/spinel/cidr.go src/spinel/cidr_test.go

test-config:
	$(gotest) src/spinel/config.go src/spinel/config_test.go

spinel:
	$(gobuild) -o bin/spinel src/main.go

test: test-token test-cidr test-config

all: install-depends spinel

