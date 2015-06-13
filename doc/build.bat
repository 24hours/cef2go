@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)

set C_INCLUDE_PATH=%GOPATH%\src\github.com\24hours\chrome\
set LIBRARY_PATH=%GOPATH%\Release\

go build -a -o Release/cef2go.exe main.go
@echo exit code = %ERRORLEVEL%