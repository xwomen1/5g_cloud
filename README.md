# 🚀 KUBERNETES DEPLOYMENT EXERCISES (MINIKUBE)

This document describes **end-to-end Kubernetes hands-on exercises** using **Minikube**, including deploying **Nginx**, **static web applications**, and an **Nginx reverse proxy for multiple services**. The content is structured clearly for learning, execution, and Git submission.

---

## 📌 ENVIRONMENT REQUIREMENTS

* Windows
* Docker Desktop
* Minikube (Kubernetes running as a container – single node)
* kubectl
* Ngrok (to expose NodePort services to the Internet)

---

# 🧪 LAB 1: DEPLOY DEFAULT NGINX

## 🎯 Objective

* Deploy an **Nginx Deployment** (2 replicas)
* Expose it externally using a **NodePort Service**
* Access the default Nginx welcome page

## 📤 Expected Output

* 1 Nginx Deployment (replicas = 2)
* 1 NodePort Service pointing to the deployment
* `curl` to the NodePort returns the default Nginx web page

---

## 1.1. Create Deployment and Service

### 📄 Resource Files

* `nginx-deployment.yaml`
* `nginx-service.yaml`

### ▶️ Apply manifests

```bash
kubectl apply -f nginx-deployment.yaml
kubectl apply -f nginx-service.yaml
```

### ℹ️ Notes (Minikube + NodePort)

* Minikube runs Kubernetes inside a container (single node).
* NodePort is mapped to the Minikube node.
* The Minikube node port is forwarded to `localhost`.
* Ngrok is used to publicly expose the NodePort service.

---

## 1.2. Install Ngrok (Windows – CMD)

```cmd
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
```

---

# 🧪 LAB 2: DEPLOY STATIC WEB APPLICATION

## 🎯 Objective

* Build a Docker image for a static website
* Use **Nginx** as the base image
* Deploy the app to Kubernetes
* Expose it externally using NodePort

## 📤 Expected Output

* A container image successfully built with static web content
* 1 Deployment running the static website (replicas = 2)
* 1 NodePort Service (`web1-service`)
* `curl` to NodePort returns the static website

---

## 2.1. Download Static Web Template

* Visit: [https://www.free-css.com/free-css-templates](https://www.free-css.com/free-css-templates)
* Choose any template
* Extract files into a directory, e.g. `my-static-web`

---

## 2.2. Create Dockerfile

Create a `Dockerfile` inside `my-static-web`:

```dockerfile
FROM nginx:latest
COPY . /usr/share/nginx/html
```

### Explanation

* `FROM nginx:latest`: Use the latest Nginx image
* `COPY`: Copy static files into Nginx web root

---

## 2.3. Build Docker Image (Minikube Docker Daemon)

```powershell
minikube docker-env | Invoke-Expression
cd my-static-web
docker build -t my-static-web:latest .
docker images
```

ℹ️ Using Minikube Docker daemon allows the image to be used directly without pushing to a registry.

---

## 2.4. Deploy Application

### 📄 Create manifests

* `web1-deployment.yaml`
* `web1-service.yaml`

### ▶️ Apply manifests

```bash
kubectl apply -f web1-deployment.yaml
kubectl apply -f web1-service.yaml
```

---

## 2.5. Verify Static Website

```bash
minikube service web1-service
```

* The static website should open in the browser via `localhost`.
* Use **Ngrok** to expose the service publicly.

---

# 🧪 LAB 3: NGINX REVERSE PROXY FOR MULTIPLE APPLICATIONS

## 🎯 Objective

* Deploy a second static website (`web2`)
* Deploy an **Nginx Proxy** in front of `web1` and `web2`
* Route traffic based on URL path

## 📤 Expected Output

```text
curl http://<node-ip>:<node-port>/web1  → returns static web 1
curl http://<node-ip>:<node-port>/web2  → returns static web 2
```

---

## 3.1. Deploy Second Static Website (Web2)

### Download Template

* Choose a different template from Free CSS
* Extract to `my-static-web2`

### Create Dockerfile

```dockerfile
FROM nginx:latest
COPY . /usr/share/nginx/html
```

### Build image

```bash
cd my-static-web2
docker build -t my-static-web2:latest .
```

---

## 3.2. Deploy Web2 to Kubernetes

### 📄 Manifests

* `web2-deployment.yaml`
* `web2-service.yaml`

### ▶️ Apply

```bash
kubectl apply -f web2-deployment.yaml
kubectl apply -f web2-service.yaml
```

---

## 3.3. Deploy Nginx Proxy

### Create Nginx configuration (`nginx.conf`)

```nginx
server {
    listen 80;

    location /web1/ {
        proxy_pass http://web1-service/;
    }

    location /web2/ {
        proxy_pass http://web2-service/;
    }
}
```

ℹ️ Ensure the trailing `/` is included for correct path handling.

---

## 3.4. Create Dockerfile for Nginx Proxy

```dockerfile
FROM nginx:latest
COPY nginx.conf /etc/nginx/conf.d/default.conf
```

### Build image

```bash
docker build -t nginx-proxy:latest .
```

---

## 3.5. Deploy Nginx Proxy to Kubernetes

### 📄 Manifests

* `nginx-proxy-deployment.yaml`
* `nginx-proxy-service.yaml`

### ▶️ Apply

```bash
kubectl apply -f nginx-proxy-deployment.yaml
kubectl apply -f nginx-proxy-service.yaml
```

---

# 🛠️ TROUBLESHOOTING COMMANDS

```bash
# Temporary debug pod
kubectl run tmp-shell --rm -it --image=alpine -- sh

# Exec into a running pod
kubectl exec -it <nginx-proxy-pod-name> -- /bin/sh

# Inspect ConfigMap
kubectl describe configmap nginx-config

# Check Nginx configuration
nginx -T
```

---

# 📚 REFERENCES

* Minikube Documentation
* Kubernetes Official Documentation
* Docker Desktop for Windows
* Nginx Official Documentation
* Kubernetes Services

---

✅ **All manifests, Dockerfiles, and execution results must be committed to Git and submitted as a repository link.**
