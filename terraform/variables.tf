
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


variable "api_image" {
  description = "Docker image URI for the AI Ops API container"
  type        = string
}

variable "execution_role_arn" {
  description = "IAM role ARN for ECS task execution"
  type        = string
}
