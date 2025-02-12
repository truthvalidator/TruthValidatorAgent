VERSION=$(shell git describe --tags --always)
GITHASH=$(shell git rev-parse HEAD 2>/dev/null)
BUILDAT=$(shell date +%FT%T%z)
API_PROTO_FILES=$(shell find api/* -name *.proto)
LDFLAGS="-s -w -X TruthValidatorSentientNet-tgbot/cmd.GitHash=${GITHASH} -X TruthValidatorSentientNet-tgbot/cmd.BuildAt=${BUILDAT} -X TruthValidatorSentientNet-tgbot/cmd.Version=${VERSION}"


buildlocal:
	mkdir -p bin
	go build -v -o bin/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp .env bin/
	cp -r app.example.toml bin/app.toml

builddarwin:
	rm -rf release/${VERSION}/darwin
	GOOS=darwin GOARCH=amd64 go build -v -o release/${VERSION}/darwin/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/darwin
	cp -r app.example.toml release/${VERSION}/darwin
	cd release/${VERSION}/darwin/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_darwin_amd64.tar.gz ./* && cd ../../

builddarwinarm:
	rm -rf release/${VERSION}/darwin-arm
	GOOS=darwin GOARCH=arm64 go build -v -o release/${VERSION}/darwin-arm/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/darwin-arm
	cp -r app.example.toml release/${VERSION}/darwin-arm/app.toml
	cd release/${VERSION}/darwin-arm/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_darwin_arm64.tar.gz ./* && cd ../../

buildlinux:
	rm -rf release/${VERSION}/linux
	GOOS=linux GOARCH=amd64 go build -v -o release/${VERSION}/linux/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/linux
	cp -r app.example.toml release/${VERSION}/linux/app.toml
	cd release/${VERSION}/linux/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_linux_amd64.tar.gz .env app.toml TruthValidatorSentientNet-tgbot && cd ../../

builddockeramd:
	rm -rf release/${VERSION}/dockeramd
	GOOS=linux GOARCH=amd64 go build -v -o release/${VERSION}/dockeramd/server/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/dockeramd/server
	cp docker/dockerfile release/${VERSION}/dockeramd/dockerfile
	cp docker/docker-compose.yml release/${VERSION}/dockeramd/docker-compose.yml
	cp docker/README.md release/${VERSION}/dockeramd/readme.md
	cp -r docker/mysql release/${VERSION}/dockeramd/mysql
	cd release/${VERSION}/dockeramd/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_docker_amd64.tar.gz ./* && cd ../../


builddockerarm:
	rm -rf release/${VERSION}/dockerarm
	GOOS=linux GOARCH=arm64 go build -v -o release/${VERSION}/dockerarm/server/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/dockerarm/server
	cp docker/dockerfile release/${VERSION}/dockerarm/dockerfile
	cp docker/docker-compose.yml release/${VERSION}/dockerarm/docker-compose.yml
	cp docker/README.md release/${VERSION}/dockerarm/readme.md
	cp -r docker/mysql release/${VERSION}/dockerarm/mysql
	cd release/${VERSION}/dockerarm/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_docker_arm64.tar.gz ./* && cd ../../


buildwin:
	rm -rf release/${VERSION}/windows
	GOOS=windows GOARCH=amd64 go build -v -o release/${VERSION}/windows/TruthValidatorSentientNet-tgbot.exe -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/windows
	cp -r app.example.toml release/${VERSION}/windows
	cd release/${VERSION}/windows/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_windows_amd64.tar.gz ./* && cd ../../

buildlinuxarm:
	rm -rf release/${VERSION}/linux-arm
	GOOS=linux GOARCH=arm64 go build -v -o release/${VERSION}/linux-arm/TruthValidatorSentientNet-tgbot -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/linux-arm
	cp -r app.example.toml release/${VERSION}/linux-arm
	cd release/${VERSION}/linux-arm/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_linux_arm64.tar.gz ./* && cd ../../

buildwinarm:
	rm -rf release/${VERSION}/windows-arm
	GOOS=windows GOARCH=arm64 go build -v -o release/${VERSION}/windows-arm/TruthValidatorSentientNet-tgbot.exe -ldflags ${LDFLAGS}
	cp  .env release/${VERSION}/windows-arm
	cp -r app.example.toml release/${VERSION}/windows-arm
	cd release/${VERSION}/windows-arm/ && tar -zcvf ../TruthValidatorSentientNet-tgbot_ce_${VERSION}_windows_arm64.tar.gz ./* && cd ../../

# Build all platforms
# 一键编译所有平台
buildall: builddarwin builddarwinarm buildlinux buildwin buildlinuxarm buildwinarm builddockerarm builddockeramd


