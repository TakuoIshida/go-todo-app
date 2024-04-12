SET search_path TO 'tenant';

ALTER TABLE todo ENABLE ROW LEVEL SECURITY;

CREATE POLICY todo_policy ON todo
  USING (tenant_id = current_setting('app.current_tenant')::uuid);

