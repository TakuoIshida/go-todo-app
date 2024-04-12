CREATE SCHEMA tenant;

CREATE SCHEMA common;

CREATE USER tenant_user WITH PASSWORD 'tenant';

CREATE USER common_user WITH PASSWORD 'common';

GRANT usage ON SCHEMA common TO tenant_user;

GRANT usage ON SCHEMA common TO common_user;

GRANT usage ON SCHEMA tenant TO tenant_user;

GRANT usage ON SCHEMA tenant TO common_user;

