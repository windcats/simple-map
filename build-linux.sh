#! /bin/bash
echo -e "Start running the script..."

VERSION=1.0.0
GIT_COMMIT=$(git rev-parse --short HEAD)
BUILD_TIME=$(date +"%Y-%m-%dT%H:%M:%S%:z")

CGO_ENABLED=0 GOOS=linux go build -ldflags="\
-s -w \
-X 'main.Version=$VERSION' \
-X 'main.BuildTime=$BUILD_TIME' \
-X 'main.Commit=$GIT_COMMIT'\
" -o SimpleMap

upx -9 SimpleMap

echo -e "End running the script!"