CREATE DATABASE tokokami;

DROP DATABASE tokokami;

USE tokokami;

DROP TABLE nota;
DROP TABLE transaksi;
DROP TABLE pelanggan;
DROP TABLE barang;
DROP TABLE pegawai;

CREATE TABLE pegawai (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL
);

CREATE TABLE barang (
  id_barang int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  id_pegawai int NOT NULL,
  nama_barang varchar(255) NOT NULL,
  info_barang text NOT NULL,
  stok_barang int NOT NULL,
  CONSTRAINT fk_barang_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE pelanggan (
  id_pelanggan int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  id_pegawai int NOT NULL,
  nama varchar(255) NOT NULL,
  CONSTRAINT fk_pelanggan_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE transaksi (
  id_transaksi int AUTO_INCREMENT PRIMARY KEY NOT NULL,
  id_pegawai int NOT NULL,
  id_pelanggan int NOT NULL,
  tgl_transaksi date NOT NULL DEFAULT (DATE(CURRENT_TIMESTAMP)),
  CONSTRAINT fk_transaksi_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT fk_transaksi_pelanggan FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE nota (
  id_transaksi int NOT NULL,
  id_barang int NOT NULL,
  qty int NOT NULL,
  CONSTRAINT fk_nota_transaksi FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT fk_nota_barang FOREIGN KEY (id_barang) REFERENCES barang(id_barang) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Pegawai
INSERT INTO
  pegawai (username, password)
VALUES
  ("x", "z"),
  ("rizal", "rizal123"),
  ("alif", "alif123"),
  ("fahmi", "fahmi123"),
  ("muhamad", "muh123"),
  ("hafidz", "haf123"),
  ("messi", "lm10"),
  ("ronaldo", "cr7"),
  ("alterra", "alta123"),
  ("putri", "put123");

-- Barang
INSERT INTO
  barang (
    id_pegawai,
    nama_barang,
    info_barang,
    stok_barang
  )
VALUES
  (1, "Indomie", "Indomie", 120),
  (1, "Nescafe", "Nescafe", 60),
  (1, "Pop Mie", "Pop mie", 24),
  (2, "Molto", "Molto pink", 48),
  (2, "Rinso", "Rinso cair", 36),
  (3, "Lifeboy", "Sabun mandi", 50),
  (4, "Nuvo", "Sabun mandi cair", 60),
  (5, "Pepsodent", "Pasta gigi", 16),
  (5, "CloseUp", "Pasta gigi warna ijo", 32),
  (5, "Beng Beng", "Rasa coklat", 200),
  (5, "Tango", "Ukuran sedang", 72),
  (7, "Baygon", "Obat nyamuk biar sehat", 12);

-- Pelanggan
INSERT INTO
  pelanggan (id_pegawai, nama)
VALUES
  (1, "Mbappe"),
  (1, "Dewi"),
  (8, "Lukman"),
  (9, "Pevita"),
  (10, "Lucinta");

-- Transaksi
INSERT INTO
  transaksi (id_pegawai, id_pelanggan)
VALUES
  (1, 1),
  (1, 5),
  (2, 4);

-- Nota
INSERT INTO
  nota (id_transaksi, id_barang, qty)
VALUES
  (1, 1, 12),
  (1, 3, 3),
  (1, 5, 6),
  (2, 12, 1),
  (3, 6, 2),
  (3, 10, 10);
  
SELECT t.id_transaksi, t.tgl_transaksi, p.nama , p2.username
FROM transaksi t, pelanggan p, pegawai p2 
WHERE p2.id = t.id_pegawai AND p.id_pelanggan = t.id_pelanggan
ORDER BY 1 DESC LIMIT 1;

SELECT t.id_transaksi, t.tgl_transaksi ,p.nama, p2.username, b.nama_barang, n.qty 
FROM transaksi t, pelanggan p, pegawai p2, barang b, nota n  
WHERE p.id_pelanggan = t.id_pelanggan 
AND p2.id = t.id_pegawai 
AND n.id_transaksi = 1
AND b.id_barang  = n.id_barang;

SELECT * FROM nota WHERE id_transaksi  = 1