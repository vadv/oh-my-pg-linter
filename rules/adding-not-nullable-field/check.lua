--[[
{
  AlterTableStmt = {
    cmds = { {
        AlterTableCmd = {
          behavior = "DROP_RESTRICT",
          def = {
            ColumnDef = {
              colname = "foo",
              constraints = { {
                  Constraint = {
                    contype = "CONSTR_DEFAULT",
                    location = 52,
                    raw_expr = {
                      A_Const = {
                        location = 60,
                        val = {
                          Integer = {
                            ival = 10
                          }
                        }
                      }
                    }
                  }
                }, {
                  Constraint = {
                    contype = "CONSTR_NOTNULL",
                    location = 63
                  }
                } },
              is_local = true,
              location = 38,
              typeName = {
                location = 44,
                names = { {
                    String = {
                      str = "pg_catalog"
                    }
                  }, {
                    String = {
                      str = "int4"
                    }
                  } },
                typemod = -1
              }
            }
          },
          subtype = "AT_AddColumn"
        }
      } },
    relation = {
      inh = true,
      location = 13,
      relname = "core_recipe",
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
        if cmd.AlterTableCmd
            and cmd.AlterTableCmd.subtype == "AT_AddColumn"
            and cmd.AlterTableCmd.def
            and cmd.AlterTableCmd.def.ColumnDef
            and cmd.AlterTableCmd.def.ColumnDef.constraints then
          local add_default, not_null = false, false
          for _, constraint in pairs(cmd.AlterTableCmd.def.ColumnDef.constraints) do
            if constraint.Constraint.contype == "CONSTR_DEFAULT" then
              add_default = true
            end
            if constraint.Constraint.contype == "CONSTR_NOTNULL" then
              not_null = true
            end
          end
          if add_default and not_null then
            table.insert(result, statement)
          end
        end
      end
    end
  end
  return result
end

return check
