return {
  { sql = [[ DROP INDEX "email_idx_not_concurrently"; ]], passed = false },
  { sql = [[ DROP INDEX CONCURRENTLY "email_idx"; ]], passed = true },
}
