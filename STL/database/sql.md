#### 导入驱动
使用数据库时，除了`database/sql`包本身，还需要引入想使用的特定数据库驱动。
```
import (
    "database/sql"
    _ "github.com/jackx/pgx/stdlib"
)
```

#### 访问数据
加载驱动包后，需要使用`sql.Open()`来创建`sql.DB`


#### 参考资料
[Go database/sql 教程 ](https://yq.aliyun.com/articles/178898?utm_content=m_29337)
