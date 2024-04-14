SET search_path TO 'tenant';

ALTER TABLE todos ENABLE ROW LEVEL SECURITY;

CREATE POLICY todo_policy ON todos
  USING (tenant_id = current_setting('app.tenant_id')::uuid);

