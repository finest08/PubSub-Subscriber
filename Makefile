# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf generate
	buf build
	cd proto && buf push

.PHONY: run
run:
	dapr run \
		--app-id pubsubsubscriber \
		--app-port 8082 \
		--app-protocol grpc \
		--config ./.dapr/config.yaml \
		--components-path ./.dapr/components \
		go run .

.PHONY: kill
kill:
	lsof -P -i TCP -s TCP:LISTEN | grep 8082 | awk '{print $2}' | { read pid; kill -9 ${pid}; }
	lsof -P -i TCP -s TCP:LISTEN | grep 9092 | awk '{print $2}' | { read pid; kill -9 ${pid}; }
.PHONY: test
test:
	go test -v ./handler/...
