
resource "aws_ecs_cluster" "aiops_cluster" {
  name = "aiops-cluster"
}

resource "aws_ecs_task_definition" "api" {
  family                   = "aiops-api"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = var.execution_role_arn
  container_definitions    = jsonencode([{
    name      = "aiops-api",
    image     = var.api_image,
    essential = true,
    portMappings = [{ containerPort = 8080, hostPort = 8080 }]
  }])
}

resource "aws_ecs_service" "api_service" {
  name            = "aiops-api-service"
  cluster         = aws_ecs_cluster.aiops_cluster.id
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets         = [aws_subnet.public.id]
    assign_public_ip = true
    security_groups = [aws_security_group.ec2_sg.id]
  }
}
