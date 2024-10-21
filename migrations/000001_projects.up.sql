
CREATE TABLE IF NOT EXISTS projects (
                                        id INT PRIMARY KEY AUTO_INCREMENT,
                                        title VARCHAR(255) NOT NULL,
                                        description TEXT,
                                        url VARCHAR(255) NOT NULL,
                                        image VARCHAR(255) NOT NULL,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
