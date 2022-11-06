{ pkgs ? import <nixpkgs> {} }:
with pkgs;
mkShell {
  # Need to disable fortify hardening because GCC is not built with -oO,
  # which means that if CGO_ENABLED=1 (which it is by default) then the golang
  # debugger fails.
  # see https://github.com/NixOS/nixpkgs/pull/12895/files
  hardeningDisable = [ "fortify" ];

  buildInputs = [
    delve
    go-outline
    go_1_19
    golangci-lint
    gopkgs
    gopls
    gotools
    just
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
    export INSTALL_ROOT=$(dirname "$workspace_root")/.workspace-$(basename "$workspace_root")

    # Use a GOPATH local to this workspace so we get a portable
    # go environment that won't ever conflict with other repos or
    # other coding happening elsewhere.
    #
    # We put the $GOPATH/$GOCACHE/$GOENV in $INSTALL_ROOT, which is
    # adjacent to this repository root, to fix problems with vscode-go and
    # gopls -- they don't like it when the $GOPATH is inside the repository
    # being worked on.
    #
    # We also ensure the GOPATH's bin dir is on our PATH so tools
    # installed with `go install` work.
    #
    # Long term, tools installed with `go install` can also be
    # included directly as a nix package. Presently we're not
    # sure which way is best-- perhaps both ways are fine :).
    # 
    # Any tools installed explicitly with `go install` will take precedence
    # over versions installed by Nix due to the ordering here. VSCode is
    # configured to act this way already, explicitly, but this makes it so that
    # the shell/environment in general has the same behavior. The downside to
    # this is that if a user `go install`s a tool that we later update with
    # nix, they will remain using the version that they `go install`ed. 
    export GOCACHE="$INSTALL_ROOT/go/cache"
    export GOENV="$INSTALL_ROOT/go/env"
    export GOPATH="$INSTALL_ROOT/go/path"
    export GOMODCACHE="$GOPATH/pkg/mod"
    export GOROOT=
    export PATH=$(go env GOPATH)/bin:$PATH
  '';
}
