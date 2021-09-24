SELECT
    u.nickname,
    ARRAY_AGG(t.tagname) AS tags
FROM
    users AS u

JOIN user_tags AS ut
    ON u.id = ut.user_id

JOIN tags AS t
    ON ut.tag_id = t.id

GROUP BY 
    u.nickname;
