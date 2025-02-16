### 开发指南

### 目录

1. [生词本API](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi)

    * 1.1. [获取所有生词本](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-getcategory)
    * 1.2. [添加一个新的生词本](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-addcategory)
    * 1.3. [重命名一个生词本](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-renamecategory)
    * 1.4. [删除一个生词本](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-deletecategory)
    * 1.5. [获取生词本中的单词](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-getwords)
    * 1.6. [添加单词到生词本](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-addwords)
    * 1.7. [删除生词本中的单词](https://my.eudic.net/OpenAPI/doc_api_study#-studylistapi-deletewords)

---

### 1.  生词本API

#### 1.1.  获取所有生词本

1. 接口说明：

    * 请求方式：`GET`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/category`
    * 本接口用于获取生词本信息
    * 生词本的ID是唯一标识，用于调用其他接口
2. | 协议须知： |       |
    | 规则 | 描述 |
    | ------------ | ------- |
    | 传输方式   | HTTPS |
    | 请求方式   | GET   |
    | 字符编码   | UTF-8 |
    | 响应格式   | JSON  |
3. | 接口调用参数： |        |       |      |      |             |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------- | ------ | ------ | ------------- |
    | language       | string | query | 语言 | True | en/fr/de/es |
4. | 输出结果： |       |            |                |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------- | ------------ | ---------------- |
    | data       | Array | 生词本信息 | 请求成功时存在 |
    | message    | text  | 提示信息   |                |

    \*data 生词本信息：

    | 字段名   | 类型 | 含义       | 备注 |
    | ---------- | ------ | ------------ | ------ |
    | id       | text | 生词本ID   |      |
    | language | text | 语言       |      |
    | name     | text | 生词本名称 |      |
5. 示例：

    * Requset Method: `GET`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/category?language=en`
    * Response Code：`200`
    * Response Body：

      ```json
        {
        "data": [
            {
            "id": "0",
            "language": "en",
            "name": "我的生词本"
            },
            {
            "id": "132303016416635230",
            "language": "en",
            "name": "nanana3"
            },
            {
            "id": "132303016647118400",
            "language": "en",
            "name": "新的名称"
            }
        ],
        "message": ""
        }
      ```

#### 1.2.  添加一个新的生词本

1. 接口说明：

    * 请求方式：`POST`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/category`
    * 本接口用于添加一个新的生词本
2. | 协议须知： |       |
    | 规则 | 描述 |
    | ------------ | ------- |
    | 传输方式   | HTTPS |
    | 请求方式   | POST  |
    | 字符编码   | UTF-8 |
    | 响应格式   | JSON  |
3. | 接口调用参数： |        |      |            |      |                                |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------ | ------------ | ------ | -------------------------------- |
    | value          | object | body | 生词本信息 | True | Content-Type: application/json |

    \*value 生词本信息：

    | 字段名   | 数据类型 | 含义     | 必填 | 备注        |
    | ---------- | ---------- | ---------- | ------ | ------------- |
    | language | text     | 语言     | True | en/fr/de/es |
    | name     | text     | 课本名称 | True |             |
4. | 输出结果： |        |                |                |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | -------- | ---------------- | ---------------- |
    | data       | object | 新的生词本信息 | 请求成功时存在 |
    | message    | text   | 提示信息       |                |

    \*data 生词本信息：

    | 字段名   | 类型 | 含义       | 备注 |
    | ---------- | ------ | ------------ | ------ |
    | id       | text | 生词本ID   |      |
    | language | text | 语言       |      |
    | name     | text | 生词本名称 |      |
5. 示例：

    * Requset Method: `POST`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/category`
    * Request Body：

      ```json
        {
        "language": "en",
        "name": "新增生词本1"
        }
      ```
    * Response Code：`201`
    * Response Body：

      ```json
        {
        "data": {
            "id": "132314173819830130",
            "language": "en",
            "name": "新增生词本1"
        },
        "message": ""
        }
      ```

#### 1.3.  重命名一个生词本

1. 接口说明：

    * 请求方式：`PATCH`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/category`
    * 本接口用于修改生词本名称
2. | 协议须知： |       |
    | 规则 | 描述 |
    | ------------ | ------- |
    | 传输方式   | HTTPS |
    | 请求方式   | PATCH |
    | 字符编码   | UTF-8 |
    | 响应格式   | JSON  |
3. | 接口调用参数： |        |      |            |      |                                |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------ | ------------ | ------ | -------------------------------- |
    | value          | object | body | 生词本信息 | True | Content-Type: application/json |

    \*value 生词本信息：

    | 字段名   | 数据类型 | 含义     | 必填 | 备注        |
    | ---------- | ---------- | ---------- | ------ | ------------- |
    | id       | text     | 生词本ID | True |             |
    | language | text     | 语言     | True | en/fr/de/es |
    | name     | text     | 课本名称 | True |             |
4. | 输出结果： |      |          |  |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------ | ---------- | -- |
    | message    | text | 提示信息 |  |
5. 示例：

    * Requset Method: `PATCH`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/category`
    * Request Body：

      ```json
        {
        "id": "132314173819830125",
        "language": "en",
        "name": "新的名称1"
        }
      ```
    * Response Code：`201`
    * Response Body：

      ```json
        {
        "message": "分类重命名成功"
        }
      ```

#### 1.4.  删除一个生词本

1. 接口说明：

    * 请求方式：`DELETE`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/category`
    * 本接口用于删除生词本
2. | 协议须知： |        |
    | 规则 | 描述 |
    | ------------ | -------- |
    | 传输方式   | HTTPS  |
    | 请求方式   | DELETE |
    | 字符编码   | UTF-8  |
    | 响应格式   | JSON   |
3. | 接口调用参数： |        |      |            |      |                                |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------ | ------------ | ------ | -------------------------------- |
    | value          | object | body | 生词本信息 | True | Content-Type: application/json |

    \*value 生词本信息：

    | 字段名   | 数据类型 | 含义     | 必填 | 备注        |
    | ---------- | ---------- | ---------- | ------ | ------------- |
    | id       | text     | 生词本ID | True |             |
    | language | text     | 语言     | True | en/fr/de/es |
    | name     | text     | 课本名称 | True |             |
4. | 输出结果： |      |          |            |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------ | ---------- | ------------ |
    | message    | text | 提示信息 | 失败时存在 |
5. 示例：

    * Requset Method: `DELETE`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/category`
    * Request Body：

      ```json
        {
        "id": "132314173819830125",
        "language": "en",
        "name": "新的名称1"
        }
      ```
    * Response Code：`204`

#### 1.5.  获取生词本中的单词

1. 接口说明：

    * 请求方式：`GET`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/words/{id}`
    * 本接口用于获取生词本中的单词
2. | 协议须知： |       |
    | 规则 | 描述 |
    | ------------ | ------- |
    | 传输方式   | HTTPS |
    | 请求方式   | GET   |
    | 字符编码   | UTF-8 |
    | 响应格式   | JSON  |
3. | 接口调用参数： |         |       |              |       |             |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | --------- | ------- | -------------- | ------- | ------------- |
    | id             | text    | query | 课本ID       | True  |             |
    | language       | text    | query | 语言         | True  | en/fr/de/es |
    | page           | integer | query | 分页页数     | False | 非负数      |
    | page\_size  | integer | query | 分页单词数量 | False | 默认100     |
4. | 输出结果： |       |          |                |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------- | ---------- | ---------------- |
    | data       | Array | 单词信息 | 请求成功时存在 |
    | message    | text  | 提示信息 |                |

    \*data 单词信息：

    | 字段名 | 类型 | 含义 | 备注 |
    | -------- | ------ | ------ | ------ |
    | word   | text | 单词 |      |
    | exp    | text | 释义 |      |
5. 示例：

    * Requset Method: `GET`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/words/0?language=en&category_id=0&page=1&page_size=2`
    * Response Code：`200`
    * Response Body：

      ```json
        {
        "data": [
            {
            "word": "action",
            "exp": "n. 行动；活动；功能；情节；战斗"
            },
            {
            "word": "and",
            "exp": "conj. 和，与；而且；然后；就；但是"
            }
        ],
        "message": ""
        }
      ```

#### 1.6.  添加单词到生词本

1. 接口说明：

    * 请求方式：`POST`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/words`
    * 本接口用于往生词本里添加单词
    * 重复单词不会添加
2. | 协议须知： |       |
    | 规则 | 描述 |
    | ------------ | ------- |
    | 传输方式   | HTTPS |
    | 请求方式   | POST  |
    | 字符编码   | UTF-8 |
    | 响应格式   | JSON  |
3. | 接口调用参数： |        |      |            |      |                                |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------ | ------------ | ------ | -------------------------------- |
    | value          | object | body | 生词本信息 | True | Content-Type: application/json |

    \*value 生词本信息：

    | 字段名   | 数据类型    | 含义     | 必填 | 备注        |
    | ---------- | ------------- | ---------- | ------ | ------------- |
    | id       | text        | 生词本ID | True |             |
    | language | text        | 语言     | True | en/fr/de/es |
    | words    | Array[text] | 单词数组 | True |             |
4. | 输出结果： |      |          |  |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------ | ---------- | -- |
    | message    | text | 提示信息 |  |
5. 示例：

    * Requset Method: `POST`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/words`
    * Request Body：

      ```json
        {
        "id": "0",
        "language": "en",
        "words": [
            "english",
            "french"
        ]
        }
      ```
    * Response Code：`201`
    * Response Body：

      ```json
        {
        "message": "单词导入成功,导入数量 : 2"
        }
      ```

#### 1.7.  删除生词本中的单词

1. 接口说明：

    * 请求方式：`DELETE`
    * 请求地址：`https://api.frdic.com/api/open/v1/studylist/words`
    * 本接口用于删除生词本中的单词
2. | 协议须知： |        |
    | 规则 | 描述 |
    | ------------ | -------- |
    | 传输方式   | HTTPS  |
    | 请求方式   | DELETE |
    | 字符编码   | UTF-8  |
    | 响应格式   | JSON   |
3. | 接口调用参数： |        |      |            |      |                                |
    | 字段名 | 数据类型 | 参数类型 | 含义 | 必填 | 备注 |
    | ---------------- | -------- | ------ | ------------ | ------ | -------------------------------- |
    | value          | object | body | 生词本信息 | True | Content-Type: application/json |

    \*value 生词本信息：

    | 字段名   | 数据类型    | 含义     | 必填 | 备注        |
    | ---------- | ------------- | ---------- | ------ | ------------- |
    | id       | text        | 生词本ID | True |             |
    | language | text        | 语言     | True | en/fr/de/es |
    | words    | Array[text] | 单词数组 | True |             |
4. | 输出结果： |      |          |            |
    | 字段名 | 类型 | 含义 | 备注 |
    | ------------ | ------ | ---------- | ------------ |
    | message    | text | 提示信息 | 失败时存在 |
5. 示例：

    * Requset Method: `DELETE`
    * Request URL：`https://api.frdic.com/api/open/v1/studylist/words?language=en&category_id=0`
    * Request Body：

      ```json
        {
        "id": "0",
        "language": "en",
        "words": [
            "english",
            "french"
        ]
        }
      ```
    * Response Code：`204`

---