{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    cmake
    jq
  ];

  packages = with pkgs; [
    git
    neovim
    httpie
    pre-commit
    go
  ];

  GIT_EDITOR = "${pkgs.neovim}/bin/nvim";

  shellHook = ''

  '';

}