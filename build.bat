@echo off

for %%i in (*.exe) do (
    Wmic Process Where "Name Like '%%%%i%%%'" Call Terminate >Nul 2>&1
)

rsrc -manifest main.manifest -ico favicon.ico -o main.syso 2>Nul

if %errorlevel% NEQ 0 (
    echo "Please download rsrc: https://github.com/akavel/rsrc"
    pause
    exit
)

go build -ldflags -H=windowsgui -tags walk_use_cgo

if %errorlevel% NEQ 0 (
    pause
)