export GOPROXY=direct

sudo apt-get update
sudo apt-get install gcc-mingw-w64-i686 gcc-multilib

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./linux_amd64_PBEBrute ./
tar -czvf linux_amd64_PBEBrute.tar.gz linux_amd64_PBEBrute

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./linux_386_PBEBrute ./
tar -czvf linux_386_PBEBrute.tar.gz linux_386_PBEBrute

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./windows_386_PBEBrute ./
tar -czvf windows_386_PBEBrute.tar.gz windows_386_PBEBrute

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./windows_amd64_PBEBrute ./
tar -czvf windows_amd64_PBEBrute.tar.gz windows_amd64_PBEBrute

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./darwin_amd64_PBEBrute ./
tar -czvf darwin_amd64_PBEBrute.tar.gz darwin_amd64_PBEBrute