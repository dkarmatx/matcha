package install

const _SQL_CHECKS_SCRIPT = `
	SELECT schema_name, schema_owner
	FROM information_schema.schemata
	WHERE schema_name = '[SCHEMA_NAME]' AND schema_owner = '[SCHEMA_OWNER]' ;
`

const _SQL_DROP_SCRIPT = `
	DROP SCHEMA IF EXISTS [SCHEMA_NAME] CASCADE ;
`

const _SQL_SCHEMA_SCRIPT = `
	CREATE SCHEMA [SCHEMA_NAME] ;
	SET search_path TO [SCHEMA_NAME] ;
`

const _SQL_TYPES_SCRIPT = `
	CREATE DOMAIN email_dom AS
		varchar(1024)
		CONSTRAINT valid_email CHECK (VALUE ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9_-]+([.][A-Za-z_-]+)+$');
	COMMENT ON DOMAIN email_dom IS 'Type of email with auto cheking';

	CREATE DOMAIN gender_dom AS
		int    DEFAULT 4
		CONSTRAINT valid_gender CHECK(VALUE = 1 OR VALUE = 2 OR VALUE = 4);
	COMMENT ON DOMAIN gender_dom  IS '1 (001) - male, 2 (010) - female, 4 (100) - Other';

	CREATE DOMAIN sexpref_dom AS
		int    DEFAULT 7
		CONSTRAINT valid_sexpref CHECK((VALUE & ~(7)) = 0);
	COMMENT ON DOMAIN sexpref_dom IS 'sexprefs is a bitset where each "gender-bit" is set if user is interested in matches with that gender';
`

const _SQL_TABLES_SCRIPT = `
	CREATE TABLE users (
		user_id         bigserial       UNIQUE NOT NULL,
		user_nickname   varchar(255)    UNIQUE NOT NULL,
		user_email      email_dom       UNIQUE NOT NULL,

		bio             text            DEFAULT '',
		birthdate       date            ,
		gender          gender_dom      ,
		sexpref         sexpref_dom     ,

		PRIMARY KEY (user_id, user_nickname, user_email)
	);

	CREATE TABLE tags (
		tag_id      bigserial       UNIQUE NOT NULL,
		tag_name    varchar(255)    UNIQUE NOT NULL,

		PRIMARY KEY (tag_id, tag_name)
	);

	CREATE TABLE user_tags (
		user_id     bigint          NOT NULL,
		tag_id      bigint          NOT NULL,

		PRIMARY KEY (user_id, tag_id),
		FOREIGN KEY (user_id)       REFERENCES users (user_id) ON DELETE CASCADE,
		FOREIGN KEY (tag_id)        REFERENCES tags (tag_id) ON DELETE CASCADE
	);
`
