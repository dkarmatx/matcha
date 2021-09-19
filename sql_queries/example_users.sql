DROP TABLE IF EXISTS User_friends;
DROP TABLE IF EXISTS Users;
DROP DOMAIN IF EXISTS email_dom;
DROP DOMAIN IF EXISTS gender_dom;
DROP DOMAIN IF EXISTS sexpref_dom;

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
COMMENT ON DOMAIN sexpref_dom IS 'sex_prefs is a bitset where each "gender-bit" is set if user is interested in matches with that gender';

CREATE TABLE IF NOT EXISTS Users (
    id          bigserial       UNIQUE NOT NULL,

    nickname    varchar(255)    UNIQUE NOT NULL,
    email       email_dom       UNIQUE NOT NULL,
    passwd      varchar(255)    ,
    salt        varchar(255)    ,

    bio         text            DEFAULT '',
    birthdate   date            ,
    gender      gender_dom      ,
    sex_prefs   sexpref_dom     ,

    longitude   float,
    latitude    float,

    PRIMARY KEY (id, nickname, email)
);

INSERT INTO users (
    nickname,
    email,
    birthdate,
    bio,
    gender,
    sex_prefs
) VALUES
(
    'George Jilligan', 'georgejill@gmail.com',
    'January 9, 1987',
    'Good guy with good smile wants to find a cute GF for long relationships ;)',
    1, 2
),
(
    'Mr-Duffy-Beef', 'happy_duffy@yahoo.com',
    'July 17, 1994',
    'Hot man is looking for another hot man, write me a letter guys!!!',
    1, 1
),
(
    'Genderfluid Person', 'someone-wants-to-fly-887@students.oxford.edu',
    'May 1, 2001',
    'I will tell you who am I, but after you tell me, who you are. ;)))',
    4, 7
),
(
    'Helicopter Girl', 'warfare_danger_1995@gmail.com',
    'June 27, 1995',
    'I like helicopters, looking for an amigo for helicopter races',
    2, 1
),
(
    'Sventala', 'sveta-rus11@inbox.ru',
    'August 12, 1983',
    'Looking for a hot girl which is hugry for love and sex',
    2, 2
);

CREATE TABLE IF NOT EXISTS User_friends (
    id          bigserial       UNIQUE NOT NULL,
    user_id     bigint          NOT NULL,
    friend_id   bigint          NOT NULL,

    PRIMARY KEY (id, user_id, friend_id),
    FOREIGN KEY (user_id)       REFERENCES Users (id) ON DELETE CASCADE,
    FOREIGN KEY (friend_id)     REFERENCES Users (id) ON DELETE CASCADE
);

INSERT INTO User_Friends (user_id, friend_id) VALUES
(1, 4),
(2, 1),
(4, 1),
(3, 5),
(1, 5),
(5, 4);
