create table "short_url" (
	"id"        integer     not null constraint "short_url_id_pk" primary key autoincrement,
	"hash"      varchar(40) not null constraint "short_url_hash_pk" unique,
	"code"      varchar(16),
	"url"       text,
	"create_at" timestamp default CURRENT_TIMESTAMP
);

create unique index "short_url_hash_uindex" on "short_url"("hash");
create unique index "short_url_id_uindex" on "short_url"("id");

create table "mock_server" (
	"id"        integer     not null constraint "mock_server_id_pk" primary key autoincrement,
	"route"     varchar(64) not null constraint "mock_server_pk" unique,
	"body"      text        not null,
	"create_at" timestamp default CURRENT_TIMESTAMP
);

create unique index "mock_server_id_uindex" on "mock_server"("id");
create unique index "mock_server_route_uindex" on "mock_server"("route");


