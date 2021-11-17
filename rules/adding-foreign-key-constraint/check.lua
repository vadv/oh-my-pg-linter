--[[
{
  AlterTableStmt = {
    cmds = { {
        AlterTableCmd = {
          behavior = "DROP_RESTRICT",
          def = {
            Constraint = {
              conname = "fk_user",
              contype = "CONSTR_FOREIGN",
              fk_attrs = { {
                  String = {
                    str = "user_id"
                  }
                } },
              fk_del_action = "a",
              fk_matchtype = "s",
              fk_upd_action = "a",
              location = 25,
              pk_attrs = { {
                  String = {
                    str = "id"
                  }
                } },
              pktable = {
                inh = true,
                location = 85,
                relname = "user",
                relpersistence = "p"
              },
              skip_validation = true
            }
          },
          subtype = "AT_AddConstraint"
        }
      } },
    relation = {
      inh = true,
      location = 13,
      relname = "email",
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
        if cmd.AlterTableCmd and cmd.AlterTableCmd.def and cmd.AlterTableCmd.def.Constraint
            and cmd.AlterTableCmd.subtype == "AT_AddConstraint"
            and cmd.AlterTableCmd.def.Constraint.contype == "CONSTR_FOREIGN" then
            if not cmd.AlterTableCmd.def.Constraint.skip_validation then
              table.insert(result, statement)
            end
        end
      end
    end
  end
  return result
end

return check
