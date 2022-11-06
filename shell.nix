{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  buildInputs = [
    delve
    go-outline
    go_1_19
    golangci-lint
    gopkgs
    gopls
    gotools
    just
    nixpkgs-fmt
  ];

  shellHook = ''
    # Figure out the workspace root, which is:
    # - pwd when this hook fires under direct nix-shell invocations
    # - dirname of the IN_LORRI_SHELL file when fired by lorri invocations
    shell_nix="''${IN_LORRI_SHELL:-$(pwd)/shell.nix}"
    workspace_root=$(dirname "$shell_nix")

    # The path to this repository
    export WORKSPACE_ROOT="$workspace_root"
    # Path to a folder that can hold build caches, etc. without them being
    # present in the main repo. When run as a daemon, Lorri evaluates this
    # shellHook without setting $USER or $HOME so it's ~impossible to put
    # this in $HOME. Instead, we make it adjacent, as
    # `.pipe-$repositoryFolderName`, so that multiple checkouts of this repo in
    # the same folder will still each have their own caches.
    export TOOLCHAIN_ROOT="$workspace_root/.toolchain"

    # Use a GOPATH local to this workspace so we get a portable
    # go environment that won't ever conflict with other repos or
    # other coding happening elsewhere.
    #
    # We put the $GOPATH/$GOCACHE/$GOENV in $TOOLCHAIN_ROOT,
    # and ensure that the GOPATH's bin dir is on our PATH so tools
    # can be installed with `go install`.
    # 
    # Any tools installed explicitly with `go install` will take precedence
    # over versions installed by Nix due to the ordering here.
    export GOROOT=
    export GOCACHE="$TOOLCHAIN_ROOT/go/cache"
    export GOENV="$TOOLCHAIN_ROOT/go/env"
    export GOPATH="$TOOLCHAIN_ROOT/go/path"
    export GOMODCACHE="$GOPATH/pkg/mod"
    export PATH=$(go env GOPATH)/bin:$PATH
  '';

  # Need to disable fortify hardening because GCC is not built with -oO,
  # which means that if CGO_ENABLED=1 (which it is by default) then the golang
  # debugger fails.
  # see https://github.com/NixOS/nixpkgs/pull/12895/files
  hardeningDisable = [ "fortify" ];
}
