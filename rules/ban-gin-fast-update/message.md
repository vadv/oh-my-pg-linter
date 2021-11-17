Использование gin-индекса без fastupdate приводит к росту pending_list.  Это может привести что рандомный запрос начнет тормозить и не успеет за таймаут перепаковать данные.

Решение:

Вместо:
```sql
create index on inventory using gin(groups);
```

Используйте: 
```sql
create index on inventory using gin(groups) with (fastupdate = false)`
```

[gin-tips](https://postgrespro.ru/docs/postgrespro/9.5/gin-tips)
