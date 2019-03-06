# totem

[![CircleCI](https://circleci.com/gh/Tufin/totem.svg?style=shield)](https://circleci.com/gh/Tufin/totem)
[![Go Report Card](https://goreportcard.com/badge/github.com/tufin/totem)](https://goreportcard.com/report/github.com/tufin/totem)


Find invalid golang imports in a mono-repo project

#### Run
Set follow environment variables:
```bash
export TOTEM_COMMON_IMPORTS=github.com/tufin/totem/common
export TOTEM_PACKAGE=github.com/tufin/orca/
export TOTEM_PATH=/Users/israel/view/go/src/github.com/tufin/orca
```
This will run *on root folder that contains multiple micro-services*.
Set below environment variable if you want to run _totem_ a specific service:
```bash
export TOTEM_SERVICE=ceribro
```
Running _totem_:
```bash
totem
```
Exit code 1 if there are invalid imports, else 0
