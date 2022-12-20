CREATE DATABASE tokokami;

USE tokokami;

CREATE TABLE pegawai (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL 
);

CREATE TABLE barang (
   id_barang int NOT NULL AUTO_INCREMENT PRIMARY KEY,
   nama_barang varchar(255) NOT NULL,
   info_barang text NOT NULL, 
   stok_barang int NOT NULL,
);

CREATE TABLE pelanggan (
  id_pelanggan int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL 
);

CREATE TABLE transaksi (
  id_transaksi int AUTO_INCREMENT NOT NULL,
  id_pelanggan int NOT NULL,
  id_barang int NOT NULL,
  nama_barang varchar(255) NOT NULL ,
  id_pegawai int NOT NULL,
  tgl_transaksi date NOT NULL DEFAULT (DATE(CURRENT_TIMESTAMP))
  CONSTRAINT fk_transaksi_pelanggan FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
  CONSTRAINT fk_transaksi_barang FOREIGN KEY (id_barang) REFERENCES barang(id_barang),
  CONSTRAINT fk_transaksi_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai)
);

CREATE TABLE nota (
  id_transaksi int NOT NULL,
  id_pegawai int NOT NULL,
  id_barang int NOT NULL,
  qty int NOT NULL,
  CONSTRAINT fk_nota_transaksi FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi),
  CONSTRAINT fk_transaksi_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai),
  CONSTRAINT fk_transaksi_barang FOREIGN KEY (id_barang) REFERENCES barang(id_barang)
);