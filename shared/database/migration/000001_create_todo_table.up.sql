SET search_path TO tenant;

CREATE TABLE todos(
  create_user_id uuid NOT NULL,
  update_user_id uuid NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  tenant_id uuid NOT NULL,
  id uuid PRIMARY KEY NOT NULL,
  title varchar(255) NOT NULL,
  description text,
  is_deleted bool NOT NULL
);

CREATE INDEX idx_todo_tenant_id ON todos(tenant_id);

