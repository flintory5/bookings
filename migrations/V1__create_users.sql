CREATE  TABLE users ( 
	id                   integer  NOT NULL  ,
	first_name           varchar(255) DEFAULT '' NOT NULL  ,
	last_name            varchar(255) DEFAULT '' NOT NULL  ,
	email                varchar(255)    ,
	"password"           varchar(255)    ,
	access_level         integer DEFAULT 1   ,
	CONSTRAINT pk_users PRIMARY KEY ( id )
 );