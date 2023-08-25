build:
	go build -o roll

clean:
	rm -f roll

install: build
	mv roll /usr/local/bin

uninstall:
	rm -rf /usr/local/bin/roll
