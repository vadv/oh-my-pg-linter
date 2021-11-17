return {
  { sql = [[ DROP INDEX "email_idx"; ]], passed = false },
  { sql = [[ DROP INDEX CONCURRENTLY "email_idx"; ]], passed = true },
}
