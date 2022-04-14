build:
	go build -o flights

test: build
# unit tests
	go test -v ./...
# end to end tests
	echo '[["SFO", "ATL"]]' | ./flights || (echo "failed $$?"; exit 1)
	echo '[["ATL", "EWR"],["SFO", "ATL"]]' | ./flights  || (echo "failed $$?"; exit 1)
	echo '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' | ./flights || (echo "failed $$?"; exit 1)
	echo '[["EWR", "IND"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' | ./flights || (echo "failed $$?"; exit 0)
	echo '[["IND", "IND"]]' | ./flights || (echo "failed $$?"; exit 0)
	echo '[["IND", "SFO"],["SFO", "IND"]]' | ./flights || (echo "failed $$?"; exit 0)

clean:
	rm flights