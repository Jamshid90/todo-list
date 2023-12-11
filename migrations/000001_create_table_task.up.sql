CREATE TABLE IF NOT EXISTS "task" (
    "uuid"        UUID NOT NULL,
    "title"       CHARACTER VARYING(300) DEFAULT '',
    "priority"    CHARACTER VARYING(20) DEFAULT '',
    "status"    CHARACTER VARYING(20) DEFAULT '',
    "description" CHARACTER VARYING(500) DEFAULT '',
    "created_at"  TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT task_uuid_pkey PRIMARY KEY (uuid)
);