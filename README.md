# terraform-provider-dhall

This providers brings the ability for terraform to use dhall as an external datasource.

Result from the dhall evaluation is provided as a JSON-encoded string.

## Usage

```terraform
# resources.tf
resource "aws_iam_policy" "assets-access" {
  name  = "assets-access"
  description = "A policy allowing access to assets."

  policy = data.dhall.iam_policy_assets_access.result
}

data "dhall" "iam_policy_assets_access" {
    entrypoint = "(./assets_access.dhall) { region = \"${var.region}\", accountId = \"${var.account-id}\" }"
}

# terraform.tf
terraform {
  required_providers {
    dhall = {
      source  = "awakesecurity/dhall"
      version = "0.0.1"
    }
  }
}
```

```dhall
-- assets_access.dhall
let predicate =
      https://raw.githubusercontent.com/mjgpy3/iam-dhall/20bcc9c507d353fb3736a633280239a922b91aa6/policy.dhall

let policy =
      https://raw.githubusercontent.com/mjgpy3/iam-dhall/20bcc9c507d353fb3736a633280239a922b91aa6/output.dhall

let Aws
    : Type
    = { accountId : Text, region : Text }

let listGetBucketAccess =
      \(bucket : Text) ->
        [     predicate.serviceAllow
                predicate.Service.S3
                [ "ListBucket" ]
                [ bucket ]
          //  { sid = "ListObjects" }
        ,     predicate.serviceAllow
                predicate.Service.S3
                [ "GetObject" ]
                [ "${bucket}/*" ]
          //  { sid = "GetObject" }
        ]

let assetsAccess =
      \(aws : Aws) ->
        policy
          aws
          (   {- merge access to public-assets and static-assets -}
              listGetBucketAccess "public-assets"
            # listGetBucketAccess "static-assets"
          )

in  assetsAccess

```
