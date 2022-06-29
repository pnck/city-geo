## 💡 简介

City Geo 仓库整理了中国城市县区的经纬度数据。

## ⚗ 用法

- [data/data.json](https://github.com/pnck/city-geo/blob/master/data/data.json) 是已经整理好的数据，可直接使用。
- 也可以作为 package 导入到其它工程中，如
```go
package any
import citygeo "github.com/pnck/city-geo/data"

var geoData = citygeo.Data // embeded
```


代码中的 `baiduAK` 请勿在生产环境使用，可能会随时删除。

## 上游
https://github.com/88250/city-geo

## 📄 开源协议

City Geo 使用 [木兰宽松许可证, 第2版](http://license.coscl.org.cn/MulanPSL2) 开源协议。

## 🙏 鸣谢

* [城市数据来源](https://github.com/modood/Administrative-divisions-of-China)
* [百度地图 API](http://lbsyun.baidu.com)
