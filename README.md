# [WIP] jumppointsearch
toy repo, currently WIP, where I want to implement jump point search in golang just for the hell of it.

## background / links
* https://github.com/golang/go/wiki/SliceTricks
* https://stackoverflow.com/a/23532104
* https://golang.google.cn/pkg/container/heap
* https://movingai.com/benchmarks/grids.html
* https://harablog.wordpress.com/2011/09/07/jump-point-search/
* https://github.com/rangercyh/jps
* https://ojs.aaai.org/index.php/AAAI/article/view/7994

## development
this is a pure golang project, but you can get a fully working setup with Nix.
I also included a config file for VSCode.

Start by entering a shell environment with `nix-shell` or `nix develop`.
Optionally, you can use `direnv` (optionally with `lorri`) to load the nix
environment as soon as you enter the repo.

```bash
# development shell using nix-shell:
nix-shell

# development shell using nix flakes:
nix develop

# to get a development shell using direnv, just allow it
direnv allow

# using lorri, force reload
lorri watch --once
```

then you should have all the deps set up appropriately. Now get to work:

```bash
# open vscode
code .

# run the tests
just test

## run the linter
just lint
```
