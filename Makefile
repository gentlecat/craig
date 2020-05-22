fmt :
	go fmt -x ./...

test : fmt
	go test ./... -bench .  -coverprofile=coverage.out

coverage : test
	go tool cover -html=coverage.out
