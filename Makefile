build:
	go build -o roll

clean:
	rm -f roll

install:
	mv roll /usr/local/bin
