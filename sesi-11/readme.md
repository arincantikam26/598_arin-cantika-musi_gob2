# SESI 11
- kode peserta    : 149368582100-598
- nama lengkap    : Arin Cantika Musi

---
## Summary sesi 11 : 

1. Unit-test
Bahasa Go telah menyediakan suatu package bernama testing, yang dapat kita gunakan untuk membuat padaaplikasi bahasa Go kita. Terdapat aturan dari penamaan file test yang harus kita ikuti guna untuk membuat filetesting pada bahasa Go. 

2. Create Failed Unit-test
    - t.Fail() : Digunakan untuk menggagalkan sebuah unit test namun proses eksekusi dari unit  test nya tidak akanterhenti.
    - t.FailNow() : Digunakan untuk menggagalkan sebuah unit test namun dan proses eksekusi nya akan terhenti pada unittest tersebut dan unit test lainnya tidak akan tereksekusi.
    - t.Error(...args) : Digunakan untuk menggagalkan sebuah unit test sekaligus melakukan logging dari errornya agar kitadapat membuat pesan dari error yang terjadi. 
    - t.Fatal(...args) : Fungsi nya mirip seperti  
    - t.FailNow(), namun t.Fatal(...args) dapat melakukan logging dari pesanerrornya.

3. Mock
Mock merupakan sebuah object yang sudah kita program dengan eksepektasi tertentu sehingga ketika objecttersebutdipanggil, maka ia akan menghasilkan data sesuai dengan yang sudah kita progam diawal.
