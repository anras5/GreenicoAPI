--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0
-- Dumped by pg_dump version 15.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: PICO; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."PICO" (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public."PICO" OWNER TO postgres;

--
-- Name: PICO_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."PICO_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."PICO_id_seq" OWNER TO postgres;

--
-- Name: PICO_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."PICO_id_seq" OWNED BY public."PICO".id;


--
-- Name: READINGS; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."READINGS" (
    id integer NOT NULL,
    temperature numeric NOT NULL,
    humidity numeric NOT NULL,
    pressure numeric NOT NULL,
    voc numeric NOT NULL,
    uv numeric NOT NULL,
    light numeric NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    pico_id integer NOT NULL
);


ALTER TABLE public."READINGS" OWNER TO postgres;

--
-- Name: READINGS_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."READINGS_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."READINGS_id_seq" OWNER TO postgres;

--
-- Name: READINGS_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."READINGS_id_seq" OWNED BY public."READINGS".id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: PICO id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PICO" ALTER COLUMN id SET DEFAULT nextval('public."PICO_id_seq"'::regclass);


--
-- Name: READINGS id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."READINGS" ALTER COLUMN id SET DEFAULT nextval('public."READINGS_id_seq"'::regclass);


--
-- Name: PICO PICO_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."PICO"
    ADD CONSTRAINT "PICO_pkey" PRIMARY KEY (id);


--
-- Name: READINGS READINGS_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."READINGS"
    ADD CONSTRAINT "READINGS_pkey" PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: READINGS READINGS_PICO_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."READINGS"
    ADD CONSTRAINT "READINGS_PICO_id_fk" FOREIGN KEY (pico_id) REFERENCES public."PICO"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

