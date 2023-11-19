-- create a psql table for projects
CREATE TABLE IF NOT EXISTS projects (
    p_id BIGSERIAL PRIMARY KEY,
    p_name TEXT NOT NULL,
    p_description TEXT NOT NULL,
    p_excerpt TEXT NOT NULL,
    p_category TEXT[] NOT NULL,
    p_assigned_to INT NOT NULL,
    p_created_by INT NOT NULL,
    p_created_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    p_due_date TIMESTAMP(0) with time zone NOT NULL,
    p_done BOOLEAN NOT NULL DEFAULT FALSE
);

