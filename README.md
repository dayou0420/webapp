# WEBAPP
1. command
```
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```
```
psql -h localhost -p 5432 -U postgres -d postgres
```
```
insert into users (first_name, last_name, email, password, is_admin, created_at, updated_at) values ('Admin', 'User', 'admin@example.com', '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK', 1, '2023-03-15 00:00:00', '2023-02-15 00:00:00');
```