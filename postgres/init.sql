CREATE SCHEMA tenant;

CREATE SCHEMA common;

CREATE USER tenant_user WITH PASSWORD 'tenant';

CREATE USER common_user WITH PASSWORD 'common';

-- common_userはcommonスキーマのみアクセス可能
GRANT USAGE ON SCHEMA common TO common_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA common TO common_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA common GRANT
SELECT
, INSERT, UPDATE, DELETE ON TABLES TO common_user;

GRANT USAGE ON ALL SEQUENCES IN SCHEMA common TO common_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA common GRANT USAGE ON SEQUENCES TO common_user;

-- tenant_userはcommon, tenantスキーマ両方へアクセス可能
GRANT USAGE ON SCHEMA tenant TO tenant_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA tenant TO tenant_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA tenant GRANT
SELECT
, INSERT, UPDATE, DELETE ON TABLES TO tenant_user;

GRANT USAGE ON ALL SEQUENCES IN SCHEMA tenant TO tenant_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA tenant GRANT USAGE ON SEQUENCES TO tenant_user;

-- tenantスキーマのtable, view , sequenceへのアクセス権限を付与
GRANT USAGE ON SCHEMA common TO tenant_user;

-- tenantスキーマのCRUD権限を付与
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA common TO tenant_user;

-- 今後作成されるテーブルに対してもCRUD権限を付与
ALTER DEFAULT PRIVILEGES IN SCHEMA common GRANT
SELECT
, INSERT, UPDATE, DELETE ON TABLES TO tenant_user;

-- tenantスキーマのシーケンス権限を付与
GRANT USAGE ON ALL SEQUENCES IN SCHEMA common TO tenant_user;

-- 今後作成されるテーブルに対してもシーケンス権限を付与
ALTER DEFAULT PRIVILEGES IN SCHEMA common GRANT USAGE ON SEQUENCES TO tenant_user;

