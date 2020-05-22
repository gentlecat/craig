fmt :
	go fmt -x ./...

test : fmt
	go test ./... -bench .
