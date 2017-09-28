variable "do_token" {}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = "${var.do_token}"
}

# Create a new Web Droplet in the nyc2 region
resource "digitalocean_droplet" "web" {
  image  = "ubuntu-14-04-x64"
  name   = "denny-vm1-1"
  region = "sfo2"
  size   = "512mb"
}
