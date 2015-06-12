@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)

go build -a -o Release/cef2go.exe main.go
@echo exit code = %ERRORLEVEL%