CREATE  TABLE "public".rooms ( 
	id                   serial  NOT NULL  ,
	room_name            varchar(255)    ,
	created_at           timestamp    ,
	updated_at           timestamp    ,
	CONSTRAINT pk_rooms PRIMARY KEY ( id )
 );