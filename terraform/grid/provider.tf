terraform {
  required_providers {
    decort = {
      version = "4.6.3"
      source  = "basis/decort/decort"
    }
  }
}

provider "decort" {
  authenticator        = "decs3o"
  controller_url       = "https://mr4.digitalenergy.online"
  oauth2_url           = "https://sso.digitalenergy.online"
  allow_unverified_ssl = true
}
