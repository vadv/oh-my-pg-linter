return {
  { sql = [[ CREATE INDEX CONCURRENTLY IF NOT EXISTS "email_idx" ON "app_user" ("email"); ]], passed = false },
  { sql = [[ CREATE INDEX CONCURRENTLY "email_idx" ON "app_user" ("email"); ]], passed = true },
}
