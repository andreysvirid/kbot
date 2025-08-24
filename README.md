# KBot CI/CD (GitOps вариант)

Автоматизированный CI/CD для Telegram бота с использованием GitHub Actions, Docker, Helm и ArgoCD (GitOps).

---

## 🚀 CI/CD Workflow

```mermaid
flowchart TD
    A[Push в ветку develop] --> B[GitHub Actions]
    B --> C[Build Docker image]
    C --> D[Push image в ghcr.io]
    D --> E[Обновление helm/values.yaml]
    E --> F[Commit и push обновлённого values.yaml]
    F --> G[ArgoCD автоматически синхронизирует приложение]
    G --> H[Развёртывание в Kubernetes]
    H --> I[Telegram бот готов к использованию]

    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#bbf,stroke:#333,stroke-width:2px
    style C fill:#bfb,stroke:#333,stroke-width:2px
    style D fill:#fbf,stroke:#333,stroke-width:2px
    style E fill:#ffb,stroke:#333,stroke-width:2px
    style F fill:#bbf,stroke:#333,stroke-width:2px
    style G fill:#bfb,stroke:#333,stroke-width:2px
    style H fill:#bbf,stroke:#333,stroke-width:2px
    style I fill:#f9f,stroke:#333,stroke-width:2px
```
---

## 🔹 Настройка GitHub Secrets

- **GITHUB_TOKEN** – создаётся автоматически, проверяем **Read and write permissions**.

> ArgoCD в этом варианте **не требует токена и CLI**, он сам видит изменения в репозитории и синхронизирует deployment.

---

## 🔹 Запуск CI/CD

1. Сделать commit в ветку `develop`:

```bash
git add .
git commit -m "Trigger CI/CD workflow"
git push origin develop
```

2. GitHub Actions автоматически:
- Соберёт Docker image  
- Пушит его в `ghcr.io/<твой_логин>/kbot`  
- Обновит Helm values и закоммитит изменения  

3. ArgoCD автоматически подхватит новые значения и развернёт бота в Kubernetes.

---

## 🔹 Контейнерный образ

Пример адреса:  

```
ghcr.io/den-vasyliev/kbot:v1.0.0-106879e-linux-amd64
```
