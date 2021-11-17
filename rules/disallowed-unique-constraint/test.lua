return {
  { sql = [[ ALTER TABLE table_name ADD CONSTRAINT field_name_constraint UNIQUE (field_name); ]], passed = false },
}
