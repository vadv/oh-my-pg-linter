local function check(tree)
  local result = {}
  local stmt
  for _, statement in pairs(tree) do
    -- проверяем что statement на drop индекса
    stmt = statement:tree()
    if stmt.DropStmt and (stmt.DropStmt.removeType == "OBJECT_INDEX") then
      if not(stmt.DropStmt.concurrent) then
        table.insert(result, statement)
      end
    end
  end
  return result
end

return check
