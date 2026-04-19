---
id: adapters
title: Adapters
description: Supported adapters and usage
keywords: [adapters, MySQL adapter, storage adapter]
authors: [hsluoyz]
---

* Adapters
  * allow
    * load policy 
      * -- `LoadPolicy()` --
    * save policy
      * -- `SavePolicy()` --
  * uses
    * by the enforcer
  * implemented -- outside -- Casbin's source code
    * Reason:🧠keep the core library small🧠

## Supported adapters

### Go

| Adapter                                                                                 | Type     | Author                                                   | AutoSave  | Description                                                                                                                                                   |
|-----------------------------------------------------------------------------------------|----------|----------------------------------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                                       | File     | Casbin                                                   | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                               |
| [Filtered File Adapter (built-in)](/docs/policy-subset-loading)                         | File     | [@faceless-saint](https://github.com/faceless-saint)     | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files with policy subset loading support                            |
| [SQL Adapter](https://github.com/Blank-Xu/sql-adapter)                                  | SQL      | [@Blank-Xu](https://github.com/Blank-Xu)                 | ✅         | MySQL, PostgreSQL, SQL Server, SQLite3 are supported in `master` branch and Oracle is supported in `oracle` branch by `database/sql`                          |
| [Xorm Adapter](https://github.com/casbin/xorm-adapter)                                  | ORM      | Casbin                                                   | ✅         | MySQL, PostgreSQL, TiDB, SQLite, SQL Server, Oracle are supported by [Xorm](https://github.com/go-xorm/xorm/)                                                 |
| [GORM Adapter](https://github.com/casbin/gorm-adapter)                                  | ORM      | Casbin                                                   | ✅         | MySQL, PostgreSQL, Sqlite3, SQL Server are supported by [GORM](https://github.com/go-gorm/gorm)                                                               |
| [GORM Adapter Ex](https://github.com/casbin/gorm-adapter-ex)                            | ORM      | Casbin                                                   | ✅         | MySQL, PostgreSQL, Sqlite3, SQL Server are supported by [GORM](https://github.com/go-gorm/gorm)                                                               |
| [Ent Adapter](https://github.com/casbin/ent-adapter)                                    | ORM      | Casbin                                                   | ✅         | MySQL, MariaDB, PostgreSQL, SQLite, Gremlin-based graph databases are supported by [ent ORM](https://entgo.io/)                                               |
| [Beego ORM Adapter](https://github.com/casbin/beego-orm-adapter)                        | ORM      | Casbin                                                   | ✅         | MySQL, PostgreSQL, Sqlite3 are supported by [Beego ORM](https://beego.wiki/docs/mvc/model/overview/)                                                          |
| [SQLX Adapter](https://github.com/memwey/casbin-sqlx-adapter)                           | ORM      | [@memwey](https://github.com/memwey)                     | ✅         | MySQL, PostgreSQL, SQLite, Oracle are supported by [SQLX](https://github.com/jmoiron/sqlx)                                                                    |
| [Sqlx Adapter](https://github.com/Blank-Xu/sqlx-adapter)                                | ORM      | [@Blank-Xu](https://github.com/Blank-Xu)                 | ✅         | MySQL, PostgreSQL, SQL Server, SQLite3 are supported in `master` branch and Oracle is supported in `oracle` branch by [sqlx](https://github.com/jmoiron/sqlx) |
| [GF ORM Adapter](https://github.com/vance-liu/gdb-adapter)                              | ORM      | [@vance-liu](https://github.com/vance-liu)               | ✅         | MySQL, SQLite, PostgreSQL, Oracle, SQL Server are supported by [GoFrame ORM](https://github.com/gogf/gf/tree/master/contrib/drivers)                          |
| [GoFrame ORM Adapter](https://github.com/kotlin2018/adapter)                            | ORM      | [@kotlin2018](https://github.com/kotlin2018)             | ✅         | MySQL, SQLite, PostgreSQL, Oracle, SQL Server are supported by [GoFrame ORM](https://github.com/gogf/gf/tree/master/contrib/drivers)                          |
| [gf-adapter](https://github.com/zcyc/gf-adapter)                                        | ORM      | [@zcyc](https://github.com/zcyc)                         | ✅         | MySQL, SQLite, PostgreSQL, Oracle, SQL Server are supported by [GoFrame ORM](https://github.com/gogf/gf/tree/master/contrib/drivers)                          |
| [Gdb Adapter](https://github.com/jxo-me/gdb-adapter)                                    | ORM      | [@jxo-me](https://github.com/jxo-me)                     | ✅         | MySQL, SQLite, PostgreSQL, Oracle, SQL Server are supported by [GoFrame ORM](https://github.com/gogf/gf/tree/master/contrib/drivers)                          |
| [GoFrame V2 Adapter](https://github.com/hailaz/gf-casbin-adapter)                       | ORM      | [@hailaz](https://github.com/hailaz)                     | ✅         | MySQL, SQLite, PostgreSQL, Oracle, SQL Server are supported by [GoFrame ORM](https://goframe.org/pages/viewpage.action?pageId=1114686)                        |
| [Bun Adapter](https://github.com/JunNishimura/casbin-bun-adapter)                       | ORM      | [@JunNishimura](https://github.com/JunNishimura)         | ✅         | MySQL, SQLite, PostgreSQL, SQL Server are supported by [Bun ORM](https://bun.uptrace.dev/guide/drivers.html)                                                  |
| [Filtered PostgreSQL Adapter](https://github.com/casbin/casbin-pg-adapter)              | SQL      | Casbin                                                   | ✅         | For [PostgreSQL](https://www.postgresql.org/)                                                                                                                 |
| [Filtered pgx Adapter](https://github.com/pckhoi/casbin-pgx-adapter)                    | SQL      | [@pckhoi](https://github.com/pckhoi)                     | ✅         | PostgreSQL is supported by [pgx](https://github.com/jackc/pgx)                                                                                                |
| [Pgx Adapter](https://github.com/gtoxlili/pgx-adapter)                                  | SQL      | [@gtoxlili](https://github.com/gtoxlili)                 | ✅         | PostgreSQL is supported by [pgx](https://github.com/jackc/pgx), supports customizable column count                                                            |
| [casbin-pgx-adapter](https://github.com/noho-digital/casbin-pgx-adapter)                | SQL      | [@noho-digital](https://github.com/noho-digital)         | ✅         | A PostgreSQL adapter for Casbin using the [pgx](https://github.com/jackc/pgx) driver.                                                                         |
| [PostgreSQL Adapter](https://github.com/cychiuae/casbin-pg-adapter)                     | SQL      | [@cychiuae](https://github.com/cychiuae)                 | ✅         | For [PostgreSQL](https://www.postgresql.org/)                                                                                                                 |
| [RQLite Adapter](https://github.com/edomosystems/rqlite-adapter)                        | SQL      | [EDOMO Systems](https://github.com/edomosystems)         | ✅         | For [RQLite](https://github.com/rqlite/rqlite/)                                                                                                               |
| [MongoDB Adapter](https://github.com/casbin/mongodb-adapter)                            | NoSQL    | Casbin                                                   | ✅         | For [MongoDB](https://www.mongodb.com) based on [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)                                               |
| [RethinkDB Adapter](https://github.com/adityapandey9/rethinkdb-adapter)                 | NoSQL    | [@adityapandey9](https://github.com/adityapandey9)       | ✅         | For [RethinkDB](https://rethinkdb.com/)                                                                                                                       |
| [Cassandra Adapter](https://github.com/casbin/cassandra-adapter)                        | NoSQL    | Casbin                                                   | ❌         | For [Apache Cassandra DB](http://cassandra.apache.org)                                                                                                        |
| [DynamoDB Adapter](https://github.com/HOOQTV/dynacasbin)                                | NoSQL    | [HOOQ](https://github.com/HOOQTV)                        | ❌         | For [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                                       |
| [Dynacasbin](https://github.com/NewbMiao/dynacasbin)                                    | NoSQL    | [NewbMiao](https://github.com/NewbMiao)                  | ✅         | For [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                                       |
| [ArangoDB Adapter](https://github.com/adamwasila/arangodb-adapter)                      | NoSQL    | [@adamwasila](https://github.com/adamwasila)             | ✅         | For [ArangoDB](https://www.arangodb.com/)                                                                                                                     |
| [Amazon S3 Adapter](https://github.com/Soluto/casbin-minio-adapter)                     | Cloud    | [Soluto](https://github.com/Soluto)                      | ❌         | For [Minio](https://github.com/minio/minio) and [Amazon S3](https://aws.amazon.com/s3/)                                                                       |
| [Go CDK Adapter](https://github.com/bartventer/casbin-go-cloud-adapter)                 | Cloud    | [@bartventer](https://github.com/bartventer)             | ✅         | Adapter based on [Go Cloud Dev Kit](https://gocloud.dev/) that supports: Amazon DynamoDB, Azure CosmosDB, GCP Firestore, MongoDB, In-Memory                   |
| [Azure Cosmos DB Adapter](https://github.com/spacycoder/cosmos-casbin-adapter)          | Cloud    | [@spacycoder](https://github.com/spacycoder)             | ✅         | For [Microsoft Azure Cosmos DB](https://docs.microsoft.com/en-us/azure/cosmos-db/introduction)                                                                |
| [GCP Firestore Adapter](https://github.com/reedom/casbin-firestore-adapter)             | Cloud    | [@reedom](https://github.com/reedom)                     | ❌         | For [Google Cloud Platform Firestore](https://cloud.google.com/firestore/)                                                                                    |
| [GCP Cloud Storage Adapter](https://github.com/qurami/casbin-cloud-storage-adapter)     | Cloud    | [qurami](https://github.com/qurami)                      | ❌         | For [Google Cloud Platform Cloud Storage](https://cloud.google.com/storage/)                                                                                  |
| [GCP Cloud Spanner Adapter](https://github.com/flowerinthenight/casbin-spanner-adapter) | Cloud    | [@flowerinthenight](https://github.com/flowerinthenight) | ✅         | For [Google Cloud Platform Cloud Spanner](https://cloud.google.com/spanner/)                                                                                  |
| [Consul Adapter](https://github.com/ankitm123/consul-adapter)                           | KV store | [@ankitm123](https://github.com/ankitm123)               | ❌         | For [HashiCorp Consul](https://www.consul.io/)                                                                                                                |
| [Redis Adapter (Redigo)](https://github.com/casbin/redis-adapter)                       | KV store | Casbin                                                   | ✅         | For [Redis](https://redis.io/)                                                                                                                                |
| [Redis Adapter (go-redis)](https://github.com/mlsen/casbin-redis-adapter)               | KV store | [@mlsen](https://github.com/mlsen)                       | ✅         | For [Redis](https://redis.io/)                                                                                                                                |
| [Etcd Adapter](https://github.com/sebastianliu/etcd-adapter)                            | KV store | [@sebastianliu](https://github.com/sebastianliu)         | ❌         | For [etcd](https://github.com/coreos/etcd)                                                                                                                    |
| [BoltDB Adapter](https://github.com/speza/casbin-bolt-adapter)                          | KV store | [@speza](https://github.com/speza)                       | ✅         | For [Bolt](https://github.com/boltdb/bolt)                                                                                                                    |
| [Bolt Adapter](https://github.com/wirepair/bolt-adapter)                                | KV store | [@wirepair](https://github.com/wirepair)                 | ❌         | For [Bolt](https://github.com/boltdb/bolt)                                                                                                                    |
| [BadgerDB Adapter](https://github.com/inits/casbin-badgerdb-adapter)                    | KV store | [@inits](https://github.com/inits)                       | ✅         | For [BadgerDB](https://github.com/dgraph-io/badger)                                                                                                           |
| [Protobuf Adapter](https://github.com/casbin/protobuf-adapter)                          | Stream   | Casbin                                                   | ❌         | For [Google Protocol Buffers](https://developers.google.com/protocol-buffers/)                                                                                |
| [JSON Adapter](https://github.com/casbin/json-adapter)                                  | String   | Casbin                                                   | ❌         | For [JSON](https://www.json.org/)                                                                                                                             |
| [String Adapter](https://github.com/qiangmzsx/string-adapter)                           | String   | [@qiangmzsx](https://github.com/qiangmzsx)               | ❌         | For String                                                                                                                                                    |
| [HTTP File Adapter](https://github.com/h4ckedneko/casbin-httpfs)                        | HTTP     | [@h4ckedneko](https://github.com/h4ckedneko)             | ❌         | For [http.FileSystem](https://golang.org/src/net/http/fs.go)                                                                                                  |
| [FileSystem Adapter](https://github.com/naucon/casbin-fs-adapter)                       | File     | [@naucon](https://github.com/naucon)                     | ❌         | For [fs.FS](https://pkg.go.dev/io/fs) and [embed.FS](https://pkg.go.dev/embed)                                                                                |
| [NATS JetStream Adapter](https://github.com/grepplabs/casbin-jetstream)                 | KV store | [grepplabs](https://github.com/grepplabs)                | ✅         | For [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream)                                                                                            |
| [Kubernetes Adapter](https://github.com/grepplabs/casbin-kube)                          | Cloud    | [grepplabs](https://github.com/grepplabs)                | ✅         | For [Kubernetes](https://kubernetes.io)                                                                                                                       |

### Java

| Adapter                                                             | Type     | Author                                     | AutoSave  | Description                                                                                                                                |
|---------------------------------------------------------------------|----------|--------------------------------------------|-----------|--------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                   | File     | Casbin                                     | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                            |
| [JDBC Adapter](https://github.com/jcasbin/jdbc-adapter)             | JDBC     | Casbin                                     | ✅         | MySQL, Oracle, PostgreSQL, DB2, Sybase, SQL Server are supported by [JDBC](https://docs.oracle.com/cd/E19226-01/820-7688/gawms/index.html) |
| [Hibernate Adapter](https://github.com/jcasbin/hibernate-adapter)   | ORM      | Casbin                                     | ✅         | Oracle, DB2, SQL Server, Sybase, MySQL, PostgreSQL are supported by [Hibernate](http://www.hibernate.org/)                                 |
| [MyBatis Adapter](https://github.com/jcasbin/mybatis-adapter)       | ORM      | Casbin                                     | ✅         | MySQL, Oracle, PostgreSQL, DB2, Sybase, SQL Server (the same as JDBC) are supported by [MyBatis 3](https://mybatis.org/mybatis-3/)         |
| [Hutool Adapter](https://github.com/mapleafgo/jcasbin-extra)        | ORM      | [@mapleafgo](https://github.com/mapleafgo) | ✅         | MySQL, Oracle, PostgreSQL, SQLite are supported by [Hutool](https://github.com/looly/hutool)                                               |
| [MongoDB Adapter](https://github.com/jcasbin/jcasbin-mongo-adapter) | NoSQL    | Casbin                                     | ✅         | MongoDB is supported by [mongodb-driver-sync](https://mongodb.github.io/mongo-java-driver/)                                                |
| [DynamoDB Adapter](https://github.com/jcasbin/dynamodb-adapter)     | NoSQL    | Casbin                                     | ❌         | For [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                    |
| [Redis Adapter](https://github.com/jcasbin/redis-adapter)           | KV store | Casbin                                     | ✅         | For [Redis](https://redis.io/)                                                                                                             |

### Node.js

| Adapter                                                                                                     | Type       | Author                                                               | AutoSave  | Description                                                                                                                                      |
|-------------------------------------------------------------------------------------------------------------|------------|----------------------------------------------------------------------|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                                                           | File       | Casbin                                                               | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                  |
| [Filtered File Adapter (built-in)](/docs/policy-subset-loading)                                             | File       | Casbin                                                               | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files with policy subset loading support               |
| [String Adapter (built-in)](https://github.com/casbin/node-casbin/blob/master/src/persist/stringAdapter.ts) | String     | [@calebfaruki](https://github.com/calebfaruki)                       | ❌         | For String                                                                                                                                       |
| [Basic Adapter](https://github.com/node-casbin/basic-adapter)                                               | Native ORM | Casbin                                                               | ✅         | pg, mysql, mysql2, sqlite3, oracledb, mssql are supported by the adapter itself                                                                  |
| [Sequelize Adapter](https://github.com/node-casbin/sequelize-adapter)                                       | ORM        | Casbin                                                               | ✅         | MySQL, PostgreSQL, SQLite, Microsoft SQL Server are supported by [Sequelize](https://github.com/sequelize/sequelize)                             |
| [TypeORM Adapter](https://github.com/node-casbin/typeorm-adapter)                                           | ORM        | Casbin                                                               | ✅         | MySQL, PostgreSQL, MariaDB, SQLite, MS SQL Server, Oracle, WebSQL, MongoDB are supported by [TypeORM](https://github.com/typeorm/typeorm)        |
| [Prisma Adapter](https://github.com/node-casbin/prisma-adapter)                                             | ORM        | Casbin                                                               | ✅         | MySQL, PostgreSQL, MariaDB, SQLite, MS SQL Server, AWS Aurora, Azure SQL are supported by [Prisma](https://www.prisma.io/)                       |
| [Drizzle Adapter](https://github.com/node-casbin/drizzle-adapter)                                           | ORM        | Casbin                                                               | ✅         | PostgreSQL, MySQL, SQLite, Turso, Neon, PlanetScale, Vercel Postgres, Xata are supported by [Drizzle ORM](https://orm.drizzle.team/)             |
| [Knex Adapter](https://github.com/knex/casbin-knex-adapter)                                                 | ORM        | [knex](https://github.com/knex)                                      | ✅         | MSSQL, MySQL, PostgreSQL, SQLite3, Oracle are supported by [Knex.js](https://knexjs.org/)                                                        |
| [Objection.js Adapter](https://github.com/willsoto/casbin-objection-adapter)                                | ORM        | [@willsoto](https://github.com/willsoto)                             | ✅         | MSSQL, MySQL, PostgreSQL, SQLite3, Oracle are supported by [Objection.js](https://vincit.github.io/objection.js/)                                |
| [MikroORM Adapter](https://github.com/baisheng/casbin-mikroorm-adapter)                                     | ORM        | [@baisheng](https://github.com/baisheng)                             | ✅         | MongoDB, MySQL, MariaDB, PostgreSQL, SQLite are supported by [MikroORM](https://mikro-orm.io/)                                                   |
| [Node PostgreSQL Native Adapter](https://github.com/touchifyapp/casbin-pg-adapter)                          | SQL        | [@touchifyapp](https://github.com/touchifyapp)                       | ✅         | PostgreSQL adapter with advanced policy subset loading support and improved performances built with [node-postgres](https://node-postgres.com/). |
| [Mongoose Adapter](https://github.com/node-casbin/mongoose-adapter)                                         | NoSQL      | [elastic.io](https://github.com/elasticio) and Casbin                | ✅         | MongoDB is supported by [Mongoose](https://mongoosejs.com/)                                                                                      |
| [Mongoose Adapter (No-Transaction)](https://github.com/minhducck/casbin-mongoose-adapter)                   | NoSQL      | [minhducck](https://github.com/minhducck)                            | ✅         | MongoDB is supported by [Mongoose](https://mongoosejs.com/)                                                                                      |
| [Node MongoDB Native Adapter](https://github.com/NathanBhanji/mongodb-casbin-adapter)                       | NoSQL      | [NathanBhanji](https://github.com/NathanBhanji)                      | ✅         | For [Node MongoDB Native](https://mongodb.github.io/node-mongodb-native/)                                                                        |
| [Node MongoDB Native Adapter](https://github.com/juicycleff/casbin-mongodb-adapter)                         | NoSQL      | [@juicycleff](https://github.com/juicycleff)                         | ✅         | For [Node MongoDB Native](https://mongodb.github.io/node-mongodb-native/)                                                                        |
| [DynamoDB Adapter](https://github.com/fospitia/casbin-dynamodb-adapter)                                     | NoSQL      | [@fospitia](https://github.com/fospitia)                             | ✅         | For [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                          |
| [Couchbase Adapter](https://github.com/MarkMYoung/casbin-couchbase-adapter)                                 | NoSQL      | [@MarkMYoung](https://github.com/MarkMYoung)                         | ✅         | For [Couchbase](https://www.couchbase.com/)                                                                                                      |
| [Redis Adapter](https://github.com/node-casbin/redis-adapter)                                               | KV store   | Casbin                                                               | ❌         | For [Redis](https://redis.io/)                                                                                                                   |
| [Redis Adapter](https://github.com/NandaKishorJeripothula/node-casbin-redis-adapter)                        | KV store   | [@NandaKishorJeripothula](https://github.com/NandaKishorJeripothula) | ❌         | For [Redis](https://redis.io/)                                                                                                                   |

### PHP

| Adapter                                                                         | Type     | Author                               | AutoSave  | Description                                                                                                                                                                                                                         |
|---------------------------------------------------------------------------------|----------|--------------------------------------|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                               | File     | Casbin                               | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                                                                                                     |
| [Database Adapter](https://github.com/php-casbin/database-adapter)              | ORM      | Casbin                               | ✅         | MySQL, PostgreSQL, SQLite, Microsoft SQL Server are supported by [techone/database](https://github.com/techoner/database)                                                                                                           |
| [Zend Db Adapter](https://github.com/php-casbin/zend-db-adapter)                | ORM      | Casbin                               | ✅         | MySQL, PostgreSQL, SQLite, Oracle, IBM DB2, Microsoft SQL Server, Other PDO Driver are supported by [zend-db](https://docs.zendframework.com/zend-db/)                                                                              |
| [Doctrine DBAL Adapter (Recommend)](https://github.com/php-casbin/dbal-adapter) | ORM      | Casbin                               | ✅         | Powerful PHP database abstraction layer ([DBAL](https://github.com/doctrine/dbal)) with many features for database schema introspection and management.                                                                             |
| [Medoo Adapter](https://github.com/php-casbin/medoo-adapter)                    | ORM      | Casbin                               | ✅         | [Medoo](https://github.com/catfan/Medoo) is a lightweight PHP Database Framework to Accelerate Development, supports all SQL databases, including `MySQL`, `MSSQL`, `SQLite`, `MariaDB`, `PostgreSQL`, `Sybase`, `Oracle` and more. |
| [Laminas-db Adapter](https://github.com/php-casbin/laminas-db-adapter)          | ORM      | Casbin                               | ✅         | MySQL, PostgreSQL, Oracle, IBM DB2, Microsoft SQL Server, PDO, etc. are supported by [laminas-db](https://github.com/laminas/laminas-db)                                                                                            |
| [Zend-db Adapter](https://github.com/php-casbin/zend-db-adapter)                | ORM      | Casbin                               | ✅         | MySQL, PostgreSQL, Oracle, IBM DB2, Microsoft SQL Server, PDO, etc. are supported by [zend-db](https://github.com/zendframework/zend-db)                                                                                            |
| [ThinkORM Adapter (ThinkPHP)](https://github.com/getandpost/tp3-adapter)        | ORM      | Casbin                               | ✅         | MySQL, PostgreSQL, SQLite, Oracle, Microsoft SQL Server, MongoDB are supported by [ThinkORM](https://github.com/top-think/think-orm)                                                                                                |
| [Redis Adapter](https://github.com/nsnake/CasbinAdapter-Redis-Adapter)          | KV store | [@nsnake](https://github.com/nsnake) | ❌         | For [Redis](https://redis.io/)                                                                                                                                                                                                      |

### Python

| Adapter                                                                                  | Type     | Author                                                       | AutoSave                               | Description                                                                                                                                                                   |
|------------------------------------------------------------------------------------------|----------|--------------------------------------------------------------|----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                                        | File     | Casbin                                                       | ❌                                      | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                                               |
| [Django ORM Adapter](https://github.com/officialpycasbin/django-orm-adapter)             | ORM      | Casbin                                                       | ✅                                      | PostgreSQL, MariaDB, MySQL, Oracle, SQLite, IBM DB2, Microsoft SQL Server, Firebird, ODBC are supported by [Django ORM](https://docs.djangoproject.com/en/3.0/ref/databases/) |
| [SQLObject Adapter](https://github.com/officialpycasbin/sqlobject-adapter)               | ORM      | Casbin                                                       | ✅                                      | PostgreSQL, MySQL, SQLite, Microsoft SQL Server, Firebird, Sybase, MAX DB, pyfirebirdsql are supported by [SQLObject](http://www.sqlobject.org/index.html)                    |
| [SQLAlchemy Adapter](https://github.com/officialpycasbin/sqlalchemy-adapter)             | ORM      | Casbin                                                       | ✅                                      | PostgreSQL, MySQL, SQLite, Oracle, Microsoft SQL Server, Firebird, Sybase are supported by [SQLAlchemy](https://www.sqlalchemy.org/)                                          |
| [Async SQLAlchemy Adapter](https://github.com/officialpycasbin/async-sqlalchemy-adapter) | ORM      | Casbin                                                       | ✅                                      | PostgreSQL, MySQL, SQLite, Oracle, Microsoft SQL Server, Firebird, Sybase are supported by [SQLAlchemy](https://www.sqlalchemy.org/)                                          |
| [Async Databases Adapter](https://github.com/officialpycasbin/casbin-databases-adapter)  | ORM      | Casbin                                                       | ✅                                      | PostgreSQL, MySQL, SQLite, Oracle, Microsoft SQL Server, Firebird, Sybase are supported by [Databases](https://www.encode.io/databases/)                                      |
| [Peewee Adapter](https://github.com/shblhy/peewee-adapter)                               | ORM      | [@shblhy](https://github.com/shblhy)                         | ✅                                      | PostgreSQL, MySQL, SQLite are supported by [Peewee](http://docs.peewee-orm.com/)                                                                                              |
| [MongoEngine Adapter](https://github.com/zhangbailong945/mongoengine_adapter)            | ORM      | [@zhangbailong945](https://github.com/zhangbailong945)       | ❌                                      | MongoDB is supported by [MongoEngine](http://mongoengine.org/)                                                                                                                |
| [Pony ORM Adapter](https://github.com/drorvinkler/pycasbin-pony-adapter)                 | ORM      | [@drorvinkler](https://github.com/drorvinkler)               | ✅                                      | MySQL, PostgreSQL, SQLite, Oracle, CockroachDB are supported by [Pony ORM](https://ponyorm.org/)                                                                              |
| [Tortoise ORM Adapter](https://github.com/thearchitector/casbin-tortoise-adapter)        | ORM      | [@thearchitector](https://github.com/thearchitector)         | ✅                                      | PostgreSQL (>=9.4), MySQL, MariaDB, and SQLite are supported by [Tortoise ORM](https://tortoise.github.io/databases.html)                                                     |
| [Async Ormar Adapter](https://github.com/shepilov-vladislav/ormar-casbin-adapter)        | ORM      | [@shepilov-vladislav](https://github.com/shepilov-vladislav) | ✅                                      | PostgreSQL, MySQL, SQLite are supported by [Ormar](https://github.com/collerek/ormar/)                                                                                        |
| [SQLModel Adapter](https://github.com/shepilov-vladislav/async-casbin-sqlmodel-adapter)  | ORM      | [@shepilov-vladislav](https://github.com/shepilov-vladislav) | ✅                                      | PostgreSQL, MySQL, SQLite are supported by [SQLModel](https://github.com/tiangolo/sqlmodel)                                                                                   |
| [Couchbase Adapter](https://github.com/ScienceLogic/casbin-couchbase-adapter)            | NoSQL    | [ScienceLogic](https://github.com/ScienceLogic)              | ✅ (without `remove_filtered_policy()`) | For [Couchbase](https://www.couchbase.com/)                                                                                                                                   |
| [DynamoDB Adapter](https://github.com/abqadeer/python-dycasbin)                          | NoSQL    | [@abqadeer](https://github.com/abqadeer/)                    | ✅                                      | For [DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                                                              |
| [Pymongo Adapter](https://github.com/officialpycasbin/pymongo-adapter)                   | NoSQL    | Casbin                                                       | ❌                                      | MongoDB is supported by [Pymongo](https://pypi.org/project/pymongo/)                                                                                                          |
| [Redis Adapter](https://github.com/officialpycasbin/redis-adapter)                       | KV store | Casbin                                                       | ✅                                      | For [Redis](https://redis.io/)                                                                                                                                                |
| [GCP Firebase Adapter](https://github.com/devrushi41/pycasbin-firebase-adapter)          | Cloud    | [@devrushi41](https://github.com/devrushi41)                 | ✅                                      | For [Google Cloud Platform Firebase](https://firebase.google.com/)                                                                                                            |

### .NET

| Adapter                                                                           | Type     | Author                                                 | AutoSave  | Description                                                                                                                                                                                                                         |
|-----------------------------------------------------------------------------------|----------|--------------------------------------------------------|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                                 | File     | Casbin                                                 | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                                                                                                     |
| [EF Adapter](https://github.com/casbin-net/EF-Adapter)                            | ORM      | Casbin                                                 | ❌         | MySQL, PostgreSQL, SQLite, Microsoft SQL Server, Oracle, DB2, etc. are supported by [Entity Framework 6](https://docs.microsoft.com/en-us/ef/ef6/)                                                                                  |
| [EFCore Adapter](https://github.com/casbin-net/EFCore-Adapter)                    | ORM      | Casbin                                                 | ✅         | MySQL, PostgreSQL, SQLite, Microsoft SQL Server, Oracle, DB2, etc. are supported by [Entity Framework Core](https://docs.microsoft.com/en-us/ef/core/)                                                                              |
| [Linq2DB Adapter](https://github.com/Tirael/Linq2DB-Adapter)                      | ORM      | [@Tirael](https://github.com/Tirael)                   | ✅         | [MySQL, PostgreSQL, SQLite, Microsoft SQL Server, Oracle, Access, Firebird, Sybase, etc.](https://github.com/linq2db/linq2db/blob/master/Tests/Base/TestProvName.cs) are supported by [linq2db](https://github.com/linq2db/linq2db) |
| [Azure Cosmos DB Adapter](https://github.com/sagarkhandelwal/Casbin-Using-Cosmos) | Cloud    | [@sagarkhandelwal](https://github.com/sagarkhandelwal) | ✅         | For [Microsoft Azure Cosmos DB](https://docs.microsoft.com/en-us/azure/cosmos-db/introduction)                                                                                                                                      |
| [Redis Adapter](https://github.com/casbin-net/redis-adapter)                      | KV store | Casbin                                                 | ✅         | For [Redis](https://redis.io/)                                                                                                                                                                                                      |
| [SqlSugar Adapter](https://github.com/SharpFort/sqlsugar-adapter)                 | ORM      | [@SharpFort](https://github.com/SharpFort)             | ✅         | MySQL, PostgreSQL, SQLite, Microsoft SQL Server, Oracle, etc. are supported by [SqlSugar](https://github.com/donet5/SqlSugar)                                                                                                       |

### Rust

| Adapter                                                                | Type    | Author                                             | AutoSave  | Description                                                                                                                                                    |
|------------------------------------------------------------------------|---------|----------------------------------------------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                      | File    | Casbin                                             | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                                |
| [Diesel Adapter](https://github.com/casbin-rs/diesel-adapter)          | ORM     | Casbin                                             | ✅         | [SQLite, PostgreSQL, MySQL](https://github.com/diesel-rs/diesel/blob/master/guide_drafts/backend_installation.md) are supported by [Diesel](http://diesel.rs/) |
| [Sqlx Adapter](https://github.com/casbin-rs/sqlx-adapter)              | ORM     | Casbin                                             | ✅         | PostgreSQL, MySQL are supported by [Sqlx](https://github.com/launchbadge/sqlx) with fully asynchronous operation                                               |
| [SeaORM Adapter](https://github.com/lingdu1234/sea_orm_casbin_adapter) | ORM     | [@lingdu1234](https://github.com/lingdu1234)       | ✅         | PostgreSQL, MySQL, SQLite are supported by [SeaORM](https://github.com/SeaQL/sea-orm) with fully asynchronous operation                                        |
| [SeaORM Adapter](https://github.com/ZihanType/sea-orm-adapter)         | ORM     | [@ZihanType](https://github.com/ZihanType)         | ✅         | PostgreSQL, MySQL, SQLite are supported by [SeaORM](https://github.com/SeaQL/sea-orm) with fully asynchronous operation                                        |
| [Rbatis Adapter](https://github.com/jiashiwen/casbin-rbatis-adapter)   | ORM     | [rbatis](https://github.com/rbatis)                | ✅         | MySQL, PostgreSQL, SQLite, SQL Server, MariaDB, TiDB, CockroachDB, Oracle are supported by [Rbatis](https://github.com/rbatis/rbatis)                          |
| [DynamodDB Adapter](https://github.com/fospitia/dynamodb-adapter)      | NoSQL   | [@fospitia](https://github.com/fospitia)           | ✅         | For [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)                                                                                                        |
| [MongoDB Adapter](https://github.com/wangjun861205/nb-mongo-adapter)   | MongoDB | [@wangjun861205](https://github.com/wangjun861205) | ✅         | For [MongoDB](https://www.mongodb.com)                                                                                                                         |
| [JSON Adapter](http://github.com/casbin-rs/json-adapter)               | String  | Casbin                                             | ✅         | For [JSON](https://json.org/)                                                                                                                                  |
| [YAML Adapter](https://github.com/casbin-rs/yaml-adapter)              | String  | Casbin                                             | ✅         | For [YAML](https://yaml.org/)                                                                                                                                  |
| [String Adapter](https://github.com/casbin-rs/string-adapter)          | String  | Casbin                                             | ❌         | For [String](https://doc.rust-lang.org/std/string/struct.String.html)                                                                                          |

### Ruby

| Adapter                                                                 | Type  | Author                                      | AutoSave  | Description                                                                                                                                                                                           |
|-------------------------------------------------------------------------|-------|---------------------------------------------|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                       | File  | Casbin                                      | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                                                                       |
| [Sequel Adapter](https://github.com/CasbinRuby/casbin-ruby-sql-adapter) | ORM   | [CasbinRuby](https://github.com/CasbinRuby) | ✅         | [ADO, Amalgalite, IBM_DB, JDBC, MySQL, Mysql2, ODBC, Oracle, PostgreSQL, SQLAnywhere, SQLite3, and TinyTDS](http://sequel.jeremyevans.net/) are supported by [Sequel](http://sequel.jeremyevans.net/) |

### Swift

| Adapter                                                                                                                   | Type   | Author | AutoSave  | Description                                                                                                                                            |
|---------------------------------------------------------------------------------------------------------------------------|--------|--------|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                                                                         | File   | Casbin | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                                        |
| [Memory Adapter (built-in)](https://github.com/casbin/SwiftCasbin/blob/master/Sources/Casbin/Adapter/MemoryAdapter.swift) | Memory | Casbin | ❌         | For memory                                                                                                                                             |
| [Fluent Adapter](https://github.com/SwiftCasbin/fluent-adapter)                                                           | ORM    | Casbin | ✅         | [PostgreSQL, SQLite, MySQL, MongoDB](https://docs.vapor.codes/4.0/fluent/overview/#drivers) are supported by [Fluent](https://github.com/vapor/fluent) |

### Lua

| Adapter                                                                      | Type  | Author                                         | AutoSave  | Description                                                                                                                        |
|------------------------------------------------------------------------------|-------|------------------------------------------------|-----------|------------------------------------------------------------------------------------------------------------------------------------|
| [File Adapter (built-in)](https://github.com/dancer1325/apache-casbin/tree/master/persist/file-adapter)                            | File  | Casbin                                         | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files                                    |
| [Filtered File Adapter (built-in)](/docs/policy-subset-loading)              | File  | Casbin                                         | ❌         | For [.CSV (Comma-Separated Values)](https://en.wikipedia.org/wiki/Comma-separated_values) files with policy subset loading support |
| [LuaSQL Adapter](https://github.com/casbin-lua/luasql-adapter)               | ORM   | Casbin                                         | ✅         | MySQL, PostgreSQL, SQLite3 are supported by [LuaSQL](http://lunarmodules.github.io/luasql/)                                        |
| [4DaysORM Adapter](https://github.com/casbin-lua/4daysorm-adapter)           | ORM   | Casbin                                         | ✅         | MySQL, SQLite3 are supported by [4DaysORM](https://github.com/itdxer/4DaysORM)                                                     |
| [OpenResty Adapter](https://github.com/tom2nonames/lua-resty-casbin-adapter) | ORM   | [@tom2nonames](https://github.com/tom2nonames) | ✅         | MySQL, PostgreSQL are supported by it                                                                                              |


## TODO:

* `e.LoadPolicy()`
  * refresh policies -- from -- storage
* `SavePolicy()`
  * uses
    * MANUALLY persist ALL rules
  * use cases
    * Auto-Save disabled

## how to use your OWN storage adapter?

TODO:
Integrate a custom adapter:

```go
import (
    "github.com/casbin/casbin/v3"
    "github.com/your-username/your-repo"
)

a := yourpackage.NewAdapter(params)
e := casbin.NewEnforcer("examples/basic_model.conf", a)
```

## Migrate/Convert between different adapter

To migrate policies from adapter A to adapter B:

1.Load policy from A to memory

```go
e, _ := NewEnforcer(m, A)
```

or

```go
e.SetAdapter(A)
e.LoadPolicy()
```

2.Switch from adapter A to B

```go
e.SetAdapter(B)
```

3.Save policy from memory to B

```go
e.SavePolicy()
```

## Load/Save at run-time

Reload models and policies or persist policy changes after initialization:

```go
// Reload the model from the model CONF file.
e.LoadModel()

// Reload the policy from file/database.
e.LoadPolicy()

// Save the current policy (usually after changed with Casbin API) back to file/database.
e.SavePolicy()
```

## AutoSave

Adapters with Auto-Save capability can persist individual policy changes directly to storage without a full save operation
* This differs from `SavePolicy()`, which wipes storage and rewrites all policies, potentially causing performance issues with large policy sets.

When an adapter supports Auto-Save, control this behavior via `Enforcer.EnableAutoSave()`
* This option defaults to enabled for compatible adapters.

:::note

1. Auto-Save is an optional feature
* Adapters may implement it or not.
2. Auto-Save only functions when the enforcer's adapter supports it.
3. Check the AutoSave column in the adapter list above to determine support.

:::

Auto-Save usage example:

```go
import (
    "github.com/casbin/casbin/v3"
    "github.com/casbin/xorm-adapter"
    _ "github.com/go-sql-driver/mysql"
)

// AutoSave is enabled by default when using compatible adapters with enforcers.
a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
e := casbin.NewEnforcer("examples/basic_model.conf", a)

// Disable AutoSave.
e.EnableAutoSave(false)

// Policy changes affect only the in-memory enforcer,
// not the storage.
e.AddPolicy(...)
e.RemovePolicy(...)

// Enable AutoSave.
e.EnableAutoSave(true)

// Policy changes now persist to storage
// in addition to updating the in-memory enforcer.
e.AddPolicy(...)
e.RemovePolicy(...)
```

Additional examples: [https://github.com/casbin/xorm-adapter/blob/master/adapter_test.go](https://github.com/casbin/xorm-adapter/blob/master/adapter_test.go)

### How to write an adapter

Implement the [Adapter](https://github.com/casbin/casbin/blob/master/persist/adapter.go) interface with at least two required methods: `LoadPolicy(model model.Model) error` and `SavePolicy(model model.Model) error`.

Three optional methods enable Auto-Save support.

| Method                 | Type      | Description                                                |
|------------------------|-----------|------------------------------------------------------------|
| LoadPolicy()           | mandatory | Load all policy rules from the storage                     |
| SavePolicy()           | mandatory | Save all policy rules to the storage                       |
| AddPolicy()            | optional  | Add a policy rule to the storage                           |
| RemovePolicy()         | optional  | Remove a policy rule from the storage                      |
| RemoveFilteredPolicy() | optional  | Remove policy rules that match the filter from the storage |

:::note

When an adapter lacks Auto-Save support, provide empty implementations for the three optional methods
* Golang example:

:::

```go
// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
    return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
    return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
    return errors.New("not implemented")
}
```

Casbin enforcers ignore the "not implemented" error when calling these optional methods.

Adapter implementation requirements:

* Data Structure
  * Support reading at minimum six columns.
* Database Name
  * Default to `casbin`.
* Table Name
  * Default to `casbin_rule`.
* Ptype Column
  * Name as `ptype` (not `p_type` or `Ptype`).
* Table definition
  * `(id int primary key, ptype varchar, v0 varchar, v1 varchar, v2 varchar, v3 varchar, v4 varchar, v5 varchar)`.
* Unique key index
  * Build on columns `ptype,v0,v1,v2,v3,v4,v5`.
* `LoadFilteredPolicy` accepts a filter parameter structured like this:

```json
{
    "p": ["", "domain1"],
    "g": ["", "", "domain1"]
}
```

### Who is responsible to create the DB?

By convention, adapters should automatically create a `casbin` database if it doesn't exist and use it for policy storage
* Reference implementation: [https://github.com/casbin/xorm-adapter](https://github.com/casbin/xorm-adapter)

## Update Adapter

The `UpdateAdapter` interface extends the basic `Adapter` interface to enable direct policy updates in storage
* This approach is more efficient than the remove-and-add sequence when modifying existing rules.

Adapters implementing `UpdateAdapter` provide these methods:

| Method                   | Type     | Description                                                      |
|--------------------------|----------|------------------------------------------------------------------|
| UpdatePolicy()           | optional | Update a single policy rule in the storage                       |
| UpdatePolicies()         | optional | Update multiple policy rules in the storage                      |
| UpdateFilteredPolicies() | optional | Update policy rules that match the filter in the storage         |

### Example

Update adapter usage:

```go
import (
    "github.com/casbin/casbin/v3"
    "github.com/casbin/gorm-adapter/v3"
)

a, _ := gormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/")
e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

// Update a single policy
// Change: p, alice, data1, read -> p, alice, data1, write
e.UpdatePolicy(
    []string{"alice", "data1", "read"},
    []string{"alice", "data1", "write"},
)

// Update multiple policies at once
e.UpdatePolicies(
    [][]string{{"alice", "data1", "write"}, {"bob", "data2", "read"}},
    [][]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}},
)

// Update all policies matching a filter
e.UpdateFilteredPolicies(
    [][]string{{"alice", "data1", "write"}},
    0,
    "alice", "data1", "read",
)
```

### How to write an update adapter

To implement `UpdateAdapter`, add the update methods to your basic `Adapter` implementation:

```go
// UpdatePolicy updates a policy rule from storage.
// This is part of the UpdateAdapter interface.
func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
    // Update the policy in storage
    // SQL example: UPDATE casbin_rule SET v0=?, v1=?, v2=? WHERE ptype=? AND v0=? AND v1=? AND v2=?
    return nil
}

// UpdatePolicies updates multiple policy rules in the storage.
// This is part of the UpdateAdapter interface.
func (a *Adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
    // Update multiple policies in storage
    // Use transactions for consistency
    return nil
}

// UpdateFilteredPolicies updates policy rules that match the filter from the storage.
// This is part of the UpdateAdapter interface.
func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) error {
    // Find policies matching the filter, then update them
    return nil
}
```

:::note

Without `UpdateAdapter` support, Casbin automatically falls back to combining `RemovePolicy()` and `AddPolicy()` operations.

:::

## Context Adapter

[ContextAdapter](https://github.com/casbin/casbin/blob/master/persist/adapter_context.go) provides context-aware operations for Casbin adapters.

Context enables features like timeout control for adapter API calls.

### Example

[gormadapter](https://github.com/casbin/gorm-adapter) supports context
* Below is timeout control using context:

```go
ca, _ := NewContextAdapter("mysql", "root:@tcp(127.0.0.1:3306)/", "casbin")
// Set 300s timeout
ctx, cancel := context.WithTimeout(context.Background(), 300*time.Microsecond)
defer cancel()

err := ca.AddPolicyCtx(ctx, "p", "p", []string{"alice", "data1", "read"})
if err != nil {
    panic(err)
}
```

### How to write an context adapter

`ContextAdapter` API adds a context processing layer on top of standard `Adapter` API
* After implementing the standard Adapter API, wrap your logic with context handling.

Reference implementation: [adapter.go](https://github.com/casbin/gorm-adapter/blob/master/adapter.go)

## Transaction

Casbin supports transactions
* Transaction usage in `gormadapter`:

```go
db, _ := gorm.Open(...)
adapter, _ := gormadapter.NewTransactionalAdapterByDB(db)
e, _ := casbin.NewTransactionalEnforcer("examples/rbac_model.conf", adapter)

ctx := context.Background()

// WithTransaction executes a function within a transaction.
// Errors trigger rollback; otherwise, automatic commit occurs.
err := e.WithTransaction(ctx, func(tx *casbin.Transaction) error {
    tx.AddPolicy("alice", "data1", "read")
    tx.AddPolicy("alice", "data1", "write")
    return nil
})

// Manual transaction handling
tx, _ := e.BeginTransaction(ctx)
tx.AddPolicy("alice", "data1", "write")
if err := tx.Commit(); err != nil {
    // handle transaction failure
}
```

Implement `TransactionalAdapter` and `TransactionContext` from [persist/transaction.go](https://github.com/casbin/casbin/blob/master/persist/transaction.go) to add transaction support.

* [implementation | Go](https://github.com/casbin/gorm-adapter/blob/master/adapter.go)
