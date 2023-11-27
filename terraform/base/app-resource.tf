resource "aws_sqs_queue" "crawler_urls_queue" {
  name = "crawler-url-${var.cluster_name}-${var.environment}-queue"

  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.crawler_queue_deadletter.arn
    maxReceiveCount     = 4
  })
}

resource "aws_sqs_queue" "crawler_queue_deadletter" {
  name = "crawler-url-${var.cluster_name}-${var.environment}-dlq"
}

resource "aws_s3_bucket" "crawler_screenshot_s3_bucket" {
  bucket = "crawler-url-${var.cluster_name}-${var.environment}-bucket"
}
