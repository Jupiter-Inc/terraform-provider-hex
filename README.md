# Terraform Hex Provider

Converts a normal string to a hex string.

Based on the [Writing Custom Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html) tutorial and the [Random Provider](https://github.com/hashicorp/terraform-provider-random)

## Example Usage

```hcl
resource "random_password" "key" {
  length = 32
}

resource "hex_string" "key-hex" {
  data = random_password.password.result
}

resource "helm_release" "jupiter" {
  name       = "jupiter"
  chart      = "jupiter"
  version    = "1.0.0"

  set_sensitive {
    name  = "key"
    value = hex_string.key_hex.result
  }
}
```

## Specific Use Case

`imgproxy` requires it's key and salt to be hex-encoded, but there only exist built-in functions for [`base64`](https://www.terraform.io/docs/configuration/functions/base64encode.html), [`json`](https://www.terraform.io/docs/configuration/functions/jsonencode.html), [`url`](https://www.terraform.io/docs/configuration/functions/urlencode.html), and [`yaml`](https://www.terraform.io/docs/configuration/functions/yamlencode.html) encoding.

## Building/Testing

1. Make sure you still have `go` installed (`brew install go`)
1. `VERSION=X.Y.Z make build` (don't include the `v` before the version)
1. Follow the instructions [here](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations) to install the plugin
1. `terraform init`
1. `terraform plan`
1. `terraform apply`
1. Repeat steps 3-7 as you make changes.

## Publishing

1. Make sure you have `hub` installed (`brew install hub`)
1. `VERSION=X.Y.Z make publish` (don't include the `v` before the version)
