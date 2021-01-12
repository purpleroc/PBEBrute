## PBEBrute
遇到个基于[JEECG](http://www.jeecg.com/) 搭建的站点，单纯爆破登录接口太慢、请求包太多。在有SQLi情况下，拿到了salt、username、password。
决定写个东西来本地爆破。

### Usage
```
$ ./PBEBrute


 ___  _         ___             _       
| . \| |_  ___ | . > _ _  _ _ _| |_ ___ 
|  _/| . \/ ._>| . \| '_>| | | | | / ._>
|_|  |___/\___.|___/|_|  |___| |_| \___.


[Error] Open  user.txt  Error!
Usage of ./PBEBrute:
  -f string
    	密码文件，按行分割 (default "dict.txt")
  -i int
    	PBE密钥生成循环次数, 默认1000 (default 1000)
  -t int
    	爆破协程数，默认10000 (default 10000)
  -u string
    	含有: username,salt,password 的待爆破本 格式如：admin,cb362cfeefbf3d8d,RCGTeGiH 按行分割 (default "user.txt")

```

### EXAMPLE
```
$ cat admin.txt
test22,5FMD48RM,ac52e15671a377cf
jeecg,vDDkDzrK,3dd8371f3cf8240e
admin,RCGTeGiH,cb362cfeefbf3d8d
zhagnxiao,go3jJ4zX,f898134e5e52ae11a2ffb2c3b57a4e90
$ 
$ ./PBEBrute -t 100000 -f rockyou.txt -u admin.txt


 ___  _         ___             _       
| . \| |_  ___ | . > _ _  _ _ _| |_ ___ 
|  _/| . \/ ._>| . \| '_>| | | | | / ._>
|_|  |___/\___.|___/|_|  |___| |_| \___.
                                        


2006-01-02T15:04:05Z07:00  [INFO] Start Brute
2006-01-02T15:04:05Z07:00  [INFO] Brute:  test22
======Congratulations======
user:  test22
pass:  123123
2006-01-02T15:04:05Z07:00  [INFO] Brute:  jeecg
======Congratulations======
user:  jeecg
pass:  123456
2006-01-02T15:04:05Z07:00  [INFO] Brute:  admin
======Congratulations======
user:  admin
pass:  123456
2006-01-02T15:04:05Z07:00  [INFO] Brute:  zhagnxiao
======Congratulations======
user:  zhagnxiao
pass:  123
2006-01-02T15:04:05Z07:00  [INFO] END Brute
```

### RESULT
腾讯云 64H 机器：
> 1030650 线程下（max user processes 1030650），大约50s，跑完单用户 14344392 字典。