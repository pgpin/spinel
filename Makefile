gobuild = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"'
gotest = go test 
version = 0.1.0

install-depends:
	go get gopkg.in/yaml.v2
	go get gopkg.in/korylprince/go-ad-auth.v2
	go get github.com/gin-gonic/gin
	go get github.com/jbmcgill/go-throttle
	go get github.com/adam-hanna/randomstrings

test-token:
	$(gotest) src/spinel/token.go src/spinel/token_test.go

test-cidr:
	$(gotest) src/spinel/cidr.go src/spinel/cidr_test.go

test-config:
	$(gotest) src/spinel/config.go src/spinel/config_test.go

spinel:
	$(gobuild) -o bin/spinel src/main.go

dist:
	mkdir -p releases build/etc/init.d build/usr/local/bin build/usr/local/share/spinel/tmpl
	cp -p bin/spinel build/usr/local/bin/spinel
	strip build/usr/local/bin/spinel
	cp -p tmpl/* build/usr/local/share/spinel/tmpl/
	cp -p src/init.d/spinel build/etc/init.d/spinel
	cp -p example-config.yaml build/etc/spinel.yaml
	cd build && fpm --post-install ../src/init.d/postinstall.sh --force --package ../releases/ -n spinel --version $(version)-$(rc) -s dir -t rpm --rpm-user nginx ./etc ./usr
	cd build && fpm --post-install ../src/init.d/postinstall.sh --force --package ../releases/ -n spinel --version $(version)-$(rc) -s dir -t deb --deb-user nginx ./etc ./usr
	cd build && tar -zcf ../releases/spinel-${version}-$(rc).tar.gz ./

test: test-token test-cidr test-config

all: install-depends spinel

