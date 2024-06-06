## 功能  
解析nginx配置文件，输出监听的端口，server name，上下文，跳转的地址，以及该规则所在的文件  
## 使用方法
```
ning-config-prase /etc/nginx/nginx.conf
```
其中 `/etc/nginx/nginx.conf` 为nginx配置文件路径,配置文件中的`include`其他文件，会自动递归分析

