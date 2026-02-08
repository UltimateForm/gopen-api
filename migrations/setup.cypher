CREATE CONSTRAINT user_email_is_unique
FOR (u:User)
REQUIRE u.email IS UNIQUE;

CREATE CONSTRAINT character_name_is_unique
FOR (c:Character)
REQUIRE c.name IS UNIQUE;