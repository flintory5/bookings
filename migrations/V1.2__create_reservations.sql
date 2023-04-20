CREATE  TABLE reservations ( 
	id                   serial  NOT NULL  ,
	first_name           varchar(255) DEFAULT ''   ,
	last_name            varchar(255) DEFAULT ''   ,
	email                varchar(255)    ,
	phone                varchar(255) DEFAULT ''   ,
	start_date           date    ,
	end_date             date    ,
	room_id              integer  NOT NULL  ,
	created_at           timestamp    ,
	updated_at           timestamp    ,
	CONSTRAINT pk_reservations PRIMARY KEY ( id )
 );