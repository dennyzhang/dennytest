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
  user_data = "#cloud-config\nruncmd:\n  - touch /tmp/a.txt"
  ssh_keys = [1968722,979830,812123]
}
