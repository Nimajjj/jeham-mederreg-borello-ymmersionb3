CREATE DATABASE IF NOT EXISTS "PetitsRochersGrosCailloux";
USE "PetitsRochersGrosCailloux";


CREATE TABLE IF NOT EXISTS Cailloux (
    ID INT auto_increment,
    Title VARCHAR(255),
    Description VARCHAR(65535),
    Price INT,
    Creation DATE,
    Quantity INT,

    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS Users (
    ID INT auto_increment,
    Email VARCHAR(255),
    EncryptedPwd VARCHAR(255),

    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS Baskets (
    ID INT auto_increment,

    FK_ID_User INT,

    FOREIGN KEY(FK_ID_User) REFERENCES Users(ID),

    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS Categories (
    ID INT auto_increment,
    Title VARCHAR(255),

    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS Photos (
    ID INT auto_increment,
    FilePath VARCHAR(255),

    PRIMARY KEY(ID)
);


CREATE TABLE IF NOT EXISTS Cailloux_Categories (
    ID_Caillou INT,
    ID_Categorie INT,

    FOREIGN KEY(ID_Caillou) REFERENCES Cailloux(ID),
    FOREIGN KEY(ID_Categorie) REFERENCES Categories(ID)
);

CREATE TABLE IF NOT EXISTS Cailloux_Photos (
    ID_Caillou INT,
    ID_Photo INT,

    FOREIGN KEY(ID_Caillou) REFERENCES Cailloux(ID),
    FOREIGN KEY(ID_Photo) REFERENCES Photos(ID)
);

CREATE TABLE IF NOT EXISTS Cailloux_Basket (
    ID_Caillou INT,
    ID_Basket INT,
    Quantity INT,

    FOREIGN KEY(ID_Caillou) REFERENCES Cailloux(ID),
    FOREIGN KEY(ID_Basket) REFERENCES Baskets(ID)
);
