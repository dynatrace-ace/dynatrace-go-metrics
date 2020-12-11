build_windows:
	echo "--- Compiling Windows Binaries ---"
	GOOS=windows GOARCH=386 go build -o bin/windows/dt-send-metrics-x86.exe
	GOOS=windows GOARCH=amd64 go build -o bin/windows/dt-send-metrics-x64.exe

build_linux:
	echo "--- Compiling Linux Binaries ---"
	GOOS=linux GOARCH=386 go build -o bin/linux/dt-send-metrics-x86
	GOOS=linux GOARCH=amd64 go build -o bin/linux/dt-send-metrics-x64

build_macos:
	echo "--- Compiling Darwin (MacOS) Binaries ---"
	GOOS=darwin GOARCH=386 go build -o bin/macos/dt-send-metrics-x86
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/dt-send-metrics-x64

build: build_windows build_linux build_macos
