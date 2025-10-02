#! /bin/bash
echo -e "Start running the script..."

VERSION=1.0.0
GIT_COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date +"%Y-%m-%dT%H:%M:%S%:z")

go build -ldflags="\
-s -w \
-X 'main.Version=$VERSION' \
-X 'main.BuildTime=$BUILD_TIME' \
-X 'main.Commit=$GIT_COMMIT'\
" -o SimpleMap.exe

upx -9 SimpleMap.exe

echo -e "End running the script!"