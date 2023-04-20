CREATE  TABLE "public".restrictions ( 
	id                   serial  NOT NULL  ,
	restriction_name     varchar(255)    ,
	created_at           timestamp    ,
	updated_at           timestamp    ,
	CONSTRAINT pk_restrictions PRIMARY KEY ( id )
 );
