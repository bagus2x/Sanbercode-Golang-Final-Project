-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 28 Agu 2023 pada 00.35
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sanbercode`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `category`
--

CREATE TABLE `category` (
  `id` bigint(20) NOT NULL,
  `name` varchar(256) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `category`
--

INSERT INTO `category` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Flora', '2023-08-27 22:14:54.833', '2023-08-27 22:15:04.214'),
(11, 'Animal', '2023-08-27 22:23:30.494', '2023-08-27 22:23:30.494');

-- --------------------------------------------------------

--
-- Struktur dari tabel `comment`
--

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `post_id` bigint(20) NOT NULL,
  `body` longtext NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `comment`
--

INSERT INTO `comment` (`id`, `author_id`, `post_id`, `body`, `created_at`, `updated_at`) VALUES
(1, 2, 18, 'mantab', '2023-08-28 05:28:19.489', '2023-08-28 05:28:19.489');

-- --------------------------------------------------------

--
-- Struktur dari tabel `post`
--

CREATE TABLE `post` (
  `id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `title` varchar(256) NOT NULL,
  `body` longtext NOT NULL,
  `thumbnail` varchar(512) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `post`
--

INSERT INTO `post` (`id`, `author_id`, `title`, `body`, `thumbnail`, `created_at`, `updated_at`) VALUES
(5, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:16:25.527', '2023-08-27 22:16:25.527'),
(8, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:16:51.230', '2023-08-27 22:16:51.230'),
(9, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:17:42.093', '2023-08-27 22:17:42.093'),
(10, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:17:59.461', '2023-08-27 22:17:59.461'),
(11, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:21:59.278', '2023-08-27 22:21:59.278'),
(12, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:23:00.850', '2023-08-27 22:23:00.850'),
(13, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:23:17.334', '2023-08-27 22:23:17.334'),
(14, 1, 'Aku cantik kali s', 'ikan badak ikan hiu', 'https://goola.com/ss.jpg', '2023-08-27 22:23:38.405', '2023-08-27 22:23:38.405'),
(15, 1, 'wdwdwdwdwd', 'ahahahahaha', 'https://goola.com/ss.jpg', '2023-08-27 22:26:25.540', '2023-08-27 22:26:25.540'),
(16, 1, 'wdwdwdwdwd', 'ahahahahaha', 'https://goola.com/ss.jpg', '2023-08-27 22:27:36.322', '2023-08-27 22:27:36.322');

-- --------------------------------------------------------

--
-- Struktur dari tabel `post_categories`
--

CREATE TABLE `post_categories` (
  `category_id` bigint(20) NOT NULL,
  `post_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `post_categories`
--

INSERT INTO `post_categories` (`category_id`, `post_id`) VALUES
(1, 9),
(1, 10),
(1, 12),
(1, 13),
(1, 15),
(1, 16),
(11, 14);

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL,
  `name` varchar(256) NOT NULL,
  `email` varchar(512) NOT NULL,
  `password` varchar(512) NOT NULL,
  `photo` varchar(512) DEFAULT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `photo`, `created_at`, `updated_at`) VALUES
(1, 'bambang', 'bambang@gmail.com', '$2a$10$cBcfvH5XD.k.7icPco5R5O/6tuF6hh15Rv3rh0TcrQNj82GqoDBiu', 'https://sosmed.com/bambang.jpg', '2023-08-27 22:13:55.923', '2023-08-27 22:36:17.239'),
(2, 'Tubagus Saifulloh', 'bagus@gmail.com', '$2a$10$ykzNOMf3ho8dFdZ51VeE8O16tngfXBb5yk6PBQy8LiCqpa0Rqg6LS', 'https://www.google.com/me.jpg', '2023-08-28 04:52:39.304', '2023-08-28 05:00:00.547');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indeks untuk tabel `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_comments` (`author_id`);

--
-- Indeks untuk tabel `post`
--
ALTER TABLE `post`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_posts` (`author_id`);

--
-- Indeks untuk tabel `post_categories`
--
ALTER TABLE `post_categories`
  ADD PRIMARY KEY (`category_id`,`post_id`),
  ADD KEY `fk_post_categories_post` (`post_id`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `category`
--
ALTER TABLE `category`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT untuk tabel `comment`
--
ALTER TABLE `comment`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `post`
--
ALTER TABLE `post`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `comment`
--
ALTER TABLE `comment`
  ADD CONSTRAINT `fk_user_comments` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- Ketidakleluasaan untuk tabel `post`
--
ALTER TABLE `post`
  ADD CONSTRAINT `fk_user_posts` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- Ketidakleluasaan untuk tabel `post_categories`
--
ALTER TABLE `post_categories`
  ADD CONSTRAINT `fk_post_categories_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_post_categories_post` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
