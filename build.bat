set PATH=C:\mingw-w64\x86_64-7.2.0-posix-seh-rt_v5-rev1\mingw64\bin;%GOROOT%\bin;%PATH%
set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-physx\=\%
set GOPATH=%BASEDIR%
set GOBIN=%CURDIR%\bin
go install ./...
pause