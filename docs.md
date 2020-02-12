[toc]

# 接口文档

## Demo

### 获取指定id的demo GET /demo/get?id={}

入参：

| 参数 | 类型   | 解释 | 必选 |
| ---- | ------ | ---- | ---- |
| id   | string | ID   | 是   |

示例：

```json
无
```

出参：

```json
{
    "id": "1",
    "name": "demo1"
}
```

解释：

| 参数 | 类型   | 解释 |
| ---- | ------ | ---- |
| id   | string | id   |
| name | string | name |



### 创建一个demo POST /demo/set

入参：

| 参数 | 类型   | 解释 | 必选 |
| ---- | ------ | ---- | ---- |
| id   | string | ID   | 是   |
| name | string | name | 是   |

示例：

```json
{
    "id": "1",
    "name": "demo1"
}
```

出参：

```json
无
```

解释：

无