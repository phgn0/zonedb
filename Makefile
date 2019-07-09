.PHONY: test update metadata/*.json

install:
	go install ./cmd/zonedb

test:
	go run cmd/zonedb/main.go
	go test ./...

zones.go: zones.txt metadata/*.json internal/* internal/*/*
	go generate

update:
	go run cmd/zonedb/main.go -update -w -c 100
	make zones.go
	make test

normalize:
	go run cmd/zonedb/main.go -w
	make zones.go

major_version=$(shell cat VERSION)
number_of_commits=$(shell git rev-list HEAD --count)
tag_version=v$(major_version).$(number_of_commits)

tag-version: .git/refs/heads/master
	git tag $(tag_version) $(shell cat .git/refs/heads/master)
	git push --tags
