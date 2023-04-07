CREATE INDEX room_restrictions_dates_idx ON "public".room_restrictions ( start_date, end_date );
CREATE INDEX room_restrictions_room_id_idx ON "public".room_restrictions ( room_id );
CREATE INDEX room_restrictions_reservation_id_idx ON "public".room_restrictions ( reservation_id );