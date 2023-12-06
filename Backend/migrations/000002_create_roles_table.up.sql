CREATE TABLE IF NOT EXISTS roles(
    r_id bigserial PRIMARY KEY,
    r_role text NOT NULL,
    r_created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    r_version integer NOT NULL DEFAULT 1
);
