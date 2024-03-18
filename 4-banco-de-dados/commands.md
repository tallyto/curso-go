```bash
docker compose up -d
```

```bash
docker compose exec mysql bash
```

```bash
mysql -u root -p goexpert
``` 

```sql
create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));
```

```sql
show tables;
```

```sql
describe products;
```