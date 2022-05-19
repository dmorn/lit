{
  description = "Literature Crawling Tool";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs-fmt.url = "github:nix-community/nixpkgs-fmt";
  };

  outputs = { self, nixpkgs, flake-utils, nixpkgs-fmt }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        formatter = pkgs.nixpkgs-fmt;
        packages = {
          lit = pkgs.buildGoModule {
            name = "lit";
            src = ./.;
            vendorSha256 = "sha256-iejiOwUqdPCv1up3ao0uyxeT5kenUfqA8kKYXoENNAM=";
          };
        };

        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
          ];
        };
      });
}
