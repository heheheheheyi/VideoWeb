create table videos
(
	id int unsigned auto_increment
		primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	title varchar(255) null,
	info varchar(255) null,
	url varchar(255) null,
	uid int unsigned null,
	img varchar(255) null,
	status int default 0 not null
)
charset=utf8mb4;

create index idx_videos_deleted_at
	on videos (deleted_at);

create index index_status
	on videos (status);

