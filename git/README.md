# tips

## submoduleの追加
```shell
# git submodule [--name <name>] <repository> [<path>]
git submodule brutus https://github.com/bradfitz/gitbrut.git bradfitz_bruteforce
[submodule "brutus"]
    path = bradfitz_bruteforce
    url = https://github.com/bradfitz/gitbrute
```
