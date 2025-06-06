variable "aws_region" {
  description = "AWS region to deploy resources in"
  type        = string
  default     = "us-west-2"
}

variable "artifact_bucket_name" {
  description = "Name of the S3 bucket for storing build artifacts"
  type        = string
}
