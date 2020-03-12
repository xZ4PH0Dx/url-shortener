CREATE TABLE urls(id serial primary key, code char(50), original_url char(255));
CREATE TABLE request_log(id serial primary key, url_id int, user_ip char(16));

CREATE PROCEDURE insert_data(short_code char(50), url char(255), ip char(16))
LANGUAGE SQL
AS $$
INSERT INTO urls(code, original_url) VALUES (short_code, url);
INSERT INTO request_log(url_id, user_ip) VALUES (LASTVAL(), ip);
$$;
