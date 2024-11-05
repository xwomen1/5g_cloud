Bài Tập Triển Khai Kubernetes
Bài 1:
Đề bài: triển khai deployment chạy nginx (default) lên kubernetes và cho phép truy cập từ bên
ngoài thông qua nodePort
Output:
• 1 deployment nginx (replicas=2 pod)

• 1 nodePort service trỏ tới deployment
• thực hiên curl tới nodePort và cho ra kết quả trang web mặc định của nginx
Nộp bài
Sinh viên thực hiện tạo các tài nguyên và lưu lại các kết quả thực hành vào thư mục trên git
và nộp bài bằng link git
1.1. Tạo 2 file deployment của nginx :
 
Cd vào thư mục chứa các file này và chạy lệnh
kubectl apply -f nginx-deployment.yaml
kubectl apply -f nginx-service.yaml
*** Em dùng Minikube và nó chạy K8s dưới dạng container  1 node thì port của container k8s sẽ ánh xạ ra localhost, NodePort ánh xạ vào service chính là port cần map ra ngoài. Nên em sẽ dùng ngrok để public web của em.
1.2. Cài ngrok bằng cmd:
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

 
 
 

Bài 2: 
Đề bài: Triển khai deployment một ứng dụng web tĩnh lên kubernetes cho phép truy cập từ
bên ngoài thông qua nodePort
Output:
• Đóng gói thành công container chứa web tĩnh
o download 1 template tại https://www.free-css.com/free-css-templates)
o sử dụng base image nginx
o lưu ý cấu hình nginx trỏ tới web tĩnh (tham khảo file cấu hình mẫu đơn giản
tại https://gist.github.com/mockra/9062657)
• 1 deployment chạy ứng dụng web tĩnh (replicas=2)
• 1 nodePort service trỏ tới deployment (service web 1)
• Thực hiên curl tới nodePort và cho ra kết quả trang web tĩnh theo template
Nộp bài
Sinh viên thực hiện tạo các tài nguyên và lưu lại các kết quả thực hành vào thư mục trên git
và nộp bài bằng link git

2.1. Tải Template Web Tĩnh
•	Chọn và tải template:
o	Truy cập Free CSS Templates.
o	Chọn một template mà bạn thích và tải về.
o	Giải nén template và đặt các file vào một thư mục, ví dụ my-static-web.
2.2. Tạo Dockerfile
•	Tạo Dockerfile trong thư mục my-static-web:
o	Giải thích:
	FROM nginx:latest: Sử dụng image Nginx mới nhất làm base.
	COPY . /usr/share/nginx/html: Copy tất cả các file trong thư mục hiện tại vào thư mục web của Nginx.
2.3. Xây dựng Image
•	Sử dụng Minikube's Docker Daemon
o	Thiết lập môi trường để sử dụng Docker daemon của Minikube:
& minikube docker-env | Invoke-Expression
o	Xây dựng image:
cd my-static-web
docker build -t my-static-web:latest .
o	Kiểm tra image đã được tạo:
docker images
o	Lưu ý: Bằng cách sử dụng Docker daemon của Minikube, image sẽ có sẵn trong cluster và không cần đẩy lên registry.
________________________________________
2.4 Triển khai Deployment và Service
•	Tạo file web1-deployment.yaml
•	Tạo file web1-service.yaml

•	Triển khai:
kubectl apply -f web1-deployment.yaml
kubectl apply -f web1-service.yaml
________________________________________
2.5. Kiểm tra Ứng dụng Web Tĩnh
minikube service web1-service
lấy được trang web từ localhost 
 
Sau đó ánh xạ ra ngoài bằng ngrok
 

 
________________________________________





Bài 3:
Đề bài: triển khai nginx proxy cho nhiều ứng dụng
- từ level 2, triển khai thêm 1 trang web static thứ hai, khác với static web đã triển khai
- service cho trang web tĩnh mới được lấy tên là web2
- triển khai thêm 1 deployment nginx-proxy đóng vai trò proxy cho cả 2 ứng dụng trên và tạo
nodePort service có tên "nginx-proxy“
- thiết lập cấu hình config của nginx-proxy sao cho:
+ khi gọi tới nginx-proxy với path /web1 > nginx-proxy filter path và forward tới service web 1 >
service web1
+ khi gọi tới nginx-proxy với path /web2 > nginx-proxy filter path và forward tới service web 2 >
service web2

Output:
+ curl http://\<node-ip>:\<node-port>/web1 > trả về static web 1
+ curl http://\<node-ip>:\<node-port>/web2 > trả về static web 2
Nộp bài
Sinh viên thực hiện tạo các tài nguyên và lưu lại các kết quả thực hành vào thư mục trên git
và nộp bài bằng link git



________________________________________
3.1 Triển khai Trang Web Tĩnh Thứ Hai
3.2. Tải Template và Tạo Image
•	Tải template mới:
o	Truy cập Free CSS Templates và chọn một template khác.
o	Giải nén và đặt các file vào thư mục my-static-web2.
•	Tạo Dockerfile cho web2:
o	Trong thư mục my-static-web2, tạo Dockerfile với nội dung:
•	Xây dựng image cho web2:
o	Xây dựng image:
cd my-static-web2
docker build -t my-static-web2:latest .
1.2. Triển khai Deployment và Service cho Web2
•	Tạo file web2-deployment.yaml:
•	Triển khai Deployment:
kubectl apply -f web2-deployment.yaml
•	Tạo file web2-service.yaml:
•	Triển khai Service:
kubectl apply -f web2-service.yaml
________________________________________
3.3. Triển khai Nginx Proxy
3.4. Tạo Cấu hình Nginx Proxy
•	Tạo thư mục nginx-proxy và tạo file nginx.conf:
nginx
•	Giải thích:
o	proxy_pass chuyển tiếp yêu cầu tới các Service nội bộ web1-service và web2-service.
o	Đảm bảo rằng bạn thêm dấu / ở cuối đường dẫn để Nginx xử lý chính xác.
3.5. Tạo Dockerfile cho Nginx Proxy
•	Trong thư mục nginx-proxy, tạo Dockerfile:
•	Xây dựng image cho Nginx Proxy:
docker build -t nginx-proxy:latest .
3.6. Triển khai Deployment và Service cho Nginx Proxy
•	Tạo file nginx-proxy-deployment.yaml:
•	Triển khai Deployment:
kubectl apply -f nginx-proxy-deployment.yaml
•	Tạo file nginx-proxy-service.yaml:
•	Triển khai Service:
kubectl apply -f nginx-proxy-service.yaml

 




________________________________________

 

 

 
 
 

Tài Nguyên Tham Khảo:
•	Minikube Documentation
•	Kubernetes Documentation
•	Docker Desktop for Windows
•	Nginx Official Documentation
•	Kubernetes Services
________________________________________
Một số câu lệnh hỗ trợ fix bug:
Truy cập vào pod tạm thời
kubectl run tmp-shell --rm -it --image=alpine – sh
truy cập vào pod:
kubectl exec -it <nginx-proxy-pod-name> -- /bin/sh
kiểm tra config map
kubectl describe configmap nginx-config
nginx -T

