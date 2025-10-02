# simple-map
simple-map sever with tiles

## 使用说明

### 1. 构建项目

#### Linux 构建
```sh
bash build-linux.sh
```

#### Windows 构建
```sh
bash build-win.sh
```

### 2. 配置文件

编辑 [config.yaml](config.yaml)，配置端口和瓦片源。例如：

```yaml
port: 8080
sources:
  - name: "blue"
    type: "dir"
    path: "F:\\maps\\ttt"
  - name: "world_mbtiles"
    type: "mbtiles"
    path: "./my_map.mbtiles"
  - name: "black"
    type: "dir"
    path: "./black"
```

### 3. 启动服务

```sh
./SimpleMap -c config.yaml
```

### 4. 访问瓦片

- 目录或主题瓦片访问格式：
  ```
  http://localhost:8080/{theme}/{z}/{x}/{y}.png
  ```
  例如：
  ```
  http://localhost:8080/blue/2/1/3.png
  ```

- MBTiles 源访问格式（使用配置中的 name）：
  ```
  http://localhost:8080/world_mbtiles/2/1/3.png
  ```

### 5. 查看版本信息

```sh
./SimpleMap -v
```

### 6. 关闭服务

按 `Ctrl+C` 即可优雅关闭服务。

---

如需更多帮助，请参考源码或 [config.go](config.go)、[tiles.go](tiles.go)、[http.go](http.go) 文件。