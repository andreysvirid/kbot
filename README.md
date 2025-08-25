Push to develop branch
          │
          ▼
   GitHub Actions Workflow
          │
   ┌──────┴──────┐
   │             │
   ▼             ▼
 Checkout       Run Tests
 Repository     (make test)
   │
   ▼
Build Docker Image
  docker buildx
  Tag: ghcr.io/<user>/<repo>:v1.0.0-<short_sha>-linux-amd64
   │
   ▼
 Push Image to GHCR
  docker login + push
   │
   ▼
Update Helm Chart
  yq updates values.yaml
  image.tag -> new version
   │
   ▼
Commit Helm Chart
  git add/commit/push
   │
   ▼
ArgoCD monitors repo
  Auto-deploy new Helm chart
   │
   ▼
Kubernetes Cluster
  New Telegram bot version deployed
