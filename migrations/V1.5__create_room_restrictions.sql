CREATE  TABLE "public".room_restrictions ( 
	id                   serial  NOT NULL  ,
	start_date           date    ,
	end_date             date    ,
	room_id              integer  NOT NULL  ,
	reservation_id       integer  NOT NULL  ,
	created_at           timestamp    ,
	updated_at           timestamp    ,
	restriction_id       integer  NOT NULL  ,
	CONSTRAINT pk_room_restrictions PRIMARY KEY ( id )
 );