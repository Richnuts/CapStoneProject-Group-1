
# Safety Protocol App
<br></br>

<p align="center">
<img src="https://safety-protocol-cyan.vercel.app/static/media/logo.81a1f7af00ea4bce08e00455cc40f4f5.svg"
</p><br></br><br></br>

## 💻 &nbsp;About The Project
Safety Protocol App is an app that help us to track person who will came to work at the office.
<details>
<summary>📖 ERD</summary>
<ul>
<img src="group_project_capstone_ERD.jpg">
</ul>
</details>

<details>
<summary>📖 MYSQL Schema</summary>
<ul>
<img src="Mysql_Schema.jpg">
</ul>
</details>

<details>
<summary>🛠️ Features</summary>
<ul>

<details>
<summary>🔒 &nbsp;Authentication</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| POST | /register | - | - | NO | NO | Register akun user / pegawai |
| POST | /login  | - | - | NO | NO | Login ke dalam sistem |
</details>


<details>
<summary>👨‍💼 &nbsp;Users</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| GET | /profile | - | - | YES | NO | Get data user yang sedang login |
| GET | /users/:id  | - | - | YES | YES | Get data user tertentu |
| PUT | /users/:id | - | name, email, password, image | YES | NO | Edit data user |
| DELETE | /users/:id  | - | - | YES | NO | Delete data user |
</details>



<details>
<summary>🏢 &nbsp;Office</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| GET | /offices | - | - | YES | NO | Get list data office |
| GET | /offices/:id  | - | - | YES | NO | Get data office tertentu |
</details>



<details>
<summary>🗓️ &nbsp;Schedule</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| GET | /schedules | page, month, year, office | - | YES | NO | Get list data schedule untuk WFO |
| POST | /schedules  | - | office_id, total_capacity, month, year | YES | YES | Menambahkan data schedule di office, bulan dan tahun tertentu |
| GET | /schedules/:id | page | - | YES | NO | Get data schedule beserta partisipannya |
| PUT | /schedules/:id  | - | total_capacity | YES | YES | Edit total capacity pada sebuah schedule |
</details>

 

<details>
<summary>📃 &nbsp;Certificates</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| GET | /certificates | page, status | - | YES | YES | Get list data user dan masing-masing sertifikat vaksin |
| POST | /certificates  | - | vaccinedose, image, description | YES | NO | Menambahkan data sertifikat vaksin user |
| GET | /mycertificates| - | - | YES | NO | Get data sertifikat vaksin dari user yang sedang login |
| PUT | /mycertificates/:id  | - | image | YES | NO | Edit sertifikat vaksin jika pengajuan ditolak oleh admin |
| GET | /certificates/:id | - | - | YES | NO | Get data sertifikat vaksin berdasarkan id sertifikat |
| PUT | /certificates/:id  | - | status | YES | YES | Edit status pengajuan sertifikat vaksin |
</details>



<details>
<summary>🖥️ &nbsp;Attendances</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| POST | /attendances | - | schedule_id, description, image | YES | NO | Booking jadwal WFO |
| PUT | /attendances/:id  | - | schedule_id, status, status_info | YES | YES | Edit status booking WFO |
| GET | /attendances/:id| - | - | YES | NO | Get data booking WFO by id |
| GET | /myattendances  | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login |
| GET | /mylatestattendances | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login dan diurutkan dari tanggal terbaru|
| GET | /mylongestattendances  | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login dan diurutkan dari tanggal terjauh |
| GET | /pendingattendances  | page | - | YES | YES | Get list data booking WFO yang berstatus pending |
</details>



<details>
<summary>🖥️ &nbsp;Check In and Check Out</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Admin Only | Fungsi |
| --- | --- | --- | --- | --- | --- | --- |
| GET | /checks | page | - | YES | YES | Get list user dan data check in dan checkout |
| GET | /checks/:id  | - | - | YES | NO | Get data check in dan check out by id |
| PUT | /checkin | - | id, temperature | YES | NO | Check in pada saat wfo |
| PUT | /checkout  | - | id | YES | NO | Check out setelah wfo |
</details>
</ul>
</details><br></br>

## 📖 Documentation
For a complete documentation, you can see OpenAPI documentation [here](https://app.swaggerhub.com/apis-docs/mufidi-a/capstone-group-1/1.0.0)<br></br>

# How to Use

### 1. Install app
```bash
git clone https://github.com/Richnuts/CapStoneProject-Group-1.git
```

### 2 Set .env configuration
```bash
AWS_Region
AWS_Access_key_ID
AWS_Secret_access_key
AWS_Bucket
SECRET
Port
DB_Driver
DB_Name 
DB_Address
DB_Port 
DB_Username
DB_Password
```

### 3. Run app
```bash
go run main.go
```

### 4. Run unit testing
```bash
go test ./delivery/... -coverprofile=cover.out && go tool cover -html=cover.out
```
<br></br>

## 🛠 &nbsp;Build App & Database
[![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)](https://aws.amazon.com/)&nbsp;
[![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)](https://www.docker.com/)&nbsp;
[![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)](https://github.com/)&nbsp;
[![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)](https://go.dev/)&nbsp;
[![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)](https://www.json.org/json-en.html)&nbsp;
[![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)]((https://www.mysql.com/))&nbsp;
[![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)](https://www.postman.com/)&nbsp;
[![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)](https://code.visualstudio.com/)&nbsp;<br></br>

## Contact

[![GitHub](https://img.shields.io/badge/-Mufidi-05122A?style=flat&logo=github&logoColor=black)](https://github.com/mufidi-a)
[![GitHub](https://img.shields.io/badge/-Richap-05122A?style=flat&logo=github&logoColor=black)](https://github.com/Richnuts)