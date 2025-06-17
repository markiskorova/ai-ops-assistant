resource "aws_db_instance" "aiops_db" {
  allocated_storage    = 20
  engine               = "postgres"
  engine_version       = "13.4"
  instance_class       = "db.t3.micro"
  username             = var.db_username
  password             = var.db_password
  skip_final_snapshot  = true
  publicly_accessible  = true
  storage_encrypted    = false
  backup_retention_period = 0

  tags = {
    Name = "aiops-db"
  }
}
