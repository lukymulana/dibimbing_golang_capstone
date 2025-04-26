# Travel Booking System API Documentation

Sistem ini merupakan backend untuk aplikasi travel booking berbasis REST API menggunakan Golang.

---

## Daftar Route

### Public Routes (Tanpa Autentikasi)

#### 1. Register User
- **Endpoint:** `POST /register`
- **Deskripsi:** Mendaftarkan user baru.
- **Request Body:**
```json
{
  "username": "user1",
  "password": "yourpassword",
  "email": "user1@example.com",
  "role": "guide" // atau "traveler"
}
```
- **Response Sukses:**
```json
{
  "id": 1,
  "username": "user1",
  "email": "user1@example.com",
  "role": "guide"
}
```

#### 2. Login
- **Endpoint:** `POST /login`
- **Deskripsi:** Login dan mendapatkan JWT token.
- **Request Body:**
```json
{
  "username": "user1",
  "password": "yourpassword"
}
```
- **Response Sukses:**
```json
{
  "token": "<jwt_token>"
}
```

#### 3. Get Trips (Public)
- **Endpoint:** `GET /trips`
- **Deskripsi:** Mendapatkan daftar trip berdasarkan filter kota dan tanggal (opsional).
- **Query Params (opsional):**
  - `city`: Nama kota
  - `start_date`: Tanggal mulai (YYYY-MM-DD)
  - `end_date`: Tanggal akhir (YYYY-MM-DD)
- **Response Sukses:**
```json
[
  {
    "id": 1,
    "city": "Jakarta",
    "start_date": "2025-05-01",
    "end_date": "2025-05-05",
    "capacity": 10,
    "price": 500000
  }
]
```

---

### Authenticated Routes (Memerlukan JWT Token)
Semua route berikut harus menggunakan header:
```
Authorization: Bearer <jwt_token>
```

#### 4. Create Trip
- **Endpoint:** `POST /auth/trips`
- **Deskripsi:** Membuat trip baru.
- **Role:** Only **Guide**
- **Request Body:**
```json
{
  "city": "Jakarta",
  "start_date": "2025-05-01",
  "end_date": "2025-05-05",
  "capacity": 10,
  "price": 500000,
  "description": "Liburan seru di Jakarta"
}
```
- **Response Sukses:**
```json
{
  "id": 2,
  "city": "Jakarta",
  "start_date": "2025-05-01",
  "end_date": "2025-05-05",
  "capacity": 10,
  "price": 500000,
  "description": "Liburan seru di Jakarta"
}
```

#### 5. Update Trip
- **Endpoint:** `PUT /auth/trips/:id`
- **Deskripsi:** Mengedit trip yang dibuat oleh user login.
- **Role:** Only **Guide**
- **Request Body:** (sama seperti create)
- **Response Sukses:**
```json
{
  "id": 2,
  "city": "Jakarta",
  "start_date": "2025-05-01",
  "end_date": "2025-05-05",
  "capacity": 10,
  "price": 500000,
  "description": "Liburan seru di Jakarta"
}
```

#### 6. Delete Trip
- **Endpoint:** `DELETE /auth/trips/:id`
- **Deskripsi:** Menghapus trip yang dibuat oleh user login.
- **Role:** Only **Guide**
- **Response Sukses:**
```json
{
  "message": "Trip deleted successfully"
}
```

#### 7. Get My Trips
- **Endpoint:** `GET /auth/my-trips`
- **Deskripsi:** Mendapatkan daftar trip yang dibuat oleh user login.
- **Response Sukses:**
```json
[
  {
    "id": 2,
    "city": "Jakarta",
    "start_date": "2025-05-01",
    "end_date": "2025-05-05",
    "capacity": 10,
    "price": 500000,
    "description": "Liburan seru di Jakarta"
  }
]
```

#### 8. Create Booking (Traveler)
- **Endpoint:** `POST /auth/bookings`
- **Deskripsi:** Traveler melakukan booking untuk trip tertentu.
- **Role:** Only **Traveler**
- **Request Body:**
```json
{
  "TripID": 1
}
```
- **Response Sukses:**
```json
{
  "booking": {
    "BookingID": 1,
    "UserID": 2,
    "TripID": 1,
    "BookingStatus": "waiting",
    "CreatedAt": "2025-04-26T14:30:00+07:00"
  }
}
```
- **Validasi:**
  - Hanya user dengan role traveler yang dapat booking trip.
  - Booking hanya bisa dilakukan jika kuota trip masih tersedia (`sisa_kuota > 0`).
  - Jika kuota penuh akan mendapat error:
    ```json
    { "error": "Trip is fully booked" }
    ```
- **Catatan:**
  - Traveler dapat melihat sisa kuota trip pada endpoint detail trip (`GET /trips/:id`).

#### 10. Get Bookings by Guide (Konfirmasi Booking)
- **Endpoint:** `GET /auth/guide/bookings`
- **Deskripsi:** Guide melihat seluruh booking untuk trip yang dimilikinya (hasil join detail).
- **Autentikasi:** Hanya untuk user dengan role guide (Authorization: Bearer {{token}})
- **Response Sukses:**
```json
[
  {
    "booking_id": 1,
    "booking_status": "waiting",
    "created_at": "2025-04-26T14:30:00+07:00",
    "user_id": 2,
    "username": "traveler1",
    "email": "traveler1@mail.com",
    "trip_id": 1,
    "city": "Bandung",
    "start_date": "2025-05-01",
    "end_date": "2025-05-03"
  }
]
```

#### 11. Update Booking Status (Konfirmasi oleh Guide)
- **Endpoint:** `PUT /auth/guide/bookings/:booking_id`
- **Deskripsi:** Guide mengubah status booking (misal dari waiting menjadi success) untuk trip miliknya.
- **Autentikasi:** Hanya untuk user dengan role guide (Authorization: Bearer {{token}})
- **Request Body:**
```json
{
  "status": "success"
}
```
- **Validasi:**
  - Hanya guide pemilik trip yang bisa mengubah status booking
  - Status hanya bisa diubah ke `success` atau `waiting`
- **Response Sukses:**
```json
{
  "message": "Booking status updated successfully"
}
```

#### 9. Get Bookings by Trip
- **Endpoint:** `GET /auth/bookings/:trip_id`
- **Deskripsi:** Mendapatkan daftar booking berdasarkan trip id.
- **Response Sukses:**
```json
[
  {
    "id": 1,
    "trip_id": 1,
    "user_id": 2
  }
]
```

---

## Catatan
- Semua endpoint `/auth/*` membutuhkan JWT token pada header Authorization.
- Format tanggal mengikuti `YYYY-MM-DD`.
- Field pada request/response dapat berubah sesuai implementasi DTO/Entity di kode.

