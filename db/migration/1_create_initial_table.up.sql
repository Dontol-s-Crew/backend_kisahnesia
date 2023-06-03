-- Adminer 4.8.1 PostgreSQL 14.5 (Debian 14.5-1.pgdg110+1) dump
CREATE OR REPLACE FUNCTION update_modified_column()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "user" (
    "id" bigint DEFAULT nextval('user_id_seq') NOT NULL,
    "email" character varying(255) UNIQUE NOT NULL,
    "password" text NOT NULL,
    "nama" character varying(255) NOT NULL,
    "is_admin" boolean default false,
    "time_updated" timestamp default current_timestamp,
    "time_created" timestamp default current_timestamp,
    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TRIGGER update_timestamp_trigger
    BEFORE UPDATE ON "user"
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE SEQUENCE cerita_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "cerita" (
    "id" bigint DEFAULT nextval('cerita_id_seq') NOT NULL,
    "user_id" bigint NOT NULL,
    "populer" bigint default 0,
    "ilutrasi" text NOT NULL,
    "cover" text NOT NULL,
    "daerah" character varying(255) NOT NULL,
    "judul" character varying(255) NOT NULL,
    "genre" character varying(255) NOT NULL,
    "status" boolean default false,
    "time_updated" timestamp default current_timestamp,
    "time_created" timestamp default current_timestamp,
    CONSTRAINT "cerita_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TRIGGER update_timestamp_trigger
    BEFORE UPDATE ON "cerita"
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE SEQUENCE isi_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "isi" (
    "id" bigint DEFAULT nextval('isi_id_seq') NOT NULL,
    "cerita_id" bigint NOT NULL,
    "order" bigint NOT NULL,
    "paragraft" text NOT NULL,
    CONSTRAINT "isi_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE ONLY cerita ADD CONSTRAINT "cerita_user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

ALTER TABLE ONLY isi ADD CONSTRAINT "isi_cerita_id_fkey" FOREIGN KEY (cerita_id) REFERENCES cerita(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

-- 2023-05-25 15:47:55.62181+00
