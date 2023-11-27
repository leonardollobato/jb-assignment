# Assignment (Work In Progress)

1. Terraform
2. Applications
3. Docker Images
4. Argo CD Application Deployment
5. Interaction with Apps
    5.1. Load Balancer URL
    5.2. Commands 
6. GitHub Actions CI
7. Assets for Testing



## 4 Argo CD Application Deployment

### Retrieve Admin password

```
kubectl -n argocd get secret argocd-initial-admin-secret \
          -o jsonpath="{.data.password}" | base64 -d; echo
```

## 5.2 Commands

### Add Product

Replace "load-balancer-url"

```
curl <load-balancer-url>/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}]'
```

``````
curl http://k8s-mainalbgroup-e963c1e2c3-1729184733.us-east-1.elb.amazonaws.com/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}]'
``````


## 7 Assets for Testing

- Jumbo Krulfriet Mild Gekruid 750g
    - https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png

- Jumbo Sinaasappelsap 1L
    - https://jumbo.com/dam-images/fit-in/720x720/Products/28092023_1695921475729_1695921478695_8718452649280_1.png 

- Broodgeluk - Volkoren Bollen - 6 Stuks
    - https://jumbo.com/dam-images/fit-in/720x720/Products/18082023_1692358384367_1692358390660_8718452713592_1.png