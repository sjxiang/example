

# 生成私钥

openssl genrsa -out server.key 2048


# 生成自签名证书
openssl req -x509 \
    -days 365 \
    -config server.cnf \
    -extensions 'req_ext' \
    -key server.key \
    -out server.crt 
