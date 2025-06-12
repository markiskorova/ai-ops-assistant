
variable "aws_region" {
  default = "us-west-1"
}

variable "ec2_ami" {
  description = "Ubuntu 22.04 AMI for us-west-1"
  default     = "ami-0b2c2a754d5b1d202"
}

variable "instance_type" {
  default = "t3.micro"
}

variable "key_pair_name" {
  default = "aiops-key"
}

variable "db_username" {}
variable "db_password" {}

