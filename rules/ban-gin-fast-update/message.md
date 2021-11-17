# Description

Using the gin index with fastupdate causes the `pending_list` to grow.
This can lead to a random query is slow down and not being able to repack gin-data within query context timeout. 

# Solution:

Instead of:
```sql
create index on inventory using gin(groups);
```

Use: 
```sql
create index on inventory using gin(groups) with (fastupdate = false)`
```

[gin-tips](https://postgrespro.ru/docs/postgrespro/9.5/gin-tips)
