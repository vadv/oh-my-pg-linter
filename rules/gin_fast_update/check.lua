local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на создание индекса
    if stmt.IndexStmt then
      if stmt.IndexStmt.accessMethod == "gin" then
        local fastupdate = true
        if stmt.IndexStmt.options then
          for _, param in pairs(stmt.IndexStmt.options) do
            if param.DefElem and param.DefElem.defname == "fastupdate" and param.DefElem.arg
                and param.DefElem.arg.String and param.DefElem.arg.String.str == "false" then
              fastupdate = false
            end
          end
        end
        if fastupdate then
          return "Запрос: "..statement:query()..[[

Использование gin-индекса без fastupdate приводит к росту pending_list.
Это может привести что рандомный запрос начнет тормозить и не успеет за таймаут перепаковать данные.

Решение:
Вместо: `create index on inventory using gin(groups);`.
Используйте: `create index on inventory using gin(groups) with (fastupdate = false);`.


Документация: https://postgrespro.ru/docs/postgrespro/9.5/gin-tips

]]
        end
      end
    end
  end
  return nil
end

return check
