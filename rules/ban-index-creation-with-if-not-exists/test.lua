return {
  { sql = [[ CREATE INDEX IF NOT EXISTS "email_idx" ON "app_user" ("email"); ]], passed = false },
  { sql = [[ CREATE INDEX "email_idx" ON "app_user" ("email"); ]], passed = true },
}
