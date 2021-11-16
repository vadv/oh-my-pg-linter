local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на создание индекса
    if stmt.IndexStmt then
      if stmt.IndexStmt.if_not_exists and stmt.IndexStmt.concurrent then
        return "Запрос: "..statement:query()..[[
Нельзя использовать конструкцию `if not exists` совместно с `concurrently`.
Это может привести к проблеме не идемпотентности миграции так как повторный запуск миграции
может привести к проблеме когда index будет в state `invalid` и не будет создан.

]]
      end
    end
  end
  return nil
end

return check
