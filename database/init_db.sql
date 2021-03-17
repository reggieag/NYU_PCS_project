create table data_table(
  id serial primary key,
  name varchar(256),
  quantity integer not null
);

create index concurrently "index_id_on_data_table"
on data_table using btree (id);
