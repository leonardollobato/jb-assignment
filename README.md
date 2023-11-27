# Assignment

1. Terraform
2. Applications
3. Docker Images
4. Argo CD Application Deployment
5. Interaction with Apps
    5.1. Load Balancer URL
    5.2. Commands 
6. GitHub Actions CI





## 5.2 Commands

### Add Product

Replace <load-balancer-url>

```
curl <load-balancer-url>/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "The Modern Sound of Betty Carter","url": "https://jumbo.com/dam-images/fit-in/360x360/Products/29092023_1695996129860_1695996141161_8718452601240_1.png"}]'
```