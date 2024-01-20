


# 返回书店中所有书架的列表
curl --location 'http://localhost:9090/v1/shelves' 


# 在书店中创建一个新的书架
curl --location 'http://localhost:9090/v1/shelves' \
--header 'Content-Type: application/json' \
--data '{
    "theme": "科普读物"
}'


# 返回书店中指定的书架（e.g. 返回第一个书架）
curl --location 'http://localhost:9090/v1/shelves/1' 


# 删除书架，包括书架上存储的所有图书（e.g.  删除第二个书架）
curl --location --request DELETE 'http://localhost:9090/v1/shelves/2'


# 返回书架上的图书列表（e.g. 列出第一个书架上的图书）
curl --location 'http://localhost:9090/v1/shelves/1/books'


# 创建一本新图书（e.g. 在第一个书架上创建一本新书）
curl --location 'http://localhost:9090/v1/shelves/1/books' \
--header 'Content-Type: application/json' \
--data '{
    "author":"foo",
    "title":"bar"
}'


# 返回特定的图书（e.g. 获取第二个书架上的第一本图书）
curl --location 'http://localhost:9090/v1/shelves/2/books/1'


# 从书架上删除一本图书（e.g. 删除第一个书架上的第一本书）
curl --location --request DELETE 'http://localhost:9090/v1/shelves/1/books/1'

