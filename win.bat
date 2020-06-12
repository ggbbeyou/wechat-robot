@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -gcflags "-N -l" -ldflags "-s -w -H windowsgui" -o windows_wechat_robot_amd64.exe