-- Type: task_status

-- DROP TYPE public.task_status;

CREATE TYPE public.task_status AS ENUM
    ('NEW', 'WORKING', 'CANCELLED', 'FINISHED', 'OVERDUE');

ALTER TYPE public.task_status
    OWNER TO postgres;


-- Table: public.task

-- DROP TABLE public.task;

CREATE TABLE public.task
(
    id_ integer NOT NULL DEFAULT nextval('task_id__seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default",
    status task_status NOT NULL,
    duedate timestamp without time zone,
    CONSTRAINT task_pkey PRIMARY KEY (id_)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.task
    OWNER to postgres;