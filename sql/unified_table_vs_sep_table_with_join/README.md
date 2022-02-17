# SQL Experiment: store data in same table v.s. seperate them and use join to combine the data

```sql
CREATE DATABASE uni-sep-table-join;

DROP TABLE IF NOT EXISTS ticket_key_and_data;

CREATE TABLE ticket_key_and_data (
    id SERIAL,
    ticket_data_a INTEGER,
    ticket_data_b INTEGER,
    ticket_data_c INTEGER,
)

DROP TABLE IF NOT EXISTS ticket_key;

CREATE TABLE ticket_key_and_data (
    id SERIAL,
);

DROP TABLE IF NOT EXISTS ticket_data;

CREATE TABLE ticket_data (
    id SERIAL,
    key_id INTEGER NOT NULL,
    ticket_data_a INTEGER,
    ticket_data_b INTEGER,
    ticket_data_c INTEGER,
    CONSTRAINT fk_key_id
        FOREIGN KEY(key_id)
            REFERENCES ticket_key(id)
);

-- How to insert key_id and ticket_data_[a|b|c] simultaneously ?
INSERT INTO ticket_data () 
SELECT 1, 2, 3 FROM generate_series(1, 100)
