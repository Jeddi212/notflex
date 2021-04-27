-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 27, 2021 at 06:34 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_notflex`
--

-- --------------------------------------------------------

--
-- Table structure for table `credits`
--

CREATE TABLE `credits` (
  `card_number` varchar(191) NOT NULL,
  `exp` longtext DEFAULT NULL,
  `cvc` longtext DEFAULT NULL,
  `user_id` varchar(191) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `credits`
--

INSERT INTO `credits` (`card_number`, `exp`, `cvc`, `user_id`) VALUES
('00828765814', '4/24', '989', 'jeddi@gmail.com'),
('123456789', '07/30', '123', 'rafael22@gmail.com'),
('77828028208124', '6/26', '787', 'fedly45@gmail.com'),
('9876543', '08/30', '333', 'jedediah@gmail.com');

-- --------------------------------------------------------

--
-- Table structure for table `films`
--

CREATE TABLE `films` (
  `id` bigint(20) NOT NULL,
  `title` longtext DEFAULT NULL,
  `genre` longtext DEFAULT NULL,
  `year` longtext DEFAULT NULL,
  `director` longtext DEFAULT NULL,
  `actor` longtext DEFAULT NULL,
  `synopsis` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `films`
--

INSERT INTO `films` (`id`, `title`, `genre`, `year`, `director`, `actor`, `synopsis`) VALUES
(1, 'Aladin Naik Haji The Series VI', 'Adventure; Comedy; Romance;Thriller', '2008', 'Reze Cantona', 'Kenny GG; Sudirman L. G.', 'Dikisahkan sebuah kerajaan pasir di jaman es yang sangat dingin. Hiduplah seorang pangeran bernama aladin. Ia sangat gagah perkasa. Namun kesombongan menjadikannya terkutuk dan harus dicampakkan dari kerjaannya. Bagaimana kisah Aladin bertahan hidup dan mengakhiri kutuk tersebut?'),
(2, 'Aladin Naik Haji 2', 'Adventure; Comedy; Romance', '2007', 'Udin Cantonas', 'Kenny GG; Sudirman L. G.', 'Dikisahkan sebuah kerajaan pasir di jaman es yang sangat dingin. Hiduplah seorang pangeran bernama aladin. Ia sangat gagah perkasa. Namun kesombongan menjadikannya terkutuk dan harus dicampakkan dari kerjaannya. Bagaimana kisah Aladin bertahan hidup dan mengakhiri kutuk tersebut?'),
(3, 'Khaleed Numpang Lewat', 'Fantasy; Historical', '2016', 'Septian Wakwaw', 'Khaleed; Selena; Claude', 'Numpang lewat ke gurun sahara. Namun ada kejadian tak terduga, ia dihadang oleh bandit yang legendaris, Claude, amun Khaleed pantang menyerah.'),
(4, 'Aladin Naik Haji', 'Adventure; Comedy', '2006', 'Udin Cantona', 'Kenny G; Sudirman L. G.', 'Dikisahkan sebuah kerajaan pasir di jaman es yang sangat dingin. Hiduplah seorang pangeran bernama aladin. Ia sangat gagah perkasa. Namun kesombongan menjadikannya terkutuk dan harus dicampakkan dari kerjaannya. Bagaimana kisah Aladin bertahan hidup dan mengakhiri kutuk tersebut?'),
(5, 'Seribu Kisah Lampau', 'Historical; Military', '2006', 'Lylia Mendar', 'Lukman Akbar; Selina Wijaya; Lutfi Gerja', 'Hmmm'),
(6, 'Seribu Kisah Lampau 2', 'Historical; Military', '2008', 'Lylia Mendar', 'Lukman Akbar; Selina Wijaya; Lutfi Gerja', 'Hmmm'),
(7, 'Fedly : Sang Pencari Cinta Sejati Vol. 2', 'Horror; Thriller; Comedy; Romance; Sci-Fi; Action; Harem; Isekai; Overpower', '2021', 'Timothy Gaskeun', 'Fedly Makaroi; Taman Susi Susanti; Novia; Jedediah', 'Fedly sedang mencari cinta di bumi ter cinta. Namun cintanya bertepuk sebelah tangan. Tetapi Dia selalu setia menunggu Novia sang orang ketiga. Lalu muncul seorang pria bernama Jedediah. Bagaimana kisah kelanjutan Fedly berburu cinta.'),
(8, 'Algoritma Lanjut', 'Comedy', '2022', 'Halo', 'Bu Inge', 'Mengkisahkan algoritma'),
(9, 'Algoritma Lanjut', 'Horror', '2023', 'Hai', 'Bu Inge; Pa Hery', 'Mengkisahkan algoritma yang menyeramkan');

-- --------------------------------------------------------

--
-- Table structure for table `histories`
--

CREATE TABLE `histories` (
  `id` bigint(20) NOT NULL,
  `user_email` varchar(191) DEFAULT NULL,
  `film_id` bigint(20) DEFAULT NULL,
  `date` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `histories`
--

INSERT INTO `histories` (`id`, `user_email`, `film_id`, `date`) VALUES
(1, 'jedediah@gmail.com', 2, '2021-04-26 21:39:41.008'),
(2, 'jedediah@gmail.com', 4, '2021-04-26 21:39:46.191'),
(3, 'jedediah@gmail.com', 6, '2021-04-26 21:39:50.513'),
(4, 'jedediah@gmail.com', 6, '2021-04-26 21:39:52.109');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `email` varchar(191) NOT NULL,
  `password` longtext DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `birth_date` longtext DEFAULT NULL,
  `gender` longtext DEFAULT NULL,
  `nationality` longtext DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `subscribe` longtext DEFAULT NULL,
  `sub_date` datetime(3) DEFAULT NULL,
  `level` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`email`, `password`, `name`, `birth_date`, `gender`, `nationality`, `status`, `subscribe`, `sub_date`, `level`) VALUES
('admin@admin.com', 'admin', 'admin', NULL, NULL, NULL, 'active', NULL, NULL, 0),
('fedly45@gmail.com', 'fedly45', 'Fedly Septian', '', 'male', 'Indonesian', 'suspend', '', '0000-00-00 00:00:00.000', 1),
('jeddi@gmail.com', 'jeddi123', 'Jedediah Fanuel', '', 'male', 'Indonesian', 'suspend', 'Basic', '2021-05-24 16:12:31.120', 1),
('jedediah23@gmail.com', 'jedediah23', 'Jedediah23', '', 'male', 'Nigerian', 'active', '', '0000-00-00 00:00:00.000', 1),
('jedediah2@gmail.com', 'jedediah2', 'Jedediah2', '', 'male', 'Nigerian', 'suspend', '', '0000-00-00 00:00:00.000', 1),
('jedediah@gmail.com', 'jedediah', 'Jedediah', '', 'male', 'Nigerian', 'active', 'Premium', '2021-04-26 21:39:12.200', 1),
('nehinehi@yahoo.com', 'nehinehi', 'Nehimala', '09/07/03', 'Female', 'Indian', 'active', '', '0000-00-00 00:00:00.000', 1),
('rafael22@gmail.com', 'rafael22', 'Rafael Gaskeun Sih', '01/12/21', 'male', 'Russian', 'active', 'Premium', '2021-04-24 19:18:44.571', 1),
('rafael33@gmail.com', 'rafael33', 'Rafael Christo', '', 'male', 'Russian', 'active', 'Premium', '2021-05-24 16:23:00.001', 1),
('saya@gmail.com', 'saya', 'Saya', '01/05/19', 'Female', 'American', 'active', '', '0000-00-00 00:00:00.000', 1),
('sunwell@gmail.com', 'sunwell', 'Sunwell', '01/20/20', 'Male', 'Indonesian', 'active', '', '0000-00-00 00:00:00.000', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `credits`
--
ALTER TABLE `credits`
  ADD PRIMARY KEY (`card_number`),
  ADD KEY `idx_credits_user_id` (`user_id`);

--
-- Indexes for table `films`
--
ALTER TABLE `films`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `histories`
--
ALTER TABLE `histories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_histories_user_email` (`user_email`),
  ADD KEY `idx_histories_film_id` (`film_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `films`
--
ALTER TABLE `films`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `histories`
--
ALTER TABLE `histories`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `credits`
--
ALTER TABLE `credits`
  ADD CONSTRAINT `fk_credits_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`email`);

--
-- Constraints for table `histories`
--
ALTER TABLE `histories`
  ADD CONSTRAINT `fk_histories_film` FOREIGN KEY (`film_id`) REFERENCES `films` (`id`),
  ADD CONSTRAINT `fk_histories_user` FOREIGN KEY (`user_email`) REFERENCES `users` (`email`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
