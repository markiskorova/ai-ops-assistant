provider "aws" {
  region = var.aws_region
}

resource "aws_s3_bucket" "artifact_bucket" {
  bucket = var.artifact_bucket_name
  force_destroy = true
}
