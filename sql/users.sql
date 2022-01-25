create table users
(
	id int unsigned auto_increment
		primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	account varchar(255) null,
	nickname varchar(255) null,
	password varchar(255) null,
	img varchar(255) null,
	status int default 0 not null
)
charset=utf8mb4;

create index idx_users_deleted_at
	on users (deleted_at);

create index index_status
	on users (status);

