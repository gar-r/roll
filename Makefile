build:
	go build -o roll

clean:
	rm -f roll

install: build
	mv roll /usr/local/bin
