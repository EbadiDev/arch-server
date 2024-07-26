# Arch-Server

Arch Server For Testing 

<!-- <p align="center"><img src="https://avatars.githubusercontent.com/u/91626055?v=4" width="128" /></p> -->

<!-- <div align="center"> -->

<!-- [![](https://img.shields.io/github/downloads/EbadiDev/Arch-Server/total.svg?style=flat-square)](https://github.com/EbadiDev/Arch-Server/releases) -->
<!-- [![](https://img.shields.io/github/v/release/EbadiDev/Arch-Server?style=flat-square)](https://github.com/EbadiDev/Arch-Server/releases)
[![docker](https://img.shields.io/docker/v/aikocute/Arch-Server?label=Docker%20image&sort=semver)](https://hub.docker.com/r/aikocute/Arch-Server)
[![Go-Report](https://goreportcard.com/badge/github.com/EbadiDev/Arch-Server?style=flat-square)](https://goreportcard.com/report/github.com/EbadiDev/Arch-Server) -->

</div>

# Description of Arch-Server

<!-- Arch-Server Supports for AikoPanel -->

An Xray-based and SingBox-based back-end framework, supporting V2ay, Trojan, Shadowsocks protocols, extremely easily extensible and supporting multi-panel connection。

If you like this project, you can click the star + view in the upper right corner to track the progress of this project.

## Disclaimer

This project is for my personal learning, development and maintenance only, I do not guarantee the availability and I am not responsible for any consequences resulting from using this software.

## Featured

- Open source `This version depends on the happy mood`
- Supports multiple protocols V2ray, Trojan, Shadowsocks.
- Supports new features like Vless and XTLS.
- Supports single connection to multiple boards and nodes without rebooting.
- Online IP support is limited
- Support node port level, user level rate limit.
- Simple and clear configuration.
- Modify the configuration to automatically restart the instance.
- Easy to compile and upgrade, can quickly update core version, support new Xray-core features.
- Support UDP and many other functions

## Featured

| Feature                              | v2ray | trojan | shadowsocks | hysteria |
| ------------------------------------ | ----- | ------ | ----------- | -------- |
| Automatically apply tls certificates | √     | √      | √           | √        |
| Automatically renew tls certificates | √     | √      | √           | √        |
| Online user statistics               | √     | √      | √           | √        |
| Audit rules                          | √     | √      | √           | √        |
| Custom DNS                           | √     | √      | √           | √        |
| Limit online IP numbers              | √     | √      | √           | √        |
| Connection limit                     | √     | √      | √           | √        |
| Cross-node IP number limit           | √     | √      | √           | √        |
| Limit speed according to users       | √     | √      | √           | √        |
| Dynamic speed limit (untested)       | √     | √      | √           | √        |

## User interface support

| Panel       | v2ray | trojan | shadowsocks | hysteria | hysteria2 |
| ----------- | ----- | ------ | ----------- | -------- | --------- |
| None   | √     | √      | √           | √        | √         |

## Software installation - release

<!-- ```
wget --no-check-certificate -O install.sh https://raw.githubusercontent.com/AikoPanel/Arch-Server-Script/master/install_v2.sh && bash install.sh
``` -->

### Docker installation

<!-- ```
docker pull aikocute/Arch-Server:latest && docker run --restart=always --name Arch-Server -d -v ${PATCH_TO_CONFIG}/aiko.json:/etc/Arch-Server/aiko.json --network=host aikocute/Arch-Server:latest
``` -->

### Docker-compose installation

<!-- Step 1 : Create Config File `aiko.json` in `/etc/Arch-Server/aiko.json`

```
mkdir -p /etc/Arch-Server/config && cd /etc/Arch-Server/config && wget https://raw.githubusercontent.com/EbadiDev/Arch-Server-Script/master/aiko.json
``` -->
<!-- 
Step 2 : Create `docker-compose.yml` in `/etc/Arch-Server/docker-compose.yml`

```
mkdir -p /etc/Arch-Server && cd /etc/Arch-Server && wget https://raw.githubusercontent.com/EbadiDev/Arch-Server-Script/master/docker-compose.yml
```

Step 3 : Run `docker-compose up -d` in `/etc/Arch-Server/`

```
cd /etc/Arch-Server/ && docker-compose up -d
``` -->

## Telgram

<!-- [Tele Aiko](https://t.me/Tele_Aiko) -->

<!-- ## Stargazers over time

[![Stargazers over time](https://starchart.cc/EbadiDev/Arch-Server.svg)](https://starchart.cc/EbadiDev/Arch-Server) -->
