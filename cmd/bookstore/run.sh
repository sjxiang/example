



echo "gRPC server 启动"
# 程序名，默认是第一个文件名 
nohup go run server/*.go &

echo "等 5 秒 ..."
sleep 5

echo "gRPC client 启动"
go run client/client.go


echo "gRPC server 下线"
killall converter

echo "结束"
