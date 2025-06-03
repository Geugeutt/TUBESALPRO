# 📈 Aplikasi Manajemen Investasi 

Aplikasi ini dibuat untuk membantu pengguna dalam mengatur portofolio investasinya secara sederhana.  
Dibuat pakai bahasa Go sebagai bentuk implementasi materi-materi yang udah kami pelajari selama perkuliahan.

## 🧠 Deskripsi Singkat

Aplikasi berbasis CLI (Command Line Interface) ini punya beberapa fitur dasar kayak ngatur portofolio, lihat harga aset, sampe nyimpen data user. Walaupun simple, tapi udah cukup buat simulasi manajemen investasi mini.

## 🔧 Fitur-Fitur

- 👤 **Profil Pengguna** – Menyimpan dan menampilkan info user.
- 💼 **Portofolio Investasi** – Bisa tambah, hapus, dan lihat aset yang dimiliki.
- 📦 **Katalog Aset** – Lihat aset-aset yang tersedia.
- 💰 **Harga Aset** – Lihat & kelola harga saham, reksadana, obligasi, dll.
- 📊 **Laporan Ringkasan** – Menampilkan total keuntungan/kerugian.

## 🗂 Struktur File 

```plaintext
fiturMenu.go          // Menu utama program
fiturProfil.go        // Buat ngatur data profil user
fiturPortofolio.go    // Fungsi CRUD portofolio
fiturKatalog.go       // Nampilin katalog aset
isidbHargaaset.go     // Data harga aset (dummy)
isidbPortofolio.go    // Data portofolio awal (dummy)
dataBentukan.go       // Struct data yang kita pake
helper.go             // Fungsi bantu kayak input & validasi
database.go           // Simulasi database (tanpa SQL)
go.mod                // Modul Go

