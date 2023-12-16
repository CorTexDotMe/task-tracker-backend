CREATE TABLE IF NOT EXISTS tasks(
    taskID SERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    status VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL,
    dateCreated DATE NOT NULL,
    dueDate DATE,
    userID INTEGER NOT NULL,
    PRIMARY KEY(taskID),
    FOREIGN KEY(userID) REFERENCES Users(userID)
);