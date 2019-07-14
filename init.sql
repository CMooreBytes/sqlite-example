CREATE TABLE Persons(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name NVARCHAR(100) NOT NULL,
    last_name NVARCHAR(100) NOT NULL
);
CREATE TABLE Services(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    service_date date NOT NULL
);
CREATE TABLE Assignments (
    person_id INTEGER NOT NULL,
    service_id INTEGER NOT NULL,
    PRIMARY KEY (person_id, service_id),
    FOREIGN KEY(person_id) REFERENCES Persons(id),
    FOREIGN KEY(service_id) REFERENCES Services(id)
);

INSERT INTO Persons (first_name, last_name) VALUES ('First', 'Last');
INSERT INTO Services (service_date) VALUES ('2019-07-14');
INSERT INTO Assignments VALUEs (1,1);