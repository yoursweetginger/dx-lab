terraform {
  required_providers {
    decort = {
      version = "4.2.2"
      source  = "digitalenergy.online/decort/decort"
    }
  }
}

provider "decort" {
   authenticator = "oauth2"
   controller_url = "https://mr4.digitalenergy.online"
   oauth2_url = "https://sso.digitalenergy.online"
   allow_unverified_ssl = true
}

resource "decort_kvmvm" "go-sample" {
   name = "go-sample"
   rg_id = 2071
   driver = "KVM_X86"
   cpu = 2
   ram = 2048 
   boot_disk_size  = 20
   image_id = 388
   started = true
   enabled = true
   
   network {
      net_type = "EXTNET"
      net_id = 3
   }
}

output ip_address {
  value = tolist(decort_kvmvm.go-sample.network)[0].ip_address
}

output login {
  value = decort_kvmvm.go-sample.os_users[0].login
}

output password {
  value = decort_kvmvm.go-sample.os_users[0].password
  sensitive = true
}
