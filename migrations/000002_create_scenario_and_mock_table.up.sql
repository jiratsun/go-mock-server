CREATE TABLE scenario (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    correlation_key VARCHAR(20) NOT NULL UNIQUE,
    current_sequence TINYINT UNSIGNED NOT NULL DEFAULT 0,
    name VARCHAR(255) NOT NULL DEFAULT '',
    description VARCHAR(255) NOT NULL DEFAULT '',
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE mock (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    scenario_id INT UNSIGNED UNIQUE,
    description VARCHAR(255) NOT NULL DEFAULT '',
    detail JSON NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (scenario_id) REFERENCES scenario (id)
);
