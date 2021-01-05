# Penggunaan dan Dokumentasi


Semua kode belum dibuat unit test.

# Jalankan Aplikasi

silahkan setup .env sesuai database yang digunakan, Disini saya menggunakan Postgre
```
 go run main.go
```


# List Endpoint

- POST /login
  
  Response body ketika melakukan login adalah token JWT yang sudah digenerate
  ![image](https://user-images.githubusercontent.com/39885486/103661577-d82fa300-4fa9-11eb-842c-8bcb711a96a3.png)
  
- DELETE /users/{id}
  
  Untuk melakukan deleting user tersebut dibutuhkan token yang sudah digenerate pada saat login,
  ![image](https://user-images.githubusercontent.com/39885486/103661659-ee3d6380-4fa9-11eb-9640-f62eee8a31b4.png)
  
- PUT /users/{id} 

  Untuk melakukan UPDATE pada user. Dibutuhkan JWT sama seperti proses DELETE diatas
  ![image](https://user-images.githubusercontent.com/39885486/103661802-162cc700-4faa-11eb-9690-fab79116c5f4.png)
  ![image](https://user-images.githubusercontent.com/39885486/103661852-280e6a00-4faa-11eb-8cc5-addd22e47c15.png)

- GET /users

  endpoint ini mengembalikan semua user yang ada pada database, but the password still in the response body 😑
  ![image](https://user-images.githubusercontent.com/39885486/103661366-a4ed1400-4fa9-11eb-8e3f-2a25b14feeb2.png)
