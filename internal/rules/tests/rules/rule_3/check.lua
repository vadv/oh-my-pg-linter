local function check(tree)
  local stmt
  for _, statement in pairs(tree) do
    -- проверяем что statement на drop индекса
    stmt = statement:tree()
    if stmt.DropStmt and (stmt.DropStmt.removeType == "OBJECT_INDEX") then
      if not(stmt.DropStmt.concurrent) then
        return statement:query()
      end
    end
  end
  return nil
end

return check
