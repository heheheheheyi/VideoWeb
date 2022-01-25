create table comments
(
	id int unsigned auto_increment
		primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	info varchar(255) null,
	vid int unsigned null,
	uid int unsigned null,
	status int default 0 not null
)
charset=utf8mb4;

create index idx_comments_deleted_at
	on comments (deleted_at);

create index index_status
	on comments (status);

