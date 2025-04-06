create database if not exists db_uco;
use db_uco;

drop table if exists trans cascade;
drop table if exists usefas cascade;
drop table if exists dcheckin cascade;
drop table if exists hcheckin cascade;
drop table if exists fasilitas cascade;
drop table if exists customer cascade;
drop table if exists kamar cascade;
drop table if exists jeniskamar cascade;

create table jeniskamar
( kodejenis varchar(3) primary key,
  namajenis varchar(20) not null,
  harga int not null
);

create table kamar
( nokamar varchar(3) primary key,
  kodejenis varchar(3) references jeniskamar(kodejenis),
  status varchar(1) not null
);

create table customer
(
  kodecust varchar(6) primary key,
  nama varchar(20) not null,
  alamat varchar(30) not null,
  kota varchar(15) not null,
  telepon varchar(15) not null
);

create table fasilitas
(
  kodefasilitas varchar(4) primary key,
  namafasilitas varchar(20) not null,
  harga int
);

create table hcheckin
(
  kodecin varchar(6) primary key,
  kodecust varchar(6) references customer(kodecust),
  tglcin date not null,
  jamcin varchar(4) not null
);

create table dcheckin
(
  kodecin varchar(6) references hcheckin(kodecin),
  nokamar varchar(3) references kamar(nokamar),
  tglcout date,
  kodeusefas varchar(9) unique,
  primary key (kodecin,nokamar)
);

create table usefas
(
  kodeusefas varchar(9) references dcheckin(kodeusefas), 
  kodefasilitas varchar(4) references fasilitas(kodefasilitas),
  jumlah int not null
);

create table trans
(
  kodetrans varchar(6) primary key,
  tgltrans date not null,
  jamtrans varchar(4) not null,
  kodecin varchar(6) references hcheckin(kodecin)
);

insert into jeniskamar values ('S01','STANDARD',110000);
insert into jeniskamar values ('D01','DELUXE',175000);
insert into jeniskamar values ('V01','VIP',250000);
insert into jeniskamar values ('P01','PRESIDENT SUITE',400000);

insert into kamar values ('101','S01',0);
insert into kamar values ('112','S01',1);
insert into kamar values ('201','D01',0);
insert into kamar values ('220','V01',0);
insert into kamar values ('405','P01',1);

insert into customer values ('C00001','YOHANES KAREL','RAMPAL KULON 12','MALANG','6503478');
insert into customer values ('C00002','YUSAK WINATA','MASPATI III/25','SURABAYA','8658898');
insert into customer values ('C00003','LISTYA ARINI','MAWAR 14','JAKARTA','5934485');
insert into customer values ('C00004','YOLANDA HIDAYATI','TALANG 1','MALANG','7894455');

insert into fasilitas values ('F001','EXTRA BED',30000);
insert into fasilitas values ('F002','EXTRA HANDUK',15000);
insert into fasilitas values ('F003','SANDAL KAMAR',45000);
insert into fasilitas values ('F004','LAUNDRY',3000);
insert into fasilitas values ('F005','HANGER',12500);

insert into hcheckin values ('T00001','C00001',str_to_date('08-06-2021','%d-%m-%Y'),'0900');
insert into hcheckin values ('T00002','C00002',str_to_date('12-07-2021','%d-%m-%Y'),'0930');
insert into hcheckin values ('T00003','C00003',str_to_date('11-08-2021','%d-%m-%Y'),'0800');
insert into hcheckin values ('T00004','C00001',str_to_date('11-08-2021','%d-%m-%Y'),'1000');

insert into dcheckin values ('T00001','101',str_to_date('12-06-2021','%d-%m-%Y'),'T00001101');
insert into dcheckin values ('T00002','201',str_to_date('14-07-2021','%d-%m-%Y'),'T00002201');
insert into dcheckin values ('T00002','220',str_to_date('14-07-2021','%d-%m-%Y'),NULL);
insert into dcheckin values ('T00003','112',str_to_date('19-08-2021','%d-%m-%Y'),'T00003112');
insert into dcheckin values ('T00004','405',str_to_date('20-08-2021','%d-%m-%Y'),NULL);

insert into usefas values ('T00001101','F001',1);
insert into usefas values ('T00001101','F002',4);
insert into usefas values ('T00002201','F001',2);
insert into usefas values ('T00003112','F004',8);
insert into usefas values ('T00003112','F003',2);

insert into trans values ('N00001',str_to_date('12-06-2021','%d-%m-%Y'),'1000','T00001');
insert into trans values ('N00002',str_to_date('14-07-2021','%d-%m-%Y'),'1300','T00002');



