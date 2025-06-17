
output "ec2_public_ip" {
  value = aws_instance.aiops_ec2.public_ip
}

output "private_key_pem_path" {
  value = local_file.private_key_pem.filename
}

output "db_endpoint" {
  value = aws_db_instance.aiops_db.endpoint
}

output "vpc_id" {
  value = aws_vpc.main.id
}

output "subnet_id" {
  value = aws_subnet.public.id
}

output "ecs_cluster_name" {
  value = aws_ecs_cluster.aiops_cluster.name
}
