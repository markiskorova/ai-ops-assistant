
output "ec2_public_ip" {
  value = aws_instance.aiops_ec2.public_ip
}

output "private_key_pem_path" {
  value = local_file.private_key_pem.filename
}

output "db_endpoint" {
  value = aws_db_instance.aiops_db.endpoint
}