return {
  { sql = [[ CREATE INDEX "email_idx" ON "app_user" using gin(groups); ]], passed = false },
  { sql = [[ CREATE INDEX "email_idx" ON "app_user" using gin(groups) with (fastupdate = FALSE) ]], passed = true },
  { sql = [[ create index on inventory using gin(groups) with (fastupdate = false); ]], passed = true },
}
