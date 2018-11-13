gobuild = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"'
gotest = go test 

gopath:
	source gopath.env	

install-depends:
	go get -u gopkg.in/yaml.v2
	go get -u gopkg.in/korylprince/go-ad-auth.v2
	go get -u github.com/gin-gonic/gin

test-token:
	$(gotest) src/spinel/token.go src/spinel/token_test.go

spinel:
	$(gobuild) -o bin/spinel src/main.go

test: test-token
all: install-depends spinel

