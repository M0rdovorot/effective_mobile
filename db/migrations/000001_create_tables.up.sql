BEGIN;

CREATE TABLE public.car 
(
  PRIMARY KEY (id),
  id          serial    NOT NULL,
  regNum      char(20)  NOT NULL,
  mark        char(50)  NOT NULL,
  model       char(50)  NOT NULL,
  year        smallint,
  name        char(50)  NOT NULL,
  surname     char(50)  NOT NULL,
  patronymic  text
);

COMMIT;
