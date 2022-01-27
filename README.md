<div id="top"></div>

  <h1 align="center">Test Bank Backend Engineer</h1>

</div>

<!-- DOCUMENTATION PROJECT -->
## Documentasi Project

### * Menjalankan program
- Akses link <b> https://labs.play-with-docker.com</b> buat instance disana. Setela berhasil clone repository pada intance tersebut
```
$ git clone https://github.com/NurilH/BANK-NURIL.git
```
- Setelah berhasil clone. masuk ke directory.
```
$ cd BANK-NURIL
```
- Setelah itu buat images docker seperti berikut
```
$ docker build -t app-bank:latest .
```
- Selalnjutnya jalankan docker compose
```
$ docker-compose up
```

  setelah program berhasil berjalan, akses sesuai port yang dibuka (80) dengan menambahkan <b>endpoint</b> yang telah dibuat melalui <b>"browser"</b> atau aplikasi API testing seperti <b>"Postman"</b>
<br/>

Berikut fitur dan endpoint Rest API yang terdapat dalam project ini :
| Method | Feature | Endpoint | Authentication
|---|---|---|---|
| POST | Menambahkan data nasabah | "link sesuai port"/nasabah | No |
| GET | Melihat data nasabah | "link sesuai port"/nasabah | No |
| GET | Melihat data nasabah berdasarkan no ktp | "link sesuai port"/nasabah/<no_ktp> | No |
| PUT | Melakukan update data nasabah | "link sesuai port"/nasabah/<id_nasabah> | No |
| DELETE | Menghapus data nasabah | "link sesuai port"/nasabah/<id_nasabah> | No |



<br/>
 
### Built With

![VS Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=FFFFFF)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=FFFFFF)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=FFFFFF)&nbsp;

<p align="right">(<a href="#top">back to top</a>)</p>

Contributor :
<br>
<p align="left">
    <a href="https://www.linkedin.com/in/nuril-huda-87b279214/" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg"
            alt="nuril huda" height="30" width="40" /></a>
    <a href="https://www.hackerrank.com/nurilhuda7337" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/hackerrank.svg"
            alt="nuril7337" height="30" width="40" /></a>
</p>

<p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: 2022 | Built with :heart: from us</p>
</h3>
<!-- end -->
