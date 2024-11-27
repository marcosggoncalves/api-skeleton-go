CREATE TABLE usuario_tipo (
  id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  nome VARCHAR(255) NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

INSERT INTO usuario_tipo (nome) values ("Comum");
INSERT INTO usuario_tipo (nome) values ("Especial");

CREATE TABLE usuario (
  id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  nome VARCHAR(255) NULL,
  cpf VARCHAR(45) NULL,
  email VARCHAR(45) NULL,
  senha VARCHAR(255) NULL,
  usuario_tipo_id INTEGER UNSIGNED, 
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_usuario_usuario_tipo 
    FOREIGN KEY (usuario_tipo_id) 
    REFERENCES usuario_tipo (id)
    ON DELETE SET NULL 
    ON UPDATE CASCADE 
);

INSERT INTO `usuario` (`id`, `nome`, `cpf`, `email`, `senha`, `usuario_tipo_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Marcos Lopes', '000.000.000-10', 'marcos.ggoncalves.pr@gmail.com', '$2a$10$bNRfGNjv31x9qLLif.9EOe7Vd5XFSAOlhJMi3a8f6nHhm0OY6ABF.', 1, '2024-11-27 03:19:44', '2024-11-27 03:19:44', NULL);
