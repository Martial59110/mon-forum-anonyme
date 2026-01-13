# Changelog

## [1.0.1](https://github.com/Martial59110/mon-forum-anonyme/compare/v1.0.0...v1.0.1) (2026-01-13)


### Bug Fixes

* **ci:** fix  DB_HOST ([e45a279](https://github.com/Martial59110/mon-forum-anonyme/commit/e45a2791765082cc42a65e08bfa9625573476d5d))
* **ci:** fix DB_HOST ([4b4739e](https://github.com/Martial59110/mon-forum-anonyme/commit/4b4739eedb457f9059384f0bf887b06e6818ef86))
* **front:** use VPS hostname for API calls ([4f3191a](https://github.com/Martial59110/mon-forum-anonyme/commit/4f3191a39d8edf6e8741230ed7d89f1735677442))
* **swarm:** avoid overlay network on broken VPS ([78ed338](https://github.com/Martial59110/mon-forum-anonyme/commit/78ed338e031c2f2fe8987ea70161f095cdd9a9e7))
* **swarm:** use host ports and pin services to manager ([5dd5d48](https://github.com/Martial59110/mon-forum-anonyme/commit/5dd5d485e054225db510451ed88761ed8cdd2f7c))
* **swarm:** use stop-first updates with host ports ([d26ebd4](https://github.com/Martial59110/mon-forum-anonyme/commit/d26ebd48dedaf509c1cbb7b1a2248f7d977a27d8))

## 1.0.0 (2026-01-10)


### Features

* **api:** implement REST API endpoints for messages ([29dd20a](https://github.com/Martial59110/mon-forum-anonyme/commit/29dd20ad51b627c511ce021866459dfba6741815))
* **database:** add database connection and table creation ([d0c8a31](https://github.com/Martial59110/mon-forum-anonyme/commit/d0c8a3138abf7ba25305e79a5e48057f78e1c8b1))
* **docker:** add environment variables for API service ([4c33a0e](https://github.com/Martial59110/mon-forum-anonyme/commit/4c33a0eec479857a5d9392398066de9ee58ad9db))
* **frontend:** add comprehensive animation system ([118ffd9](https://github.com/Martial59110/mon-forum-anonyme/commit/118ffd981e8b4a9af88588614af3ae5370a33d45))
* **frontend:** add Google Fonts and improve HTML metadata ([affdd62](https://github.com/Martial59110/mon-forum-anonyme/commit/affdd6275bb5a142041fed9a4dd7d093d4c6fe3c))
* **frontend:** add icons and animations to MessageList ([932394f](https://github.com/Martial59110/mon-forum-anonyme/commit/932394f968e2ab8eea33689625d8f6d61bc96f20))
* **frontend:** add icons and success feedback to MessageForm ([7b664f7](https://github.com/Martial59110/mon-forum-anonyme/commit/7b664f7d6e1f609f020175abaab95e53b75da4f4))
* **frontend:** add MessageForm component with validation ([243cb69](https://github.com/Martial59110/mon-forum-anonyme/commit/243cb69bafd668868d39ae7db3ca205d26b04410))
* **frontend:** add MessageList component with API integration ([c5dee8f](https://github.com/Martial59110/mon-forum-anonyme/commit/c5dee8f7655b0e26956b706e3177e717c8e264d7))
* **frontend:** integrate animations and enhance component imports ([36d6dad](https://github.com/Martial59110/mon-forum-anonyme/commit/36d6dad8d5d2ccbdf1eb8c36959847526277f15a))
* **frontend:** integrate MessageForm and MessageList components ([ec5d6ee](https://github.com/Martial59110/mon-forum-anonyme/commit/ec5d6ee80c7f4da72b90882c890cdd13c2de54ae))
* **frontend:** use two-column layout structure ([78f414c](https://github.com/Martial59110/mon-forum-anonyme/commit/78f414c4acaa15f149b3eb3e40d3bf039481cba8))
* **migrations:** add messages table migration ([5df655a](https://github.com/Martial59110/mon-forum-anonyme/commit/5df655af0c9e1e371e329c301429acd1d04281f0))
* **models:** add Message model ([10e442b](https://github.com/Martial59110/mon-forum-anonyme/commit/10e442b0175f6c2b3bca6332887316e99f6b6c33))
* **repository:** add message repository with CRUD operations ([a02bacc](https://github.com/Martial59110/mon-forum-anonyme/commit/a02bacc187678dc3035c44d44ab321b7091ba075))
* **swarm:** add docker stack file for swarm deployment ([3d468b1](https://github.com/Martial59110/mon-forum-anonyme/commit/3d468b1a275d6df31998e522aa480508edff3940))


### Bug Fixes

* **api:** add database connection retry logic ([397b0cf](https://github.com/Martial59110/mon-forum-anonyme/commit/397b0cf69751d431d9dc25565174794dc2713cb9))
* **frontend:** add null safety checks in MessageList component ([4a6b8c6](https://github.com/Martial59110/mon-forum-anonyme/commit/4a6b8c65e3e19b105f086702b597f6f757fc744e))
* **frontend:** move CSS files to correct component directories ([fe552fe](https://github.com/Martial59110/mon-forum-anonyme/commit/fe552fee556ef029e6d120085137df7ccb160994))
* **main.go:** add minimal server ([2239f55](https://github.com/Martial59110/mon-forum-anonyme/commit/2239f5514d45e94d935941521756ea03e8860dd0))
