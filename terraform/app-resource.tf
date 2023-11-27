################################################################################
# SQS
################################################################################
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

################################################################################
# S3
################################################################################
resource "aws_s3_bucket" "crawler_screenshot_s3_bucket" {
  bucket        = "crawler-url-${var.cluster_name}-${var.environment}-bucket"
  force_destroy = true
}

# resource "aws_s3_bucket_ownership_controls" "crawler_screenshot_s3_bucket_ownership" {
#   bucket = aws_s3_bucket.crawler_screenshot_s3_bucket.id
#   rule {
#     object_ownership = "BucketOwnerPreferred"
#   }
# }

# resource "aws_s3_bucket_public_access_block" "crawler_screenshot_s3_bucket_public_access" {
#   bucket = aws_s3_bucket.crawler_screenshot_s3_bucket.id

#   block_public_acls       = false
#   block_public_policy     = false
#   ignore_public_acls      = false
#   restrict_public_buckets = false
# }
