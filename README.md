# Go Join Event

Go Join Event adalah sebuah REST API yang memungkinkan pengguna untuk melakukan registrasi acara dengan autentikasi akun. API ini dibangun menggunakan Go (Golang) dan beberapa package eksternal untuk mendukung fungsionalitasnya.

## Fitur Proyek

Fitur utama dari Go Join Event meliputi:

- Registrasi akun pengguna dengan keamanan hash password.
- Sistem login dengan Basic Auth untuk memperoleh token JWT.
- Permintaan token JWT untuk mengakses daftar event yang terdaftar oleh pengguna.
- Penyediaan daftar event yang telah diregistrasi oleh pengguna.
- Fitur CRUD dasar untuk setiap model yang terlibat.

## Package Eksternal yang Digunakan

- github.com/gin-gonic/gin v1.9.1
- github.com/golang-jwt/jwt/v5 v5.2.1
- github.com/joho/godotenv v1.5.1
- github.com/lib/pq v1.10.9
- github.com/rubenv/sql-migrate v1.6.1
- golang.org/x/crypto v0.22.0

## Endpoints API

### POST /register-user

Mendaftarkan akun pengguna.

**Contoh Request Body:**

```json
{
  "id": 1,
  "username": "John",
  "email": "john@email.com",
  "password": "icikiwir"
}
```

### GET /login-user

Login pengguna dengan Basic Auth dan meminta token JWT.

**Contoh Request Body:**

```json
{
  "username": "John",
  "password": "icikiwir"
}
```

### POST /register-event/:id

Mendaftarkan akun pengguna pada suatu event.

**Contoh Request Body:**

```json
{
  "user_id": 1,
  "event_id": 1
}
```

### GET /registered-events/:id

Mendaftarkan akun pengguna pada suatu event.

**Contoh Response:**

```json
{
  "data": [
    {
      "username": "John",
      "regis_date": "2024-04-27T21:09:27.606152Z",
      "title": "Coachella Valley Music and Arts Festival",
      "location": "Indio, California, USA",
      "start_date": "2024-04-19T00:00:00Z",
      "end_date": "2024-04-21T00:00:00Z"
    }
  ]
}
```
