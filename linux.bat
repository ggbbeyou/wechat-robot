@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -gcflags "-N -l" -ldflags "-s -w" -o linux_wechat_robot_amd64