variable "do_token" {}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = "${var.do_token}"
}

# Create a new Web Droplet in the nyc2 region
resource "digitalocean_volume" "volume1" {
  region      = "sfo2"
  name        = "volume1"
  size        = 20
  description = "an example volume"
}

resource "digitalocean_droplet" "elasticsearch" {
  image  = "ubuntu-14-04-x64"
  name   = "denny-es-test1"
  region = "sfo2"
  size   = "512mb"
  volume_ids = ["${digitalocean_volume.volume1.id}"]
  ssh_keys = [1968722,979830,812123]
}
