--[[
{
  AlterTableStmt = {
    cmds = { {
        AlterTableCmd = {
          behavior = "DROP_RESTRICT",
          def = {
            Constraint = {
              conname = "field_name_constraint",
              contype = "CONSTR_UNIQUE",
              keys = { {
                  String = {
                    str = "field_name"
                  }
                } },
              location = 28
            }
          },
          subtype = "AT_AddConstraint"
        }
      } },
    relation = {
      inh = true,
      location = 13,
      relname = "table_name",
      relpersistence = "p"
    },
    relkind = "OBJECT_TABLE"
  }
}

--]]
local function check(tree)
  local stmt
  local result = {}
  for _, statement in pairs(tree) do
    stmt = statement:tree()
    -- проверяем что statement на alter table
    if stmt.AlterTableStmt and stmt.AlterTableStmt.cmds then
      for _, cmd in pairs(stmt.AlterTableStmt.cmds) do
        if cmd.AlterTableCmd and cmd.AlterTableCmd.def and cmd.AlterTableCmd.def.Constraint then
          if cmd.AlterTableCmd.def.Constraint.contype == "CONSTR_UNIQUE" then
            table.insert(result, statement)
          end
        end
      end
    end
  end
  return result
end

return check
