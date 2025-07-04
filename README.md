# go_regions

> go语言实现的省市区街道库(4级)，数据更新时间 2025-01-14

# 💡 安装

`go get -u -v github.com/itmisx/go_regions`


# 💡使用

### RegionList(pid) 获取区域列表
  ```go
  go_regions.RegionList(0) // id为0时，获取的为省份列表
  ```
  结果示例
  - id，地区编码
  - name，地区名称
  - level，地区级别 0-省 1-市 2-区 3-街道
  ```json
  [
      {
          "id": 110000, 
          "name": "北京市", 
          "level": 0 
      }, 
      {
          "id": 120000, 
          "name": "天津市", 
          "level": 0
      }, 
      {
          "id": 130000, 
          "name": "河北省", 
          "level": 0
      }
  ]
  ```
### RegionInfo获取区域信息
  ```go
  go_regions.RegionInfo(120000)
  ```
  结果示例

  ```json
  {
      "id": 120000, 
      "name": "天津市", 
      "level": 0
  }
  ```
### RegionName获取区域名称
  ```go
  go_regions.RegionName(120000)
  ```
  结果示例
  ```
  天津市
  ```
# 📝 文档
https://pkg.go.dev/github.com/itmisx/go_regions@v0.0.2#pkg-functions
