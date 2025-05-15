worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';
    
    access_log /var/log/nginx/access.log main;
    
    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    keepalive_timeout 65;
    types_hash_max_size 2048;
    
    include /etc/nginx/mime.types;
    default_type application/octet-stream;
    
    # 启用gzip压缩
    gzip on;
    gzip_disable "msie6";
    
    # 商户服务负载均衡
    upstream merchant_service {
        least_conn;
        server merchant-service-1:8080 max_fails=3 fail_timeout=30s;
        server merchant-service-2:8080 max_fails=3 fail_timeout=30s;
        server merchant-service-3:8080 max_fails=3 fail_timeout=30s;
    }
    
    # 域名服务负载均衡
    upstream domain_service {
        least_conn;
        server domain-service-1:8080 max_fails=3 fail_timeout=30s;
        server domain-service-2:8080 max_fails=3 fail_timeout=30s;
    }
    
    # API网关
    server {
        listen 80;
        server_name api.example.com;
        
        # 跨域请求配置
        add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Allow-Methods "GET, POST, PUT, DELETE, OPTIONS";
        add_header Access-Control-Allow-Headers "Origin, X-Requested-With, Content-Type, Accept, Authorization";
        
        # 健康检查
        location /health {
            return 200 "OK";
        }
        
        # 商户服务API
        location /api/v1/merchants {
            proxy_pass http://merchant_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
        
        # 域名服务API
        location /api/v1/domains {
            proxy_pass http://domain_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
        
        # 其他服务路由...
    }
}