-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 23, 2022 at 04:23 PM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 8.0.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_crowdfunding`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'test CaMpaIgn POSTMAN different SLuG', 'campaign test from postman', 'long description', 'perks 1, perks 2, perks 3', 0, 10000000, 0, 'test-campaign-postman-different-slug-1', '2022-05-14 12:55:20', '2022-05-14 12:55:20'),
(2, 1, 'test CaMpaIgn POSTMAN', 'campaign test from postman Updated', 'long description Updated', 'perks 1, perks 2, perks 3, perks 4, perks 5', 0, 999000000, 0, 'campaign-two', '2022-05-10 15:53:42', '2022-05-14 15:57:16'),
(3, 2, 'Campaign 3', 'testt', 'testsadsadasd', 'perk1, 213, 213123', 0, 20100000, 0, 'campaign-three', '2022-05-10 15:55:35', '2022-05-10 15:55:35'),
(4, 1, 'Crowdfunding for start up', 'Startup is a new company', 'Startup is a new company and we need to raise money for it', 'reward 1, reward 2, reward 3', 0, 100000000, 0, 'crowdfunding-for-start-up-1', '2022-05-14 12:21:59', '2022-05-14 12:21:59'),
(5, 2, 'Crowdfunding for start up', 'Startup is a new company', 'Startup is a new company and we need to raise money for it', 'reward 1, reward 2, reward 3', 0, 100000000, 0, 'crowdfunding-for-start-up-2', '2022-05-14 12:23:54', '2022-05-14 12:23:54'),
(6, 1, 'test CaMpaIgn POSTMAN', 'campaign test from postman', 'long description', 'perks 1, perks 2, perks 3', 0, 10000000, 0, 'test-campaign-postman-1', '2022-05-14 12:46:31', '2022-05-14 12:46:31');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(9, 1, 'campaign-images/1-1652542455367-profile.png', 1, '2022-05-14 22:34:15', '2022-05-14 22:34:15');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `payment_url` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `payment_url`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 100000, 'paid', NULL, '', '2022-05-15 16:42:09', '2022-05-15 16:42:09'),
(2, 1, 2, 350000, 'pending', NULL, '', '2022-05-15 16:42:09', '2022-05-15 16:42:09'),
(3, 1, 1, 500000, 'pending', NULL, '', '2022-05-15 16:58:44', '2022-05-15 16:58:44'),
(4, 5, 1, 1500000, 'pending', 'TRX-511653142571370', '', '2022-05-21 21:16:11', '2022-05-21 21:16:11'),
(5, 2, 1, 10000000, 'pending', 'TRX-211653143563270', '', '2022-05-21 21:32:43', '2022-05-21 21:32:43');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Alvin Martin', 'Developer', 'alvin@gmail.com', '$2a$04$q0yIbLB6T1HLb6qghWDCi.McuB8R3UvooXRGwsK4RJdwgjyjE0DYi', 'avatars/1-1652368121492-profile2.png', 'user', '', '2022-05-04 21:56:35', '2022-05-12 22:08:41'),
(2, 'Veiros', 'Gamers', 'veiros@gmail.com', '$2a$04$FbLO/mWF8AX1tcw7YRwoeOi8/JuoeqgjmoKPUHk7v2YMOvKr15c8K', 'avatars/2-1652021866887-profile.png', 'user', '', '2022-05-07 15:02:10', '2022-05-08 21:53:15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
