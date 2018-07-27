variable "do_token" {}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "elasticsearch" {
  image  = "ubuntu-14-04-x64"
  name   = "denny-es-test1"
  region = "sfo2"
  size   = "512mb"
  user_data = "#cloud-config\nruncmd:\n  - wget -O /tmp/userdata.sh https://raw.githubusercontent.com/DennyZhang/dennytest/master/hashicorp_terraform/userdata.sh\n  - bash /tmp/userdata.sh"
  ssh_keys = [1968722,979830,812123]
}
