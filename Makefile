build:
	go build -o flights

test: build
	go test -v ./...
	./flights -jsonFile testdata/input.json

clean:
	rm flights