# api-genesis

Api golang conversão de moedas

## Configurar arquivo .env

## Criar o banco de dados com o seguinte script

-- Copiando estrutura do banco de dados para genesis
CREATE DATABASE IF NOT EXISTS `genesis`
USE `genesis`;

-- Copiando estrutura para tabela genesis.exchanges
CREATE TABLE IF NOT EXISTS `exchanges` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`amount` decimal(12,2) DEFAULT 0.00,
`fromExchange` varchar(6) DEFAULT NULL,
`toExchange` varchar(6) DEFAULT NULL,
`rate` decimal(12,2) DEFAULT NULL,
`dtCreated` datetime DEFAULT NULL,
PRIMARY KEY (`id`)
);
