{ lib, buildGoModule, nix-gitignore }:

buildGoModule rec {
  pname = "terraform-provider-dhall";
  version = "0.0.1";

  src = nix-gitignore.gitignoreSource [] ./.;

  vendorSha256 = "0m11cpis171j9aicw0c66y4m1ckg41gjknj86qvblh57ai96gc1n";

  postInstall = "mv $out/bin/terraform-provider-dhall{,_v${version}}";
}
