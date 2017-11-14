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
```go
func main() {
    db, err := sql.Open("pgx","postgres://user:pwd@localhost:5432/db?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
```

#### 获取结果
Go将数据库操作分为两类：`Query`与`Exec`。两者的区别在于前者会返回结果，而后者不会。


#### 参考资料
[Go database/sql 教程 ](https://yq.aliyun.com/articles/178898?utm_content=m_29337)

[Go database/sql tutorial](http://go-database-sql.org/index.html)

[go-database-sql-tutorial](https://github.com/VividCortex/go-database-sql-tutorial)
