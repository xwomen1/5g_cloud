Bài Tập Triển Khai Kubernetes và Docker
Repository này chứa các file cấu hình cần thiết để triển khai Nginx và ứng dụng web tĩnh trên Kubernetes sử dụng Minikube. Làm theo từng bài tập để thiết lập deployments, services và proxy Nginx để chuyển hướng traffic đến nhiều ứng dụng web.

Nội Dung
Bài 1: Triển Khai Nginx trên Kubernetes
Bài 2: Triển Khai Ứng Dụng Web Tĩnh
Bài 3: Nginx Proxy cho Nhiều Ứng Dụng
Tài Nguyên Tham Khảo
Bài 1: Triển Khai Nginx trên Kubernetes
Yêu Cầu
Triển khai Nginx trên Kubernetes với 2 replicas và tạo một dịch vụ NodePort để cho phép truy cập từ bên ngoài.

Các Bước Thực Hiện
Tạo File Deployment và Service
Tạo file nginx-deployment.yaml với 2 replicas Nginx.
Tạo file nginx-service.yaml để expose Nginx qua NodePort.
Triển Khai lên Kubernetes
bash
 
kubectl apply -f nginx-deployment.yaml
kubectl apply -f nginx-service.yaml
Công Khai qua Ngrok
Cài đặt Ngrok và sử dụng để công khai dịch vụ Nginx:
bash
 
ngrok http <nodePort>
Bài 2: Triển Khai Ứng Dụng Web Tĩnh
Yêu Cầu
Triển khai một ứng dụng web tĩnh trên Kubernetes với HTML template tùy chỉnh và expose thông qua NodePort.

Các Bước Thực Hiện
Tải Template Web
Tải template từ Free CSS Templates và giải nén vào thư mục (ví dụ: my-static-web).
Tạo Dockerfile
Trong thư mục my-static-web, tạo file Dockerfile với nội dung sau:
Dockerfile
 
FROM nginx:latest
COPY . /usr/share/nginx/html
Build Docker Image
bash
 
minikube docker-env | Invoke-Expression
docker build -t my-static-web:latest .
Tạo File Deployment và Service
Tạo file web1-deployment.yaml và web1-service.yaml.
Triển Khai lên Kubernetes
bash
 
kubectl apply -f web1-deployment.yaml
kubectl apply -f web1-service.yaml
Kiểm Tra với Minikube và Ngrok
bash
 
minikube service web1-service
ngrok http <nodePort>
Bài 3: Nginx Proxy cho Nhiều Ứng Dụng
Yêu Cầu
Triển khai thêm một ứng dụng web tĩnh và thiết lập Nginx proxy để chuyển hướng traffic đến cả hai ứng dụng dựa trên đường dẫn.

Các Bước Thực Hiện
Triển Khai Ứng Dụng Web Tĩnh Thứ Hai

Tải template khác từ Free CSS Templates cho web2 và đặt vào thư mục my-static-web2.
Tạo file Dockerfile trong thư mục my-static-web2 và build image:
Dockerfile
 
FROM nginx:latest
COPY . /usr/share/nginx/html
bash
 
cd my-static-web2
docker build -t my-static-web2:latest .
Tạo file web2-deployment.yaml và web2-service.yaml, sau đó triển khai:
bash
 
kubectl apply -f web2-deployment.yaml
kubectl apply -f web2-service.yaml
Cấu Hình và Triển Khai Nginx Proxy

Tạo file cấu hình nginx-proxy/nginx.conf với các quy tắc định tuyến:
nginx
 
server {
    location /web1/ {
        proxy_pass http://web1-service;
    }
    location /web2/ {
        proxy_pass http://web2-service;
    }
}
Tạo Dockerfile cho Nginx Proxy trong thư mục nginx-proxy:
Dockerfile
 
FROM nginx:latest
COPY nginx.conf /etc/nginx/nginx.conf
Build image cho Nginx Proxy và tạo file deployment, service:
bash
 
cd nginx-proxy
docker build -t nginx-proxy:latest .
kubectl apply -f nginx-proxy-deployment.yaml
kubectl apply -f nginx-proxy-service.yaml
Kiểm Tra Nginx Proxy

Sử dụng lệnh curl để kiểm tra proxy:
bash
 
curl http://<node-ip>:<node-port>/web1
curl http://<node-ip>:<node-port>/web2
Tài Nguyên Tham Khảo
Minikube Documentation
Kubernetes Documentation
Docker Desktop for Windows
Nginx Official Documentation
Kubernetes Services

https://husteduvn-my.sharepoint.com/:w:/g/personal/linh_dt213982_sis_hust_edu_vn/EXsgKJulDbBNsq3NKlqaeugBUSHYKH99xCAnsMI6phpBsw?e=bPU7dA