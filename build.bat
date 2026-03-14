@echo off
cd /d "%~dp0"
cd web && call npm run build:prod
if errorlevel 1 exit /b 1
cd ..
set CGO_ENABLED=0
go build -o k8s-admin.exe .
if errorlevel 1 exit /b 1
echo Build succeeded: k8s-admin.exe
