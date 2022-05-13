# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	buf mod update
	buf generate
	buf build
