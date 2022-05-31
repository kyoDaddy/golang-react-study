# go-react-backend-app

### etc
go get -u github.com/julienschmidt/httprouter

#### run
go run ./api/*.go

#### action *.sql
psql -d db1 -U userA -f /path/xxx.sql
go get github.com/graphql-go/graphql

#### install postgresql 
go get -u github.com/lib/pq

```sql
select * from public.genres;
select * from public.movies;

insert into public.genres (id, genre_name, created_at, updated_at) values (1,	'Drama',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (2,	'Crime',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (3,	'Action',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (4,	'Comic Book',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (5,	'Sci-Fi',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (6,	'Mystery',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (7,	'Adventure',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (8,	'Comedy',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.genres (id, genre_name, created_at, updated_at) values (9,	'Romance',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');

insert into public.movies (id, title, description, "year", release_date, runtime, rating, mpaa_rating, created_at, updated_at) values (1,	'The Shawshank Redemption',	'Two imprisoned men bond over a number of years',	'1994',	'1994-10-14',	'142',	'5',	'R',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.movies (id, title, description, "year", release_date, runtime, rating, mpaa_rating, created_at, updated_at) values (2,	'The Godfather',	'The aging patriarch of an organized crime dynasty transfers control to his son',	'1972',	'1972-03-24',	'175',	'5',	'R',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.movies (id, title, description, "year", release_date, runtime, rating, mpaa_rating, created_at, updated_at) values (4,	'American Psycho',	 'A wealthy New York investment banking executive hides his alternate psychopathic ego'	,'2000',	'2000-04-14',	'102',	'4',	'R',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');
insert into public.movies (id, title, description, "year", release_date, runtime, rating, mpaa_rating, created_at, updated_at) values (3,	'The Dark Knight',	'The menace known as the Joker wreaks havoc on Gotham City',	'2008',	'2008-07-18',	'152',	'5'	,'PG13',	'2021-05-17 00:00:00',	'2021-05-17 00:00:00');

alter table public.movies alter column id set default nextval('public.movies_id_seq'::regclass);
```

#### jwt
go get github.com/pascaldekloe/jwt
go get -u golang.org/x/crypto/bcrypt
go get github.com/justinas/alice



#### env
```text
export GO_MOVIES_JWT='2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160'
// read jwt secret from env
cfg.jwt.secret = os.Getenv("GO_MOVIES_JWT")
```

#### build
```shell
env GOOS=linux GOARCH=amd64 go build -o gomovies ./cmd/api 
```

#### profile viper
```shell
go get github.com/spf13/viper
```





