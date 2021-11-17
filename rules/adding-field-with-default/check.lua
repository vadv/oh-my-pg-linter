local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на alter table
    if stmt.AlterTableStmt then
      for _, cmd in pairs(stmt.AlterTableStmt.cmds) do
        if cmd.AlterTableCmd and cmd.AlterTableCmd and cmd.AlterTableCmd.def and cmd.AlterTableCmd.def.ColumnDef
            and cmd.AlterTableCmd.def.ColumnDef.constraints then
          for _, constraint in pairs(cmd.AlterTableCmd.def.ColumnDef.constraints) do
            if constraint.Constraint and constraint.Constraint.contype == "CONSTR_NOTNULL" then
              table.insert(result, statement)
            end
          end
        end
      end
    end
  end
  return result
end

return check
