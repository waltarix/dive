build:
	goreleaser --rm-dist --snapshot --skip-publish

test:
	CGO_ENABLED=0 go test -v ./...

fmt:
	go fmt -x ./...

clean:
	$(RM) -r dist

.PHONY: build test fmt clean
