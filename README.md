# dts-gol3-6-project2
Project individu digitalent kominfo batch 3 kelas 6 - Saeful Rizal MM
simple CRUD dengan golang dan html ajax.

installasi
- update mysql.go sesuaikan koneksi dan databasenya
- create database sesuai yang didaftarkan di mysql.go (contoh : dts_gol3_6_project2)
- create table task seperti berikut 
    CREATE TABLE task (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255),
    employee VARCHAR(30),
    deadline DATE,
    status INT(6)
    )
- go build
- .\dts-gol3-6-project2

