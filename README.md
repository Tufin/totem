# totem

Find invalid golang imports in a mono-repo project

#### Run
Set follow environment variables:
```bash
export COMMON_IMPORTS=github.com/tufin/totem/common
export PACKAGE=github.com/tufin/orca/
export PATH=/Users/israel/view/go/src/github.com/tufin/orca
```
Running totem:
```bash
totem
```
Exit code 1 if there are invalid imports, else 0
