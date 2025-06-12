resource "aws_iam_role" "aiops_ec2_role" {
  name = "aiops-ec2-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = {
        Service = "ec2.amazonaws.com"
      },
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_instance_profile" "aiops_instance_profile" {
  name = "aiops-instance-profile"
  role = aws_iam_role.aiops_ec2_role.name
}
