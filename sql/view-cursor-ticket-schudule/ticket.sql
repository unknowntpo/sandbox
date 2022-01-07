create database tickets;

-- \c time;

drop table tickets;

create table tickets (
	id SERIAL Primary key, 
	due timestamp
);

insert into tickets (due) 
select timestamp '2014-01-10 10:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00') from generate_series(1, 10);
