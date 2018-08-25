# z-nomp

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/CoinMintTech/go-z-nomp.svg?branch=master)](https://travis-ci.org/CoinMintTech/go-z-nomp/)
[![Coverage Status](https://coveralls.io/repos/CoinMintTech/go-z-nomp/badge.svg?branch=master&service=github)](https://coveralls.io/github/CoinMintTech/go-z-nomp?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/CoinMintTech/go-z-nomp)](https://goreportcard.com/report/github.com/CoinMintTech/go-z-nomp)

[![Project](https://www.openhub.net/p/go-CoinMintTech-z-nomp/widgets/project_thin_badge.gif)][project]

Golang client for the [z-nomp][z-nomp] mining pool.

## Environment variables

- **ZNOMP_AUTH**: authentication method (default: `http`).
- **ZNOMP_AUTH_HTTP_PASSWORD**: password for HTTP basic authentication method.
- **ZNOMP_AUTH_HTTP_USER**: username for HTTP basic authentication method.
- **ZNOMP_ENDPOINT**: the z-nomp endpoint.

[license]:  https://raw.githubusercontent.com/CoinMintTech/go-z-nomp/master/LICENSE   "Apache License 2.0"
[project]:  https://www.openhub.net/p/go-CoinMintTech-z-nomp/    "OpenHub project page"
[z-nomp]:   https://github.com/z-classic/z-nomp "z-nomp"
