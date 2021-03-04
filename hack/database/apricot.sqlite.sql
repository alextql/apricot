create table "short_url" (
	"id"        integer     not null constraint "short_url_id_pk" primary key autoincrement,
	"hash"      varchar(40) not null constraint "short_url_hash_pk" unique,
	"code"      varchar(16),
	"url"       text,
	"create_at" timestamp default CURRENT_TIMESTAMP
);

create unique index "short_url_hash_uindex" on "short_url"("hash");
create unique index "short_url_id_uindex" on "short_url"("id");
