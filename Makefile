dev:
	go build -gcflags "-N -l" -i -x  &&env=dev ./go-web
