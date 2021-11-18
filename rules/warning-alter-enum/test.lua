return {
  { sql = [[ alter type job_type add value 'NewJobType' ]], passed = false },
  { sql = [[ ALTER TYPE colors RENAME VALUE 'purple' TO 'mauve' ]], passed = false },
}
