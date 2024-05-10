BEGIN;

CREATE TABLE public.car 
(
  PRIMARY KEY (id),
  id          serial    NOT NULL,
  regnum      varchar(20)  NOT NULL,
  mark        varchar(50)  NOT NULL,
  model       varchar(50)  NOT NULL,
  year        smallint,
  name        varchar(50)  NOT NULL,
  surname     varchar(50)  NOT NULL,
  patronymic  text
);

COMMIT;
