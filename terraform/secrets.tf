
resource "aws_secretsmanager_secret" "db_credentials" {
  name = "aiops/db_credentials"
}

resource "aws_secretsmanager_secret_version" "db_credentials_version" {
  secret_id     = aws_secretsmanager_secret.db_credentials.id
  secret_string = jsonencode({
    username = var.db_username,
    password = var.db_password
  })
}

resource "aws_secretsmanager_secret" "jwt_secret" {
  name = "aiops/jwt_secret"
}

resource "aws_secretsmanager_secret_version" "jwt_secret_version" {
  secret_id     = aws_secretsmanager_secret.jwt_secret.id
  secret_string = jsonencode({ secret = "your-jwt-secret" })
}
