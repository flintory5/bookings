CREATE SCHEMA IF NOT EXISTS "instance";

CREATE  TABLE "instance".users ( 
	id                   integer  NOT NULL  ,
	first_name           varchar(255)  NOT NULL  ,
	last_name            varchar(255)  NOT NULL  ,
	email                varchar(255)    ,
	"password"           varchar(255)    ,
	created_at           timestamp    ,
	updated_at           timestamp    ,
	access_level         integer    ,
	CONSTRAINT pk_users PRIMARY KEY ( id )
 );