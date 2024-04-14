SET search_path TO common;

CREATE TABLE tenants(
  create_user_id uuid NOT NULL,
  update_user_id uuid NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  id uuid PRIMARY KEY NOT NULL,
  tenant_name varchar(255) NOT NULL,
  is_deleted bool NOT NULL
);

