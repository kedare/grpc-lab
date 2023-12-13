build:
	echo "Building gomonitor (Go)"
	(cd gomonitor; go build -o ../bin/gomonitor cmd/gomonitor/main.go)
	echo "Building Blazorinfra (C#)"
	(cd blazinfra; dotnet publish --sc -o ../bin/)
	echo "Done"

clean:
	rm -rf bin
	mkdir bin

start:
	./bin/gomonitor &
	(cd ./bin; ./BlazInfra &)

stop:
	killall gomonitor
	killall BlazInfra