CREATE TABLE authors (
	id         INTEGER PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT NOW(),
  	updated_at timestamp,
	name       text NOT NULL,
	bio        text
);
