local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на drop индекса
    if stmt.DropStmt and (stmt.DropStmt.removeType == "OBJECT_INDEX") then
      if not(stmt.DropStmt.concurrent) then
        return "Запрос: "..statement:query()..[[
Удаление индекса - потенциально тяжелая операция, при выполнении запроса при создании индекса доступ к таблице
без ключевого слова `concurrently` будет заблокирован.

Решение:
Вместо: `DROP INDEX "email_idx";`.
Используйте: `DROP INDEX CONCURRENTLY "email_idx";`.

Документация: https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

]]
      end
    end
  end
  return nil
end

return check
