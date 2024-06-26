provider "aws" {
  region = "us-east-1" # Substitua pela sua região
}

variable "TF_LAMBDA_ZIP_PATH" {
  type = string
}

resource "aws_lambda_function" "lambda-registros" {
  function_name = "lambda-registros"
  role         = aws_iam_role.lambda-registros.arn
  handler      = "main"
  runtime      = "provided.al2023"

  filename     = var.TF_LAMBDA_ZIP_PATH # Recupera o zip da lambda disponibilizado pela esteira

  environment {
    variables = {
      EXAMPLE_ENV_VAR = "example"
    }
  }
}

resource "aws_iam_role_policy" "lambda_exec_policy" {
  name = "crud-api-exec-role-policy"
  role = aws_iam_role.lambda-registros.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "dynamodb:*",
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Action": "sns:*",
            "Effect": "Allow",
            "Resource": "*"
        }     
      ]  
}  
EOF
}

resource "aws_iam_role" "lambda-registros" {
  name = "lambda-registros"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda-registros" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda-registros.name
}





