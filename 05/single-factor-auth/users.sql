CREATE DATABASE IF NOT EXISTS Globomantics;

USE Globomantics;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id INT AUTO_INCREMENT,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO users (username, password) VALUES 
('luke.skywalker', '$2a$09$S2F2gdxnAdVlcehD0drkNOKSVWHxwC6COzROdicjBeXEpZ9LZ0j5a'), 
('han.solo', '$2a$09$wnn34Auc8SpK7Df18IYBlukdutz0dbt6/WFK4lpKbNy0Q0ruWeWGG'), 
('leia.organa', '$2a$09$Iexii9A8hsFf/XhZeHD.S.b8I/10FYykgJOLVlQmqx.23c/6oeAhy'), 
('padme.amidala', '$2a$09$GeOkwgtGG/btGCDO8AkTPe9x/RRpEra0qL18IKBcmXdi2tuPSUaAK'), 
('obiwan.kenobi', '$2a$09$T2fOy8dz6PlM5r9wISOCcu4pLvdlMXs3cs4/0rbq3TVmCDo4WadJ.'), 
('anakin.skywalker', '$2a$09$TcxKEhjfiwD30rS.HSMzoumH0uWpof5E8/J3JNP3vDPcI7yKJmSJO'), 
('mace.windu', '$2a$09$9bPw7mJaLOWfxwwfwpOAN.Z2C.xWEx/tXhI8JDi0.QHK3t4mUmGiq'), 
('qui.gon.jinn', '$2a$09$4er91VLnM3vSnR.acNiEEe5pZ9l54W/RzkxrXvjgn6BdBC4gurlLK'), 
('ahsoka.tano', '$2a$09$uO1/r5ZDRl7lobws/RYPTeaGQ3sYbvvNYTAWmGxBXkz/Piy3MLJje');
