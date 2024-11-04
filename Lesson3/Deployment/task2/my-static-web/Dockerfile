# Sử dụng base image nginx
FROM nginx:latest

# Sao chép toàn bộ nội dung template vào thư mục mặc định của Nginx
COPY . /usr/share/nginx/html

# Thay thế file cấu hình Nginx để trỏ đến thư mục chứa template web
COPY nginx.conf /etc/nginx/conf.d/default.conf
