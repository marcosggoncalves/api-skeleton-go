DROP DATABASE IF EXISTS usuarios;

CREATE DATABASE usuarios;

USE usuarios;

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
  usuario_tipo_id INTEGER UNSIGNED, -- Adicionado UNSIGNED para coincidir com o tipo do campo na tabela `usuario_tipo`
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_usuario_usuario_tipo -- Nomeando a constraint de chave estrangeira
    FOREIGN KEY (usuario_tipo_id) 
    REFERENCES usuario_tipo (id)
    ON DELETE SET NULL -- Configuração sugerida para evitar problemas ao excluir registros
    ON UPDATE CASCADE -- Para manter a consistência em caso de alterações no `id` da tabela referenciada
);
