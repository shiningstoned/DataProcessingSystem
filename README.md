# DataProcessingSystem
数据处理系统
# 功能实现
* 注册
* 登录
* 上传文件
* 下载文件
* 数据处理
  * ip去重
  * ip访问次数排序
# 框架
* kitex: rpc框架
* hertz: http框架
* gorm: orm框架
* mysql: 数据库
* hdfs: 分布式存储框架
# 目录结构
```
.
├── api
│   ├── biz
│   │   ├── handler
│   │   │   ├── api
│   │   │       └── api_service.go
│   │   │   
│   ├── global
│   │   └── client.go
│   ├── initialize
│   │   ├── register.go
│   │   └── rpc
│   │       ├── file_srv.go
│   │       ├── init.go
│   │       └── user_srv.go
│   ├── main.go
|
├── file
│   ├── build.sh
│   ├── handler.go
│   ├── initialize
│   │   ├── hdfs.go
│   │   └── registry.go
│   ├── main.go
│   ├── pkg
│       ├── hdfs.go
│       └── mapreduce.go
├── idl
│   ├── api.thrift
│   ├── file.thrift
│   └── user.thrift
├── temp
└── user
    ├── config
    │   └── config.toml
    ├── dal
    │   └── mysql.go
    ├── handler.go
    ├── initialize
    │   ├── config.go
    │   ├── db.go
    │   └── registry.go
    ├── main.go
```
# 代码实现
## 文件上传
```
	fileHeader, err := c.FormFile("file")
	src, err := fileHeader.Open()
	defer src.Close()
	client, err := hdfs.New("192.168.254.128:9000")
	dst, err := client.Create("/user/" + username.(string) + fileHeader.Filename)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	}
```
## 文件切分
```
        file, err := h.client.Open(filename)
	
	defer file.Close()

	var res [][]string
	var chunk []string
	var count = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if count < 10000 {
			chunk = append(chunk, line)
			count++
		} else {
			res = append(res, chunk)
			chunk = chunk[:0]
		}
	}

	if err = scanner.Err(); err != nil {
		hlog.Error("read file failed")
		return nil, err
	}

	return res, nil
```
## 数据去重
```
  datas, err := s.HDFS.PrePare(req.Filename)

	var lock sync.Mutex
  var wg sync.WaitGroup
	var result map[string]int
	for _, data := range datas {
    wg.Add(1)
		go func() {
      defer wg.Done()
			values := s.MapReduce.Map(data)
			for _, value := range values {
				lock.Lock()
				if _, ok := result[value.Key]; !ok {
					result[value.Key] = 1
				} else {
					result[value.Key]++
				}
				lock.Unlock()
			}
		}()
	}
  wg.Wait()
	var res []byte
	for key := range result {
		keyByte := []byte(key)
		res = append(res, keyByte...)
	}
	file, err := s.HDFS.CreateFile(req.Filename, "rrt")

	_, err = file.Write(res)
```
## 数据排序
``` 
  datas, err := s.HDFS.PrePare(req.Filename)

	var intermediate map[string][]string
	var lock sync.Mutex
	var wg sync.WaitGroup

	for _, data := range datas {
		wg.Add(1)
		go func(data []string) {
			defer wg.Done()
			keyValues := s.MapReduce.Map(data)
			lock.Lock()
			for _, keyValue := range keyValues {
				intermediate[keyValue.Key] = append(intermediate[keyValue.Key], keyValue.Value)
			}
			lock.Unlock()

		}(data)
	}
	wg.Wait()

	var results []KeyValue
	for key, value := range intermediate {
		result := s.MapReduce.Reduce(key, value)
		lock.Lock()
		results = append(results, result)
		lock.Unlock()
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	var res []byte
	for _, keyValue := range results {
		s := fmt.Sprintf("%s %s\n", keyValue.Key, keyValue.Value)
		keyByte := []byte(s)
		res = append(res, keyByte...)
	}
	file, err := s.HDFS.CreateFile(req.Filename, "sbn")

	_, err = file.Write(res)

```
