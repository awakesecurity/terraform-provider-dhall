resource "aws_iam_policy" "assets-access" {
  name  = "assets-access"
  description = "A policy allowing access to assets."

  policy = data.dhall.iam_policy_assets_access.result
}

data "dhall" "iam_policy_assets_access" {
    entrypoint = "(./assets_access.dhall) { region = \"${var.region}\", accountId = \"${var.account-id}\" }"
}
