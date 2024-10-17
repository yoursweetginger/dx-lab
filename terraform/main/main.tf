resource "decort_resgroup" "lab" {
  account_id = var.account_id
  gid = var.gid

  name = "lab"
}

resource "decort_k8s" "lab" {
  rg_id = decort_resgroup.lab.id
  k8sci_id = 194

  name = "lab"
  wg_name = "wg_lab"
  network_plugin = "flannel"

  masters {
    num = 1
    cpu = 2
    ram = 4096
    disk = 100
  }

  workers {
    name = "wg_lab"
    num = 1
    cpu = 4
    ram = 8192
    disk = 100
  }
}