local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на создание индекса
    if stmt.IndexStmt then
      if stmt.IndexStmt.if_not_exists and stmt.IndexStmt.concurrent then
        table.insert(result, statement)
      end
    end
  end
  return result
end

return check
