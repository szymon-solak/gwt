{
  description = "gwt - Simplified git worktree management";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      packages.default = pkgs.buildGoModule {
        pname = "gwt";
        version = "0.1.0";

        src = ./.;

        vendorHash = "sha256-9jK3jKbFp+5WSQfMbNzwIB55bC5KScZOaFHItffTF00=";

        subPackages = [ "." ];

        ldflags = [ "-s" "-w" ];

        meta = with pkgs.lib; {
          description = "Simplified git worktree management";
          homepage = "https://github.com/szymon-solak/gwt";
          license = licenses.mit;
          maintainers = [ ];
        };
      };

      devShells.default = pkgs.mkShell {
        name = "gwt";

        buildInputs = [
          pkgs.go
          pkgs.gopls
        ];
      };
    });
}
