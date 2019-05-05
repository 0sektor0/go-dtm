USE go_dtm;


DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS attachment CASCADE;
DROP TABLE IF EXISTS attachment_type CASCADE;
DROP TABLE IF EXISTS task CASCADE;
DROP TABLE IF EXISTS task_status CASCADE;
DROP TABLE IF EXISTS task_type CASCADE;
DROP TABLE IF EXISTS developer CASCADE;


CREATE TABLE developer (
    id int AUTO_INCREMENT PRIMARY KEY,
    login varchar(50) NOT NULL UNIQUE,
    password varchar(50) NOT NULL,
    picture varchar(200) NOT NULL,
    is_admin tinyint DEFAULT 0
);

CREATE TABLE task_type (
    id int AUTO_INCREMENT PRIMARY KEY,
    task_type_name varchar(50) NOT NULL
);

CREATE TABLE task_status (
    id int AUTO_INCREMENT PRIMARY KEY,
    task_status_name varchar(50) NOT NULL
);

CREATE TABLE task (
    id int AUTO_INCREMENT PRIMARY KEY,
    creator_id int,
    asignee_id int,
    task_type_id int NOT NULL,
    task_status_id int,
    title varchar(200) NOT NULL,
    task_text varchar(2000) NOT NULL,
    creation_date Date NOT NULL,
    start_date Date NOT NULL,
    end_date Date NOT NULL,
    update_date Date NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES developer(id) ON DELETE SET NULL,
    FOREIGN KEY (asignee_id) REFERENCES developer(id) ON DELETE SET NULL,
    FOREIGN KEY (task_type_id) REFERENCES task_type(id) ON DELETE CASCADE,
    FOREIGN KEY (task_status_id) REFERENCES task_status(id) ON DELETE SET NULL
);

CREATE TABLE attachment_type (
    id int AUTO_INCREMENT PRIMARY KEY,
    attachment_type_name varchar(50) NOT NULL
);

CREATE TABLE attachment (
    id int AUTO_INCREMENT PRIMARY KEY,
    task_id int NOT NULL,
    attachment_type_id int,
    attachment_path varchar(200) NOT NULL,
    creation_date Date NOT NULL,
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE,
    FOREIGN KEY (attachment_type_id) REFERENCES attachment_type(id) ON DELETE SET NULL
);

CREATE TABLE comment (
    id int AUTO_INCREMENT PRIMARY KEY,
    developer_id int,
    task_id int NOT NULL,
    comment_text varchar(1000),
    FOREIGN KEY (developer_id) REFERENCES developer(id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE
);



CREATE INDEX developer_id_idx ON developer(id);
CREATE INDEX attachment_id_idx ON attachment(id);
CREATE INDEX attachment_type_id_idx ON attachment_type(id);
CREATE INDEX comment_id_idx ON comment(id);
CREATE INDEX task_id_idx ON task(id);
CREATE INDEX task_Status_id_idx ON task_status(id);
CREATE INDEX task_type_id_idx ON task_type(id);
CREATE INDEX developer_login_idx ON developer(login);