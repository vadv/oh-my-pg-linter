Использование gin-индекса без fastupdate приводит к росту pending_list.
Это может привести что рандомный запрос начнет тормозить и не успеет за таймаут перепаковать данные.

Решение:
Вместо: `create index on inventory using gin(groups);`.
Используйте: `create index on inventory using gin(groups) with (fastupdate = false);`.


Документация: https://postgrespro.ru/docs/postgrespro/9.5/gin-tips
